package plugin

import (
	"github.com/kirillveshnyakov/LogLintGo/internal/loglinter"
	"golang.org/x/tools/go/analysis"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		loglinter.Analyzer,
	}, nil
}
