package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"testing"
)

func TestBuildError(t *testing.T) {
	var tests = []struct {
		idx int // Test index, useful for navigation
		src string
		err error
	}{
		{
			0,
			"import *",
			fmt.Errorf("1:8: unexpected *, expected identifier, string or newline"),
		},
		{
			1,
			"import name +",
			fmt.Errorf("1:13: unexpected +, expected string"),
		},
		{
			2,
			"const A = ]",
			fmt.Errorf("1:11: unexpected ], expected expression"),
		},
	}

	for i, test := range tests {
		if i != test.idx {
			t.Fatalf("Invalid test index %d, expected %d", test.idx, i)
		}

		stream, err := token.Parse([]byte(test.src))
		if err != nil {
			t.Fatalf("%d: token.Parse: %v, expected nil", i, err)
		}

		_, err = Build(stream)
		if err == nil {
			t.Fatalf("%d: err == nil, expected != nil", i)
		}

		if err.Error() != test.err.Error() {
			t.Fatalf("%d:\n got: %v\nwant: %v", i, err, test.err)
		}
	}
}
