package ins

import (
	"fmt"
	"log"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func insBlackbox(typeChain []prs.Functionality) (*fn.Blackbox, error) {
	typeChainStr := fmt.Sprintf("debug: instantiating blackbox, type chain: %s", typeChain[0].Name())
	for i := 1; i < len(typeChain); i++ {
		typeChainStr = fmt.Sprintf("%s -> %s", typeChainStr, typeChain[i].Name())
	}
	log.Print(typeChainStr)

	f, err := makeFunctionality(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	bb := fn.Blackbox{}
	bb.Func = f

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyBlackboxType(&bb, typ)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	if bb.Size == 0 {
		last := typeChain[len(typeChain)-1]
		return &bb, tok.Error{
			Msg:  fmt.Sprintf("'%s' of type 'blackbox' must have 'size' property set", last.Name()),
			Toks: []tok.Token{last.Tok()},
		}
	}

	return &bb, nil
}

func applyBlackboxType(bb *fn.Blackbox, typ prs.Functionality) error {
	for _, p := range typ.Props() {
		if err := util.IsValidProperty(p.Name, "blackbox"); err != nil {
			return fmt.Errorf(": %v", err)
		}
		if err := checkProp(p); err != nil {
			return fmt.Errorf("%s: %v", p.Loc(), err)
		}

		v, err := p.Value.Eval()
		if err != nil {
			return err
		}

		switch p.Name {
		case "size":
			if bb.Size != 0 {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "masters")
			}
			bb.Size = int64(v.(val.Int))
		default:
			panic(fmt.Sprintf("unhandled '%s' property", p.Name))
		}
	}

	return nil
}
