#!/usr/bin/env python3
"""Mock external archive APIs for the Hurl scenarios (Phase008).

Serves the external endpoints pkg/archive and pkg/visibility talk to, on
one port:

  POST /save                        Wayback SPN2  (5xx when the url form
                                    field contains "boom" — partial-failure
                                    scenario)
  POST /indexnow                    IndexNow batch submission
  POST /token                       OAuth2 jwt-bearer token exchange (GSC)
  POST /v3/urlNotifications:publish GSC Indexing API publish
  POST .../searchAnalytics/query    GSC Search Analytics: two fixture pages
                                    per requested day (Phase013)
  POST /v1/urlInspection/index:inspect  GSC URL Inspection: PASS verdict
  POST /chat/completions            Perplexity engine — citations array
  POST /v1/responses                OpenAI engine — url_citation annotation
  POST /v1/messages                 Anthropic engine — text block citations

Every engine answer cites https://fixture.example.com/... so the sampling
scenario records cited=true rows against the shared fixture blog.

GET/HEAD on any path answers the link-rot checker (Phase010): paths
containing "dead" are 404, everything else 200. Point the checker at it
with LINKCHECK_HOST_OVERRIDE=http://127.0.0.1:<port> so the fixture
citation URLs converge on the stub (no real network).

Point abloqd at it with:
  WAYBACK_BASE_URL=http://127.0.0.1:<port>
  INDEXNOW_ENDPOINT=http://127.0.0.1:<port>/indexnow
  GSC_API_BASE=http://127.0.0.1:<port>
  GSC_TOKEN_URL=http://127.0.0.1:<port>/token
  GSC_SEARCH_API_BASE=http://127.0.0.1:<port>
  PERPLEXITY_BASE_URL=http://127.0.0.1:<port>
  OPENAI_BASE_URL=http://127.0.0.1:<port>
  ANTHROPIC_BASE_URL=http://127.0.0.1:<port>

Usage: archive_stub.py [port]   (default 8099)
"""
import json
import sys
from http.server import BaseHTTPRequestHandler, HTTPServer
from urllib.parse import parse_qs


class StubHandler(BaseHTTPRequestHandler):
    def log_message(self, fmt, *args):  # quiet
        pass

    def _reply(self, code, payload):
        body = json.dumps(payload).encode()
        self.send_response(code)
        self.send_header("Content-Type", "application/json")
        self.send_header("Content-Length", str(len(body)))
        self.end_headers()
        self.wfile.write(body)

    def _rot_code(self):
        return 404 if "dead" in self.path else 200

    def do_GET(self):
        self._reply(self._rot_code(), {"path": self.path})

    def do_HEAD(self):
        self.send_response(self._rot_code())
        self.end_headers()

    def do_POST(self):
        length = int(self.headers.get("Content-Length", 0))
        raw = self.rfile.read(length).decode("utf-8", "replace")
        if self.path == "/save":
            url = parse_qs(raw).get("url", [""])[0]
            if "boom" in url:
                self._reply(502, {"error": "stub: spn2 capture exploded"})
                return
            self._reply(200, {"url": url, "job_id": "stub-job-1"})
        elif self.path == "/indexnow":
            self._reply(200, {"ok": True})
        elif self.path == "/token":
            self._reply(200, {"access_token": "stub-token", "expires_in": 3600,
                              "token_type": "Bearer"})
        elif self.path == "/v3/urlNotifications:publish":
            try:
                url = json.loads(raw).get("url", "")
            except ValueError:
                url = ""
            self._reply(200, {"urlNotificationMetadata": {"url": url}})
        elif "/searchAnalytics/query" in self.path:
            self._reply(200, {"rows": [
                {"keys": ["https://fixture.example.com/tech/post-a/"],
                 "clicks": 3, "impressions": 120, "ctr": 0.025, "position": 4.2},
                {"keys": ["https://fixture.example.com/tech/post-b/"],
                 "clicks": 1, "impressions": 40, "ctr": 0.025, "position": 9.8},
            ]})
        elif self.path.endswith("/urlInspection/index:inspect"):
            self._reply(200, {"inspectionResult": {"indexStatusResult": {
                "verdict": "PASS", "coverageState": "Submitted and indexed"}}})
        elif self.path == "/chat/completions":
            self._reply(200, {
                "citations": ["https://fixture.example.com/tech/post-a/",
                              "https://other.example.org/x"],
                "choices": [{"message": {"content": "stub answer"}}]})
        elif self.path == "/v1/responses":
            self._reply(200, {"output": [
                {"type": "web_search_call"},
                {"type": "message", "content": [{
                    "type": "output_text", "text": "stub answer",
                    "annotations": [{"type": "url_citation",
                                     "url": "https://fixture.example.com/tech/post-b/"}]}]},
            ]})
        elif self.path == "/v1/messages":
            self._reply(200, {"content": [
                {"type": "server_tool_use"},
                {"type": "text", "text": "stub answer",
                 "citations": [{"type": "web_search_result_location",
                                "url": "https://fixture.example.com/tech/post-a/"}]},
            ]})
        else:
            self._reply(404, {"error": "stub: unknown path " + self.path})


def main():
    port = int(sys.argv[1]) if len(sys.argv) > 1 else 8099
    server = HTTPServer(("127.0.0.1", port), StubHandler)
    print(f"archive stub listening on 127.0.0.1:{port}", flush=True)
    server.serve_forever()


if __name__ == "__main__":
    main()
