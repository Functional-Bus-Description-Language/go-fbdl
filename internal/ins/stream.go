package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	fbdl "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func insStream(typeChain []prs.Element) (*elem.Stream, error) {
	e, err := makeElem(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	stream := elem.Stream{}
	stream.SetElem(e)

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

func applyStreamType(stream *elem.Stream, typ prs.Element) error {
	for _, s := range typ.Symbols() {
		pe, ok := s.(*prs.Inst)
		if !ok {
			continue
		}

		e := insElement(pe)

		if !util.IsValidInnerType(e.Type(), "func") {
			return fmt.Errorf(invalidInnerTypeMsg, e.Name(), e.Type(), "func")
		}

		if stream.HasElement(e.Name()) {
			return fmt.Errorf(elemWithNameAlreadyInstMsg, e.Name())
		}

		err := addStreamInnerElement(stream, e)
		if err != nil {
			return fmt.Errorf(
				"line %d: cannot instantiate element '%s': %v",
				pe.LineNum(), e.Name(), err,
			)
		}
	}

	return nil
}

func addStreamInnerElement(stream *elem.Stream, e fbdl.Element) error {
	if (e.Type() == "return" && len(stream.Params()) > 0) ||
		(e.Type() == "param" && len(stream.Returns()) > 0) {
		return fmt.Errorf(
			"all 'stream' inner elements must be of the same base type and must be 'param' or 'return', "+
				"'%s' base type is '%s'", e.Name(), e.Type(),
		)
	}

	switch e.(type) {
	case (*elem.Param):
		stream.AddParam(e.(*elem.Param))
	case (*elem.Return):
		stream.AddReturn(e.(*elem.Return))
	default:
		panic("should never happen")
	}

	return nil
}
