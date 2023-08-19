package ins

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
)

func typeChainIter(typeChain []prs.Functionality) func() (prs.Functionality, bool) {
	tc := typeChain
	i := 0
	return func() (prs.Functionality, bool) {
		if i == len(tc) {
			return nil, false
		}
		resolvedArgs := make(map[string]prs.Expr)
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
