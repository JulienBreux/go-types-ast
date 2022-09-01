package variables

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
)

// RenameVariable renames a variable in a file
func RenameVariable(filename, oldVar, newVar string) ([]byte, error) {
	// Parse file into new file set
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// TODO: WIP
	// info := types.Info{
	// 	Types: make(map[ast.Expr]types.TypeAndValue),
	// 	Defs:  make(map[*ast.Ident]types.Object),
	// 	Uses:  make(map[*ast.Ident]types.Object),
	// }
	// var conf types.Config
	// pkg, err = conf.Check("fib", fset, []*ast.File{f}, &info)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// TODO: WIP
	// usesByObj := make(map[types.Object][]string)
	// for id, obj := range info.Uses {
	// 	posn := fset.Position(id.Pos())
	// 	lineCol := fmt.Sprintf("%d:%d", posn.Line, posn.Column)
	// 	usesByObj[obj] = append(usesByObj[obj], lineCol)
	// }

	// Used to print AST, to be removed
	ast.Inspect(f, func(n ast.Node) bool {
		ast.Print(fset, n)
		return true
	})

	// Core of transformation
	for _, ds := range f.Decls {
		switch d := ds.(type) {
		case *ast.GenDecl:
			for _, ss := range d.Specs {
				switch s := ss.(type) {
				case *ast.ValueSpec:
					if s.Names[0].Name == oldVar {
						s.Names[0].Name = newVar
					}
				}
			}
		case *ast.FuncDecl:
			for _, ss := range d.Body.List {
				switch s := ss.(type) {
				case *ast.AssignStmt:
					for _, exps := range s.Rhs {
						switch exp := exps.(type) {
						case *ast.Ident:
							if exp.Name == oldVar {
								exp.Name = newVar
							}
						}
					}
				}
			}
		}
	}

	var buf bytes.Buffer
	err = printer.Fprint(&buf, fset, f)

	return buf.Bytes(), err
}
