package main

import (
	"github.com/ek220/excesspkgname/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.New())
}
