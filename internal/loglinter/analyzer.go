package loglinter

import (
	"go/ast"

	"github.com/kirillveshnyakov/LogLintGo/internal/loglinter/rules"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "loglint",
	Doc:      "checks log messages",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{(*ast.CallExpr)(nil)}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		call := n.(*ast.CallExpr)

		lc, ok := checkCall(pass, call)
		if !ok {
			return
		}

		checkRules(pass, lc)

	})

	return nil, nil
}

func checkRules(pass *analysis.Pass, li LogInfo) {
	if ok := rules.CheckLowercase(li.Msg); ok {
		pass.Reportf(
			li.Expr.Pos(),
			"log message should start with a lowercase letter, got %q",
			li.Msg,
		)
	}
	if ok := rules.CheckEnglish(li.Msg); ok {
		pass.Reportf(
			li.Expr.Pos(),
			"log message must be in English, got %q",
			li.Msg,
		)
	}
	if ok := rules.CheckNoSpecialSymbols(li.Msg); ok {
		pass.Reportf(
			li.Expr.Pos(),
			"log message must not contain special symbols, got %q",
			li.Msg,
		)
	}
	if keyWord, ok := rules.CheckSensitiveWords(li.Msg); ok {
		pass.Reportf(
			li.Expr.Pos(),
			"log message may contain sensitive data (found %q), avoid logging sensitive information",
			keyWord,
		)
	}
}
