package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

type memDiary struct {
	accessSet          bool
	byteWriteEnableSet bool
	readLatencySet     bool
	sizeSet            bool
	widthSet           bool
}

func insMemory(typeChain []prs.Functionality) (*fn.Memory, error) {
	f, err := makeFunctionality(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	mem := fn.Memory{}
	mem.Func = f

	diary := memDiary{}

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyMemoryType(&mem, typ, &diary)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	err = fillMemoryProps(&mem, diary)
	if err != nil {
		last := typeChain[len(typeChain)-1]
		return nil, tok.Error{
			Msg:  fmt.Sprintf("%v", err),
			Toks: []tok.Token{last.Tok()},
		}
	}

	return &mem, nil
}

func applyMemoryType(mem *fn.Memory, typ prs.Functionality, diary *memDiary) error {
	for _, p := range typ.Props() {
		if err := util.IsValidProperty(p.Name, "memory"); err != nil {
			return fmt.Errorf(": %v", err)
		}
		if err := checkProp(p); err != nil {
			return err
		}

		v, err := p.Value.Eval()
		if err != nil {
			return err
		}

		switch p.Name {
		case "access":
			if diary.accessSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "access")
			}
			mem.Access = (string(v.(val.Str)))
			diary.accessSet = true
		case "byte-write-enable":
			if diary.byteWriteEnableSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "byte-write-enable")
			}
			mem.ByteWriteEnable = (bool(v.(val.Bool)))
			diary.byteWriteEnableSet = true
		case "read-latency":
			if diary.readLatencySet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "read-latency")
			}
			mem.ReadLatency = int64(v.(val.Int))
			diary.readLatencySet = true
		case "size":
			if diary.sizeSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "size")
			}
			mem.Size = int64(v.(val.Int))
			diary.sizeSet = true
		case "width":
			if diary.widthSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "width")
			}
			mem.Width = int64(v.(val.Int))
			diary.widthSet = true
		default:
			panic(fmt.Sprintf("unhandled '%s' property", p.Name))
		}
	}

	return nil
}

func fillMemoryProps(mem *fn.Memory, diary memDiary) error {
	if !diary.accessSet {
		mem.Access = "Read Write"
	}

	if !diary.sizeSet {
		return fmt.Errorf("'memory' must have 'size' property set")
	}

	if !diary.readLatencySet && mem.Access != "Write Only" {
		return fmt.Errorf(
			"'memory' must have 'read-latency' property set when its access equals %q", mem.Access,
		)
	}

	if !diary.widthSet {
		mem.Width = busWidth
	}

	return nil
}
