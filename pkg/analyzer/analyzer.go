package analyzer

import (
	"github.com/canmor/go-clean-arch/pkg/cleanarch"

	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "gocleanarch",
	Doc:      "Checks that module dependencies for Clean Architecture.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (result interface{}, err error) {
	checker := cleanarch.NewImportChecker(pass.Pkg.Path())
	if checker == nil {
		return
	}

	nodeFilter := []ast.Node{(*ast.ImportSpec)(nil)}
	pass.ResultOf[inspect.Analyzer].(*inspector.Inspector).Preorder(nodeFilter, func(node ast.Node) {
		importSpec := node.(*ast.ImportSpec)
		err := checker.Check(importSpec.Path.Value)
		if err != nil {
			pass.Reportf(importSpec.Pos(), "%s", err)
		}
	})
	return
}