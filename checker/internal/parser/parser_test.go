package parser

import "testing"

func TestParseBasic(t *testing.T) {
	raw := `<!-- observed: 2026-11-11 | status: active -->

- Prefer concise responses.
`
	r := Parse(raw)
	if len(r.Directives) != 1 {
		t.Fatalf("want 1 directive, got %d", len(r.Directives))
	}
	if r.Directives[0].Observed != "2026-11-11" {
		t.Errorf("observed = %q", r.Directives[0].Observed)
	}
	if r.Directives[0].Status != "active" {
		t.Errorf("status = %q", r.Directives[0].Status)
	}
	if len(r.Errors) != 0 {
		t.Errorf("errors: %v", r.Errors)
	}
}

func TestParseSkipsMultilineComment(t *testing.T) {
	raw := `<!-- big comment
<!-- observed: 2026-10-01 | status: superseded -->
- fake directive
-->

<!-- observed: 2026-11-11 | status: active -->

- Real directive.
`
	r := Parse(raw)
	if len(r.Directives) != 1 {
		t.Fatalf("want 1 directive, got %d: %+v", len(r.Directives), r.Directives)
	}
	if r.Directives[0].Text != "- Real directive." {
		t.Errorf("text = %q", r.Directives[0].Text)
	}
}

func TestParseInvalidDate(t *testing.T) {
	raw := `<!-- observed: not-a-date | status: active -->

- x
`
	r := Parse(raw)
	if len(r.Errors) == 0 {
		t.Fatal("expected error")
	}
}
