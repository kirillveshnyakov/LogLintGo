package tests

import (
	"testing"

	"github.com/kirillveshnyakov/LogLintGo/internal/loglinter"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestSlogRules(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), loglinter.Analyzer, "slog_tests")
}

func TestZapRules(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), loglinter.Analyzer, "zap_tests")
}
