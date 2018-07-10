package pp_test

import (
	"testing"

	"github.com/msawangwan/pp"
)

func TestPrettyPrintStructs(t *testing.T) {
	var s string

	s = pp.Print(struct{ F string }{"a field"})
	t.Log(s)
	s = pp.Print(struct {
		F  string
		FF string
	}{
		"a field",
		"another field",
	})
	t.Log(s)
	s = pp.Print(struct {
		F   string
		FF  string
		FFF string
		I   int
	}{
		"a field",
		"another field",
		"one more field",
		34,
	})
	t.Log(s)
}

func TestPrettyPrintStructsWithErrorChecking(t *testing.T) {
	var (
		s string
		e error
	)

	s, e = pp.PrintSafe(nil)
	if e != nil {
		t.Log(e)
	} else {
		t.Log(s)
	}

	s, e = pp.PrintSafe(struct{ F string }{"some field"})
	if e != nil {
		t.Log(e)
	} else {
		t.Log(s)
	}
}
