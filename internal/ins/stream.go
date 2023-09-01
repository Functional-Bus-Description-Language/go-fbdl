package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/stream"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func insStream(typeChain []prs.Functionality) (*fn.Stream, error) {
	f, err := makeFunctionality(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	stream := fn.Stream{}
	stream.Func = f

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyStreamType(&stream, typ)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	return &stream, nil
}

func applyStreamType(strm *fn.Stream, typ prs.Functionality) error {
	for _, s := range typ.Symbols() {
		pe, ok := s.(*prs.Inst)
		if !ok {
			continue
		}

		e := insFunctionality(pe)

		if !util.IsValidInnerType(e.Type(), "stream") {
			return fmt.Errorf(invalidInnerTypeMsg, e.GetName(), e.Type(), "stream")
		}

		if stream.HasElement(strm, e.GetName()) {
			return fmt.Errorf(funcWithNameAlreadyInstMsg, e.GetName())
		}

		err := addStreamInnerElement(strm, e)
		if err != nil {
			return fmt.Errorf(
				"%d:%d: cannot instantiate '%s' functionality: %v",
				pe.Line(), pe.Col(), e.GetName(), err,
			)
		}
	}

	return nil
}

func addStreamInnerElement(s *fn.Stream, e fn.Functionality) error {
	if (e.Type() == "return" && len(s.Params) > 0) ||
		(e.Type() == "param" && len(s.Returns) > 0) {
		return fmt.Errorf(
			"all 'stream' inner functionalities must be of the same base type and must be 'param' or 'return', "+
				"'%s' base type is '%s'", e.GetName(), e.Type(),
		)
	}

	switch e := e.(type) {
	case (*fn.Param):
		s.Params = append(s.Params, e)
	case (*fn.Return):
		s.Returns = append(s.Returns, e)
	default:
		panic("should never happen")
	}

	return nil
}
