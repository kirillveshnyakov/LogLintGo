package main

import (
	"github.com/kirillveshnyakov/LogLintGo/internal/loglinter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(loglinter.Analyzer)
}
