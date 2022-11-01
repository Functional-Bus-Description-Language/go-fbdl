package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/fun"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func insFunc(typeChain []prs.Element) (*elem.Func, error) {
	e, err := makeElem(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	fun := elem.Func{}
	fun.Elem = e

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyFuncType(&fun, typ)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	return &fun, nil
}

func applyFuncType(f *elem.Func, typ prs.Element) error {
	for _, s := range typ.Symbols() {
		pe, ok := s.(*prs.Inst)
		if !ok {
			continue
		}

		e := insElement(pe)

		if !util.IsValidInnerType(elem.Type(e), "func") {
			return fmt.Errorf(invalidInnerTypeMsg, elem.Name(e), elem.Type(e), "func")
		}

		if fun.HasElement(f, elem.Name(e)) {
			return fmt.Errorf(elemWithNameAlreadyInstMsg, elem.Name(e))
		}
		addFuncInnerElement(f, e)
	}

	return nil
}

func addFuncInnerElement(fun *elem.Func, e elem.Element) {
	switch e := e.(type) {
	case (*elem.Param):
		fun.Params = append(fun.Params, e)
	case (*elem.Return):
		fun.Returns = append(fun.Returns, e)
	default:
		panic("should never happen")
	}
}
