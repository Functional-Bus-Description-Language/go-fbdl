package ins

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/parser"
)

func typeChainIter(typeChain []parser.Functionality) func() (parser.Functionality, bool) {
	tc := typeChain
	i := 0
	return func() (parser.Functionality, bool) {
		if i == len(tc) {
			return nil, false
		}
		resolvedArgs := make(map[string]parser.Expr)
		if (i+1) < len(tc) && tc[i+1].ResolvedArgs() != nil {
			resolvedArgs = tc[i+1].ResolvedArgs()
		}
		typ := tc[i]
		if resolvedArgs != nil {
			typ.SetResolvedArgs(resolvedArgs)
		}
		i += 1
		return typ, true
	}
}
