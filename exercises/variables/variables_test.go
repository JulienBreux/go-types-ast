package variables

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestRenameVariable(t *testing.T) {
	got, err := RenameVariable("files/from.go", "x", "y")
	if err != nil {
		t.Errorf("got no error, wanted %q", err)
	}

	want, _ := ioutil.ReadFile("files/to.go")
	if !bytes.Equal(got, want) {
		t.Errorf("got:\n%s\nwanted:\n%s\n", got, want)
	}
}
