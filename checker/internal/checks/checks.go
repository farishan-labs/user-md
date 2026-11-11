package checks

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/farishan-labs/user-md/checker/internal/parser"
)

type Severity string

const (
	SevError Severity = "error"
	SevWarn  Severity = "warn"
	SevInfo  Severity = "info"
	SevPass  Severity = "pass"
)

type Finding struct {
	Check    string   `json:"check"`
	Severity Severity `json:"severity"`
	Message  string   `json:"message"`
	LineNo   int      `json:"line,omitempty"`
}

func truncate(s string, n int) string {
	s = strings.ReplaceAll(s, "\n", " ")
	if len([]rune(s)) <= n {
		return s
	}
	r := []rune(s)
	return string(r[:n]) + "…"
}

// --- parse ---

func ParseCheck(p parser.Result) []Finding {
	if len(p.Directives) == 0 {
		if len(p.Errors) > 0 {
			out := []Finding{}
			for _, e := range p.Errors {
				out = append(out, Finding{Check: "parse", Severity: SevError, Message: e})
			}
			return out
		}
		return []Finding{{Check: "parse", Severity: SevWarn, Message: "no directives found"}}
	}
	if len(p.Errors) > 0 {
		out := []Finding{}
		for _, e := range p.Errors {
			out = append(out, Finding{Check: "parse", Severity: SevError, Message: e})
		}
		return out
	}
	active, superseded := 0, 0
	for _, d := range p.Directives {
		if d.Status == "active" {
			active++
		} else {
			superseded++
		}
	}
	return []Finding{{Check: "parse", Severity: SevPass, Message: fmt.Sprintf("%d directives (%d active, %d superseded), all metadata valid", len(p.Directives), active, superseded)}}
}

// --- cap ---

func Cap(charCount, budget int) []Finding {
	n := charCount
	pct := float64(n) / float64(budget) * 100
	if n > budget {
		return []Finding{{Check: "cap", Severity: SevError, Message: fmt.Sprintf("%d chars exceeds budget %d (%.0f%%)", n, budget, pct)}}
	}
	if pct >= 80 {
		return []Finding{{Check: "cap", Severity: SevWarn, Message: fmt.Sprintf("%d chars = %.0f%% of budget %d (consolidate above 80%%)", n, pct, budget)}}
	}
	return []Finding{{Check: "cap", Severity: SevPass, Message: fmt.Sprintf("%d chars = %.0f%% of budget %d", n, pct, budget)}}
}

// --- secrets ---

type secretPattern struct {
	name string
	re   *regexp.Regexp
}

var secretPatterns = []secretPattern{
	{"AWS access key", regexp.MustCompile(`AKIA[0-9A-Z]{16}`)},
	{"GitHub token", regexp.MustCompile(`gh[pousr]_[A-Za-z0-9]{36,}`)},
	{"GitHub PAT", regexp.MustCompile(`github_pat_[A-Za-z0-9_]{82}`)},
	{"private key", regexp.MustCompile(`-----BEGIN (RSA |EC |OPENSSH |DSA |PGP )?PRIVATE KEY-----`)},
	{"JWT", regexp.MustCompile(`eyJ[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+`)},
	{"Slack token", regexp.MustCompile(`xox[baprs]-[A-Za-z0-9-]+`)},
	{"Stripe key", regexp.MustCompile(`sk_(live|test)_[A-Za-z0-9]{24,}`)},
	{"email", regexp.MustCompile(`[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}`)},
	{"generic secret assignment", regexp.MustCompile(`(?i)(api[_-]?key|token|secret|password|passwd|pwd)\s*[:=]\s*['"]?[A-Za-z0-9_\-]{16,}`)},
}

func Secrets(raw string) []Finding {
	lines := strings.Split(raw, "\n")
	var out []Finding
	for i, line := range lines {
		for _, sp := range secretPatterns {
			if sp.re.MatchString(line) {
				out = append(out, Finding{Check: "secrets", Severity: SevError, Message: fmt.Sprintf("potential %s detected on line %d", sp.name, i+1), LineNo: i + 1})
			}
		}
	}
	if len(out) == 0 {
		out = append(out, Finding{Check: "secrets", Severity: SevPass, Message: "no secrets detected"})
	}
	return out
}

