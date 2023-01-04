package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/stream"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func insStream(typeChain []prs.Element) (*elem.Stream, error) {
	e, err := makeElem(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	stream := elem.Stream{}
	stream.Elem = e

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

func applyStreamType(strm *elem.Stream, typ prs.Element) error {
	for _, s := range typ.Symbols() {
		pe, ok := s.(*prs.Inst)
		if !ok {
			continue
		}

		e := insElement(pe)

		if !util.IsValidInnerType(elem.Type(e), "stream") {
			return fmt.Errorf(invalidInnerTypeMsg, elem.Name(e), elem.Type(e), "stream")
		}

		if stream.HasElement(strm, elem.Name(e)) {
			return fmt.Errorf(elemWithNameAlreadyInstMsg, elem.Name(e))
		}

		err := addStreamInnerElement(strm, e)
		if err != nil {
			return fmt.Errorf(
				"line %d: cannot instantiate element '%s': %v",
				pe.LineNum(), elem.Name(e), err,
			)
		}
	}

	return nil
}

func addStreamInnerElement(s *elem.Stream, e elem.Element) error {
	if (elem.Type(e) == "return" && len(s.Params) > 0) ||
		(elem.Type(e) == "param" && len(s.Returns) > 0) {
		return fmt.Errorf(
			"all 'stream' inner elements must be of the same base type and must be 'param' or 'return', "+
				"'%s' base type is '%s'", elem.Name(e), elem.Type(e),
		)
	}

	switch e := e.(type) {
	case (*elem.Param):
		s.Params = append(s.Params, e)
	case (*elem.Return):
		s.Returns = append(s.Returns, e)
	default:
		panic("should never happen")
	}

	return nil
}
