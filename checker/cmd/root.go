package cmd

import (
	"fmt"
	"os"

	"github.com/farishan-labs/user-md/checker/internal/buildinfo"
	"github.com/farishan-labs/user-md/checker/internal/checks"
	"github.com/farishan-labs/user-md/checker/internal/parser"
	"github.com/farishan-labs/user-md/checker/internal/report"
	"github.com/spf13/cobra"
)

var (
	budget  int
	jsonOut bool
	strict  bool
)

var rootCmd = &cobra.Command{
	Use:     "usermd-check [file]",
	Short:   "Validate a USER.md operator profile",
	Long:    "usermd-check validates a USER.md file: parse, char cap, secrets, contradictions, scope mixing, harness compatibility.",
	Version: buildinfo.Version,
	Args:    cobra.MaximumNArgs(1),
	RunE:    run,
}

func init() {
	rootCmd.Flags().IntVarP(&budget, "budget", "b", 1375, "character budget (default 1375; OpenClaw uses 4000)")
	rootCmd.Flags().BoolVar(&jsonOut, "json", false, "output JSON")
	rootCmd.Flags().BoolVar(&strict, "strict", false, "treat warnings as errors (non-zero exit)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	path := "USER.md"
	if len(args) > 0 {
		path = args[0]
	}

	var raw string
	if path == "-" {
		b, err := os.ReadFile("/dev/stdin")
		if err != nil {
			return err
		}
		raw = string(b)
		path = "<stdin>"
	} else {
		b, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("cannot read %s: %w", path, err)
		}
		raw = string(b)
	}

	p := parser.Parse(raw)

	var findings []checks.Finding
	findings = append(findings, checks.ParseCheck(p)...)
	findings = append(findings, checks.Cap(p.CharCount, budget)...)
	findings = append(findings, checks.Secrets(raw)...)
	findings = append(findings, checks.Contradictions(p.Directives)...)
	findings = append(findings, checks.ScopeMixing(p.Directives)...)
	findings = append(findings, checks.Compat(p.CharCount)...)

	if jsonOut {
		report.PrintJSON(path, p.CharCount, budget, findings)
	} else {
		report.PrintHuman(path, findings)
	}

	// exit code logic
	s := report.Summarize(findings)
	if s.Errors > 0 {
		os.Exit(1)
	}
	if strict && s.Warnings > 0 {
		os.Exit(1)
	}
	return nil
}
