package variables

import (
	"bytes"
	"go/parser"
	"go/printer"
	"go/token"
)

// RenameVariable renames a variable in a file
func RenameVariable(filename, from, to string) ([]byte, error) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, filename, nil, 0)
	if err != nil {
		return nil, err
	}

	// TODO: Exercise

	var buf bytes.Buffer
	err = printer.Fprint(&buf, fs, f)

	return buf.Bytes(), err
}
