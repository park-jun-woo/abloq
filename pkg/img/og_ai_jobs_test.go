//ff:func feature=image type=generator control=iteration dimension=1
//ff:what ogAIJobs 검증 — 안×count 평탄 전개(순서·샘플 번호 1부터), count<1 보정, 잡별 출력 경로 위임
package img

import (
	"path/filepath"
	"testing"
)

func TestOGAIJobs(t *testing.T) {
	variants := []OGVariant{{Name: "a"}, {Name: "b"}}

	// 2 variants x count 2 = 4 jobs, variant-major order, n starts at 1
	spec := OGAISpec{Slug: "post", DraftDir: "og", Multi: true, Count: 2}
	jobs := ogAIJobs(spec, variants)
	if len(jobs) != 4 {
		t.Fatalf("jobs = %d, want 4", len(jobs))
	}
	want := []struct {
		name string
		n    int
		out  string
	}{
		{"a", 1, filepath.Join("og", "post", "a-1.webp")},
		{"a", 2, filepath.Join("og", "post", "a-2.webp")},
		{"b", 1, filepath.Join("og", "post", "b-1.webp")},
		{"b", 2, filepath.Join("og", "post", "b-2.webp")},
	}
	for i, w := range want {
		if jobs[i].variant.Name != w.name || jobs[i].n != w.n || jobs[i].out != w.out {
			t.Errorf("job[%d] = {%s %d %s}, want %+v", i, jobs[i].variant.Name, jobs[i].n, jobs[i].out, w)
		}
	}

	// count < 1 clamps to 1 sample per variant; single path goes direct
	jobs = ogAIJobs(OGAISpec{Slug: "post", OutDir: "static", Count: 0}, variants[:1])
	if len(jobs) != 1 || jobs[0].out != filepath.Join("static", "post.webp") {
		t.Errorf("clamped jobs = %+v, want 1 direct job", jobs)
	}
}
