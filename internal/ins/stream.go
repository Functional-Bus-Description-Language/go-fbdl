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

		f := insFunctionality(pe)

		if !util.IsValidInnerType(f.Type(), "stream") {
			return fmt.Errorf(invalidInnerTypeMsg, f.GetName(), f.Type(), "stream")
		}

		if stream.HasFunctionality(strm, f.GetName()) {
			return fmt.Errorf(funcWithNameAlreadyInstMsg, f.GetName())
		}

		err := addStreamInnerFunctionality(strm, f)
		if err != nil {
			return fmt.Errorf(
				"%d:%d: cannot instantiate '%s' functionality: %v",
				pe.Line(), pe.Col(), f.GetName(), err,
			)
		}
	}

	return nil
}

func addStreamInnerFunctionality(s *fn.Stream, f fn.Functionality) error {
	if (f.Type() == "return" && len(s.Params) > 0) ||
		(f.Type() == "param" && len(s.Returns) > 0) {
		return fmt.Errorf(
			"all 'stream' inner functionalities must be of the same base type and must be 'param' or 'return', "+
				"'%s' base type is '%s'", f.GetName(), f.Type(),
		)
	}

	switch f := f.(type) {
	case (*fn.Param):
		s.Params = append(s.Params, f)
	case (*fn.Return):
		s.Returns = append(s.Returns, f)
	default:
		panic("should never happen")
	}

	return nil
}
