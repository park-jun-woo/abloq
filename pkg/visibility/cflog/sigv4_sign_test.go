//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what sigv4 서명을 AWS 공개 테스트 벡터로 검증 — iam ListUsers 예제의 서명값과 파생 키, 세션 토큰 헤더 포함 여부
package cflog

import (
	"encoding/hex"
	"net/http"
	"strings"
	"testing"
	"time"
)

// TestSignRequest checks the signer against the AWS-published Signature
// Version 4 example (docs "Example: Signature calculations"): GET
// https://iam.amazonaws.com/?Action=ListUsers&Version=2010-05-08 at
// 20150830T123600Z with the AKIDEXAMPLE credentials must produce signature
// 5d672d79c15b13162d9279b0855cfba6789a8edb4c82c400e06b5924a6f2b5d7, and the
// derived signing key for 20150830/us-east-1/iam is the documented
// c4afb1cc5771d871763a393e44b703571b55cc28424d1a5e86da6ed3c154a4b9.
func TestSignRequest(t *testing.T) {
	key := signingKey("wJalrXUtnFEMI/K7MDENG+bPxRfiCYEXAMPLEKEY", "20150830", "us-east-1", "iam")
	wantKey := "c4afb1cc5771d871763a393e44b703571b55cc28424d1a5e86da6ed3c154a4b9"
	if got := hex.EncodeToString(key); got != wantKey {
		t.Errorf("signingKey = %s, want %s", got, wantKey)
	}

	req, err := http.NewRequest(http.MethodGet, "https://iam.amazonaws.com/?Action=ListUsers&Version=2010-05-08", nil)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	now := time.Date(2015, 8, 30, 12, 36, 0, 0, time.UTC)
	signRequest(req, "AKIDEXAMPLE", "wJalrXUtnFEMI/K7MDENG+bPxRfiCYEXAMPLEKEY", "", "us-east-1", "iam", now, emptyPayloadHash)

	auth := req.Header.Get("Authorization")
	wantSig := "Signature=5d672d79c15b13162d9279b0855cfba6789a8edb4c82c400e06b5924a6f2b5d7"
	if !strings.HasSuffix(auth, wantSig) {
		t.Errorf("Authorization = %q, want suffix %q", auth, wantSig)
	}
	wantSigned := "SignedHeaders=content-type;host;x-amz-date"
	if !strings.Contains(auth, wantSigned) {
		t.Errorf("Authorization = %q, want %q", auth, wantSigned)
	}
	if req.Header.Get("X-Amz-Date") != "20150830T123600Z" {
		t.Errorf("X-Amz-Date = %q", req.Header.Get("X-Amz-Date"))
	}
	if req.Header.Get("X-Amz-Security-Token") != "" {
		t.Errorf("session token header set without a token")
	}

	tokReq, _ := http.NewRequest(http.MethodGet, "https://b.s3.us-east-1.amazonaws.com/", nil)
	signRequest(tokReq, "AKIDEXAMPLE", "secret", "tok", "us-east-1", "s3", now, emptyPayloadHash)
	if tokReq.Header.Get("X-Amz-Security-Token") != "tok" {
		t.Errorf("X-Amz-Security-Token not set for session credentials")
	}
}