// --- contradictions ---

var stopwords = map[string]bool{
	"prefer": true, "always": true, "never": true, "avoid": true,
	"the": true, "a": true, "an": true, "to": true, "of": true, "in": true,
	"on": true, "for": true, "with": true, "and": true, "or": true, "is": true,
	"are": true, "be": true, "it": true, "you": true, "i": true, "my": true,
	"your": true, "use": true, "using": true, "do": true, "not": true,
}

func contentTokens(s string) []string {
	low := strings.ToLower(s)
	var toks []string
	var cur strings.Builder
	for _, r := range low {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			cur.WriteRune(r)
		} else {
			if cur.Len() >= 3 {
				w := cur.String()
				if !stopwords[w] {
					toks = append(toks, w)
				}
			}
			cur.Reset()
		}
	}
	if cur.Len() >= 3 {
		w := cur.String()
		if !stopwords[w] {
			toks = append(toks, w)
		}
	}
	return toks
}

func hasNegation(s string) bool {
	low := strings.ToLower(s)
	for _, w := range []string{" not ", " never ", " avoid ", " don't ", " no ", " shouldn't ", " must not "} {
		if strings.Contains(low, w) {
			return true
		}
	}
	// leading verb check
	trim := strings.TrimSpace(low)
	for _, p := range []string{"never ", "avoid ", "don't "} {
		if strings.HasPrefix(trim, p) {
			return true
		}
	}
	return false
}

func normalize(s string) string {
	return strings.Join(contentTokens(s), " ")
}

func intersectCount(a, b []string) int {
	set := map[string]bool{}
	for _, t := range a {
		set[t] = true
	}
	n := 0
	seen := map[string]bool{}
	for _, t := range b {
		if set[t] && !seen[t] {
			seen[t] = true
			n++
		}
	}
	return n
}

func Contradictions(dirs []parser.Directive) []Finding {
	var actives []parser.Directive
	for _, d := range dirs {
		if d.Status == "active" {
			actives = append(actives, d)
		}
	}
	var out []Finding

	// exact duplicates
	seen := map[string]int{}
	for _, d := range actives {
		key := normalize(d.Text)
		if key == "" {
			continue
		}
		if prev, ok := seen[key]; ok {
			out = append(out, Finding{Check: "contradictions", Severity: SevWarn, Message: fmt.Sprintf("duplicate active directive (lines %d and %d): %s", prev, d.LineNo, truncate(d.Text, 60)), LineNo: d.LineNo})
		} else {
			seen[key] = d.LineNo
		}
	}

	// potential conflicts: negation polarity
	for i := 0; i < len(actives); i++ {
		for j := i + 1; j < len(actives); j++ {
			a, b := actives[i], actives[j]
			overlap := intersectCount(contentTokens(a.Text), contentTokens(b.Text))
			if hasNegation(a.Text) != hasNegation(b.Text) && overlap >= 1 {
				out = append(out, Finding{Check: "contradictions", Severity: SevWarn, Message: fmt.Sprintf("potential conflict between lines %d and %d (shared subject, opposite polarity): %q vs %q", a.LineNo, b.LineNo, truncate(a.Text, 50), truncate(b.Text, 50)), LineNo: a.LineNo})
			}
		}
	}

	// antonym pairs
	antonyms := [][2]string{
		{"concise", "detailed"}, {"terse", "verbose"}, {"brief", "thorough"},
		{"short", "long"}, {"minimal", "comprehensive"}, {"always", "never"},
		{"enable", "disable"}, {"allow", "block"}, {"fast", "slow"},
	}
	for i := 0; i < len(actives); i++ {
		for j := i + 1; j < len(actives); j++ {
			a, b := actives[i], actives[j]
			la, lb := strings.ToLower(a.Text), strings.ToLower(b.Text)
			for _, pair := range antonyms {
				if (strings.Contains(la, pair[0]) && strings.Contains(lb, pair[1])) ||
					(strings.Contains(la, pair[1]) && strings.Contains(lb, pair[0])) {
					if intersectCount(contentTokens(a.Text), contentTokens(b.Text)) >= 1 {
						out = append(out, Finding{Check: "contradictions", Severity: SevWarn, Message: fmt.Sprintf("potential conflict between lines %d and %d (antonym pair %s/%s): %q vs %q", a.LineNo, b.LineNo, pair[0], pair[1], truncate(a.Text, 50), truncate(b.Text, 50)), LineNo: a.LineNo})
					}
				}
			}
		}
	}

	// dedupe by message
	seenMsg := map[string]bool{}
	deduped := []Finding{}
	for _, f := range out {
		if seenMsg[f.Message] {
			continue
		}
		seenMsg[f.Message] = true
		deduped = append(deduped, f)
	}
	out = deduped

	if len(out) == 0 {
		out = append(out, Finding{Check: "contradictions", Severity: SevPass, Message: "no contradictions detected"})
	}
	return out
}

