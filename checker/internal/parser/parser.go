package parser

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Directive struct {
	Observed string
	Status   string
	Text     string
	LineNo   int
}

type Result struct {
	Directives []Directive
	Preamble   string
	CharCount  int
	Errors     []string
}

var metaRe = regexp.MustCompile(`<!--\s*observed:\s*([^|>]+?)\s*\|\s*status:\s*(\w+)\s*-->`)

func Parse(raw string) Result {
	r := Result{CharCount: len([]rune(raw))}
	lines := strings.Split(raw, "\n")
	inComment := false
	var current *Directive
	flush := func() {
		if current != nil {
			r.Directives = append(r.Directives, *current)
			current = nil
		}
	}
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if inComment {
			if strings.Contains(trimmed, "-->") {
				inComment = false
			}
			continue
		}
		if strings.HasPrefix(trimmed, "<!--") {
			flush()
			if strings.Contains(trimmed, "-->") {
				if m := metaRe.FindStringSubmatch(trimmed); m != nil {
					current = &Directive{
						Observed: strings.TrimSpace(m[1]),
						Status:   strings.TrimSpace(m[2]),
						LineNo:   i + 1,
					}
				}
			} else {
				inComment = true
			}
			continue
		}
		if current != nil {
			if trimmed != "" {
				if current.Text != "" {
					current.Text += "\n"
				}
				current.Text += trimmed
			}
		} else if trimmed != "" {
			if r.Preamble != "" {
				r.Preamble += "\n"
			}
			r.Preamble += trimmed
		}
	}
	flush()

	for idx := range r.Directives {
		d := &r.Directives[idx]
		if d.Observed == "" {
			r.Errors = append(r.Errors, fmt.Sprintf("line %d: missing observed date", d.LineNo))
		} else if _, err := time.Parse("2006-01-02", d.Observed); err != nil {
			r.Errors = append(r.Errors, fmt.Sprintf("line %d: invalid observed date %q (want YYYY-MM-DD)", d.LineNo, d.Observed))
		}
		if d.Status != "active" && d.Status != "superseded" {
			r.Errors = append(r.Errors, fmt.Sprintf("line %d: invalid status %q (want active or superseded)", d.LineNo, d.Status))
		}
		if strings.TrimSpace(d.Text) == "" {
			r.Errors = append(r.Errors, fmt.Sprintf("line %d: directive has no body", d.LineNo))
		}
	}
	return r
}
