package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/stream"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type streamDiary struct {
	delaySet bool
}

func insStream(typeChain []prs.Functionality) (*fn.Stream, error) {
	f, err := makeFunctionality(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	stream := fn.Stream{}
	stream.Func = f

	diary := streamDiary{}

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyStreamType(&stream, typ, &diary)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	return &stream, nil
}

func applyStreamType(strm *fn.Stream, typ prs.Functionality, diary *streamDiary) error {
	for _, prop := range typ.Props() {
		if err := util.IsValidProperty(prop.Name, "stream"); err != nil {
			return fmt.Errorf(": %v", err)
		}
		if err := checkProp(prop); err != nil {
			return fmt.Errorf("%s: %v", prop.Loc(), err)
		}

		v, err := prop.Value.Eval()
		if err != nil {
			return fmt.Errorf("cannot evaluate expression")
		}

		switch prop.Name {
		case "delay":
			if diary.delaySet {
				return fmt.Errorf(propAlreadySetMsg, prop.Loc(), "delay")
			}
			t := v.(val.Time)
			delay := fbdlVal.Time{S: t.S, Ns: t.Ns}

			strm.Delay = &delay
			diary.delaySet = true
		default:
			panic("should never happen")
		}

	}

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
