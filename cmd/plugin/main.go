package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/kirillveshnyakov/LogLintGo/internal/loglinter"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("loglinter", New)
}

func New(settings any) (register.LinterPlugin, error) {
	return &plugin{}, nil
}

type plugin struct{}

var _ register.LinterPlugin = new(plugin)

func (*plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		loglinter.Analyzer,
	}, nil
}

func (*plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
