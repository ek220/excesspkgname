package analyzer

import (
	"flag"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func New() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "excesspkgname",
		Doc:      "A linter that checks excess package name",
		Run:      run,
		Flags:    flags(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func flags() flag.FlagSet {
	flags := flag.NewFlagSet("", flag.ContinueOnError)
	return *flags
}

func run(pass *analysis.Pass) (any, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{(*ast.ImportSpec)(nil)}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		importSpec, ok := node.(*ast.ImportSpec)
		if !ok {
			return
		}

		if importSpec.Name == nil {
			return
		}

		name := importSpec.Name.Name
		if name == "_" {
			return
		}

		path := strings.Trim(importSpec.Path.Value, "\"")
		pathElements := strings.Split(path, "/")
		pathSuffix := pathElements[len(pathElements)-1]

		if name != pathSuffix {
			return
		}

		pass.Reportf(importSpec.Pos(), `excess package name: %s "%s"`, name, path)
	})

	return nil, nil
}