// --- scope mixing ---

var projectKeywords = []string{
	"project", "repo", "repository", "codebase", "module", "endpoint", "database",
	"schema", "migration", "deploy", "ci", "build", "test", "lint", "function",
	"class", "api", "url", "port", "server", "client", "framework", "library",
	"package", "dependency", "config", "environment",
}

var decisionKeywords = []string{
	"decided", "chose", "will use", "we use", "standard is", "convention is", "policy is",
}

func ScopeMixing(dirs []parser.Directive) []Finding {
	var out []Finding
	flagged := false
	for _, d := range dirs {
		if d.Status != "active" {
			continue
		}
		low := strings.ToLower(d.Text)
		tokens := contentTokens(d.Text)
		tokenSet := map[string]bool{}
		for _, t := range tokens {
			tokenSet[t] = true
		}
		var reasons []string
		for _, k := range projectKeywords {
			if strings.Contains(k, " ") {
				if strings.Contains(low, k) {
					reasons = append(reasons, "project fact ("+k+")")
					break
				}
				continue
			}
			re := regexp.MustCompile(`(?i)\b` + regexp.QuoteMeta(k) + `\b`)
			if re.MatchString(d.Text) {
				reasons = append(reasons, "project fact ("+k+")")
				break
			}
		}
		for _, k := range decisionKeywords {
			if strings.Contains(low, k) {
				reasons = append(reasons, "durable decision")
				break
			}
		}
		if len(reasons) > 0 {
			flagged = true
			out = append(out, Finding{Check: "scope", Severity: SevWarn, Message: fmt.Sprintf("line %d looks like %s — consider AGENTS.md/MEMORY.md: %s", d.LineNo, strings.Join(reasons, ", "), truncate(d.Text, 60)), LineNo: d.LineNo})
		}
	}
	if !flagged {
		out = append(out, Finding{Check: "scope", Severity: SevPass, Message: "no scope mixing detected"})
	}
	// stable order by line
	sort.Slice(out, func(i, j int) bool { return out[i].LineNo < out[j].LineNo })
	return out
}

// --- compat ---

type harness struct {
	name string
	cap  int
}

func Compat(charCount int) []Finding {
	harnesses := []harness{
		{"OpenClaw", 4000},
		{"Hermes", 1375},
		{"Claude Code (global)", 0},
		{"OptMem", 0},
	}
	var parts []string
	for _, h := range harnesses {
		if h.cap == 0 {
			parts = append(parts, fmt.Sprintf("%s: OK (no hard cap)", h.name))
		} else if charCount <= h.cap {
			parts = append(parts, fmt.Sprintf("%s: OK (%d cap)", h.name, h.cap))
		} else {
			parts = append(parts, fmt.Sprintf("%s: TOO BIG (%d > %d)", h.name, charCount, h.cap))
		}
	}
	return []Finding{{Check: "compat", Severity: SevInfo, Message: strings.Join(parts, " | ")}}
}
