package loglinter

import (
	"go/ast"
	"go/token"
	"go/types"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
)

type LogInfo struct {
	Msg  string
	Expr ast.Expr
}

func isLoggerMethod(method string) bool {
	switch method {
	case "Debug", "Info", "Warn", "Error", "Fatal", "Panic",
		"Debugw", "Infow", "Warnw", "Errorw", "Fatalw", "Panicw":
		return true
	}
	return false
}

func isLoggerType(pass *analysis.Pass, x ast.Expr) bool {
	if pass.TypesInfo == nil {
		return false
	}

	if ident, ok := x.(*ast.Ident); ok {
		obj := pass.TypesInfo.ObjectOf(ident)
		if obj == nil {
			return false
		}
		if pkgName, ok := obj.(*types.PkgName); ok {
			return pkgName.Imported().Path() == "log/slog"
		}
	}

	t := pass.TypesInfo.TypeOf(x)
	if t == nil {
		return false
	}
	if ptr, ok := t.(*types.Pointer); ok {
		t = ptr.Elem()
	}

	typStr := types.TypeString(t, nil)
	return strings.Contains(typStr, "go.uber.org/zap")
}

func checkCall(pass *analysis.Pass, call *ast.CallExpr) (LogInfo, bool) {
	sel, isSel := call.Fun.(*ast.SelectorExpr)
	if !isSel {
		return LogInfo{}, false
	}

	methodName := sel.Sel.Name
	if !isLoggerMethod(methodName) {
		return LogInfo{}, false
	}

	if !isLoggerType(pass, sel.X) {
		return LogInfo{}, false
	}

	if len(call.Args) == 0 {
		return LogInfo{}, false
	}

	msgExpr := call.Args[0]
	msg, ok := getStringMessage(pass, msgExpr)
	if !ok {
		return LogInfo{}, false
	}
	return LogInfo{msg, msgExpr}, true
}

func getStringMessage(pass *analysis.Pass, expr ast.Expr) (string, bool) {
	switch e := expr.(type) {

	case *ast.BasicLit:
		if e.Kind != token.STRING {
			return "", false
		}
		val, err := strconv.Unquote(e.Value)
		if err != nil {
			return "", false
		}
		return val, true

	case *ast.BinaryExpr:
		if e.Op != token.ADD {
			return "", false
		}
		left, okX := getStringMessage(pass, e.X)
		right, okY := getStringMessage(pass, e.Y)
		if okX || okY {
			return left + right, true
		}
		return "", false

	case *ast.Ident:
		if pass.TypesInfo != nil {
			obj := pass.TypesInfo.ObjectOf(e)
			if c, ok := obj.(*types.Const); ok {
				val, err := strconv.Unquote(c.Val().String())
				if err == nil {
					return val, true
				}
			}
		}
		return "", false

	case *ast.ParenExpr:
		return getStringMessage(pass, e.X)

	default:
		return "", false
	}
}
