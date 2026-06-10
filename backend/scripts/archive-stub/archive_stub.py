#!/usr/bin/env python3
"""Mock external archive APIs for the Hurl scenarios (Phase008).

Serves the four endpoints pkg/archive talks to, on one port:

  POST /save                        Wayback SPN2  (5xx when the url form
                                    field contains "boom" — partial-failure
                                    scenario)
  POST /indexnow                    IndexNow batch submission
  POST /token                       OAuth2 jwt-bearer token exchange (GSC)
  POST /v3/urlNotifications:publish GSC Indexing API publish

GET/HEAD on any path answers the link-rot checker (Phase010): paths
containing "dead" are 404, everything else 200. Point the checker at it
with LINKCHECK_HOST_OVERRIDE=http://127.0.0.1:<port> so the fixture
citation URLs converge on the stub (no real network).

Point abloqd at it with:
  WAYBACK_BASE_URL=http://127.0.0.1:<port>
  INDEXNOW_ENDPOINT=http://127.0.0.1:<port>/indexnow
  GSC_API_BASE=http://127.0.0.1:<port>
  GSC_TOKEN_URL=http://127.0.0.1:<port>/token

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
        else:
            self._reply(404, {"error": "stub: unknown path " + self.path})


def main():
    port = int(sys.argv[1]) if len(sys.argv) > 1 else 8099
    server = HTTPServer(("127.0.0.1", port), StubHandler)
    print(f"archive stub listening on 127.0.0.1:{port}", flush=True)
    server.serve_forever()


if __name__ == "__main__":
    main()
