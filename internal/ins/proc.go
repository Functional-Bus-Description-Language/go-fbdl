package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/proc"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func insProc(typeChain []prs.Element) (*elem.Proc, error) {
	e, err := makeElem(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	proc := elem.Proc{}
	proc.Elem = e

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyProcType(&proc, typ)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	return &proc, nil
}

func applyProcType(p *elem.Proc, typ prs.Element) error {
	for _, s := range typ.Symbols() {
		pe, ok := s.(*prs.Inst)
		if !ok {
			continue
		}

		e := insElement(pe)

		if !util.IsValidInnerType(elem.Type(e), "proc") {
			return fmt.Errorf(invalidInnerTypeMsg, elem.Name(e), elem.Type(e), "proc")
		}

		if proc.HasElement(p, elem.Name(e)) {
			return fmt.Errorf(elemWithNameAlreadyInstMsg, elem.Name(e))
		}
		addProcInnerElement(p, e)
	}

	return nil
}

func addProcInnerElement(p *elem.Proc, e elem.Element) {
	switch e := e.(type) {
	case (*elem.Param):
		p.Params = append(p.Params, e)
	case (*elem.Return):
		p.Returns = append(p.Returns, e)
	default:
		panic("should never happen")
	}
}
