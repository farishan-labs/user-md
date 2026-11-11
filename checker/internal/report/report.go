package report

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/farishan-labs/user-md/checker/internal/checks"
)

type Summary struct {
	Errors   int `json:"errors"`
	Warnings int `json:"warnings"`
	Passed   int `json:"passed"`
	Infos    int `json:"infos"`
}

type JSONReport struct {
	File      string           `json:"file"`
	CharCount int              `json:"char_count"`
	Budget    int              `json:"budget"`
	Findings  []checks.Finding `json:"findings"`
	Summary   Summary          `json:"summary"`
}

func Summarize(findings []checks.Finding) Summary {
	s := Summary{}
	for _, f := range findings {
		switch f.Severity {
		case checks.SevError:
			s.Errors++
		case checks.SevWarn:
			s.Warnings++
		case checks.SevPass:
			s.Passed++
		case checks.SevInfo:
			s.Infos++
		}
	}
	return s
}

func PrintHuman(file string, findings []checks.Finding) {
	fmt.Printf("usermd-check %s\n\n", file)
	for _, f := range findings {
		sev := strings.ToUpper(string(f.Severity))
		fmt.Printf("%-5s %-14s %s\n", sev, f.Check, f.Message)
	}
	s := Summarize(findings)
	fmt.Printf("\nResult: %d errors, %d warnings, %d passed, %d info\n", s.Errors, s.Warnings, s.Passed, s.Infos)
}

func PrintJSON(file string, charCount, budget int, findings []checks.Finding) {
	rep := JSONReport{
		File:      file,
		CharCount: charCount,
		Budget:    budget,
		Findings:  findings,
		Summary:   Summarize(findings),
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	_ = enc.Encode(rep)
}
