package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/iface"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
)

func insFunc(typeChain []prs.Element) (*elem.Func, error) {
	e, err := makeElem(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	fun := elem.Func{}
	fun.SetElem(e)

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

func applyFuncType(fun *elem.Func, typ prs.Element) error {
	for _, s := range typ.Symbols() {
		pe, ok := s.(*prs.Inst)
		if !ok {
			continue
		}

		e := insElement(pe)

		if util.IsValidInnerType(e.Type(), "func") == false {
			return fmt.Errorf(invalidInnerTypeMsg, e.Name(), e.Type(), "func")
		}

		if fun.HasElement(e.Name()) {
			return fmt.Errorf(elemWithNameAlreadyInstMsg, e.Name())
		}
		addFuncInnerElement(fun, e)
	}

	return nil
}

func addFuncInnerElement(fun *elem.Func, e iface.Element) {
	switch e.(type) {
	case (*elem.Param):
		fun.AddParam(e.(*elem.Param))
	case (*elem.Return):
		fun.AddReturn(e.(*elem.Return))
	default:
		panic("should never happen")
	}
}
