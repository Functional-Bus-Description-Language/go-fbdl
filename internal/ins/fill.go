package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

// fillProps fills required properties that have not been set by the user.
// Some properties have default values and user is not obliged to set them explicitly.
func fillProps(e *Element) {
	switch e.Type {
	case "block":
		fillPropsBlock(e)
	case "bus":
		fillPropsBus(e)
	case "config", "status":
		fillPropsConfig(e)
	case "func":
		fillPropsFunc(e)
	case "mask":
		fillPropsMask(e)
	case "param":
		fillPropsParam(e)
	case "return":
		fillPropsReturn(e)
	default:
		msg := `no implementation for base type '%s'`
		msg = fmt.Sprintf(msg, e.Type)
		panic(msg)
	}
}

func fillPropsBlock(b *Element) {
	if _, ok := b.Props["masters"]; !ok {
		b.Props["masters"] = val.Int(1)
	}
}

func fillPropsBus(b *Element) {
	if _, ok := b.Props["masters"]; !ok {
		b.Props["masters"] = val.Int(1)
	}

	if _, ok := b.Props["width"]; !ok {
		b.Props["width"] = val.Int(int64(busWidth))
	}
}

func fillPropsConfig(c *Element) {
	if _, ok := c.Props["width"]; !ok {
		c.Props["width"] = val.Int(int64(busWidth))
	}

	if _, ok := c.Props["atomic"]; !ok {
		c.Props["atomic"] = val.Bool(true)
	}
}

func fillPropsFunc(f *Element) {
	return
}

func fillPropsMask(m *Element) {
	if _, ok := m.Props["width"]; !ok {
		m.Props["width"] = val.Int(int64(busWidth))
	}

	if _, ok := m.Props["atomic"]; !ok {
		m.Props["atomic"] = val.Bool(true)
	}
}

func fillPropsParam(p *Element) {
	if _, ok := p.Props["width"]; !ok {
		p.Props["width"] = val.Int(int64(busWidth))
	}
}

func fillPropsReturn(p *Element) {
	if _, ok := p.Props["width"]; !ok {
		p.Props["width"] = val.Int(int64(busWidth))
	}
}
