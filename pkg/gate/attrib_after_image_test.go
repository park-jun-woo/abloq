//ff:func feature=gate type=parser control=iteration dimension=1
//ff:what attribAfterImage가 이미지 다음 첫 비공백 라인의 이탤릭 표기를 찾고 일반 텍스트·볼드는 거부하는지 검증
package gate

import (
	"strings"
	"testing"
)

func TestAttribAfterImage(t *testing.T) {
	cases := []struct {
		name, body string
		want       int
	}{
		{"direct", "![i](/i)\n*Image: AI generated*\n", 1},
		{"blank between", "![i](/i)\n\n*Image: AI generated*\n", 2},
		{"plain text", "![i](/i)\ntext\n", -1},
		{"bold is not attribution", "![i](/i)\n**bold**\n", -1},
		{"nothing after", "![i](/i)\n", -1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d := &Doc{FirstContentLine: 0, FirstIsImage: true, BodyLines: strings.Split(tc.body, "\n")}
			if got := attribAfterImage(d); got != tc.want {
				t.Errorf("attribAfterImage = %d, want %d", got, tc.want)
			}
		})
	}
}
