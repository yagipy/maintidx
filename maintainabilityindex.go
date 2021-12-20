package maintainabilityindex

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"maintainabilityindex/pkg/maintidx"
)

const doc = "maintainabilityindex measures the maintainability index of each function."

var Analyzer = &analysis.Analyzer{
	Name: "maintainabilityindex",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var under int

func init() {
	Analyzer.Flags.IntVar(&under, "under", 20, "show functions with maintainability < N only.")
}

func run(pass *analysis.Pass) (interface{}, error) {
	i := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	i.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			v := maintidx.Analyze(n)

			v.Coef.Cyc.Calc()
			v.Coef.HalstVol.Calc()
			v.Calc(loc(pass.Fset, n))
			if v.MaintIdx < under {
				pass.Reportf(n.Pos(), "Function name: %v, Cyclomatic Complexity: %v, Halstead Volume: %0.2f, Maintainability Index: %v", n.Name, v.Coef.Cyc.Val, v.Coef.HalstVol.Val, v.MaintIdx)
			}
		}
	})

	return nil, nil
}

func loc(fs *token.FileSet, n *ast.FuncDecl) int {
	f := fs.File(n.Pos())
	startLine := f.Line(n.Pos())
	endLine := f.Line(n.End())
	return endLine - startLine + 1
}
