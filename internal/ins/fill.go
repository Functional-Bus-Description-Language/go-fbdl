package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

// fillProperties fills required properties that have not been set by the user.
// Some properties have default values and user is not obliged to set them explicitly.
func fillProperties(e *Element) {
	switch e.Type {
	case "block":
		fillPropertiesBlock(e)
	case "bus":
		fillPropertiesBus(e)
	case "config", "status":
		fillPropertiesConfig(e)
	case "func":
		fillPropertiesFunc(e)
	case "mask":
		fillPropertiesMask(e)
	case "param":
		fillPropertiesParam(e)
	default:
		msg := `no implementation for base type '%s'`
		msg = fmt.Sprintf(msg, e.Type)
		panic(msg)
	}
}

func fillPropertiesBlock(b *Element) {
	if _, ok := b.Props["masters"]; !ok {
		b.Props["masters"] = val.Int(1)
	}
}

func fillPropertiesBus(b *Element) {
	if _, ok := b.Props["masters"]; !ok {
		b.Props["masters"] = val.Int(1)
	}

	if _, ok := b.Props["width"]; !ok {
		b.Props["width"] = val.Int(int64(busWidth))
	}
}

func fillPropertiesConfig(c *Element) {
	if _, ok := c.Props["width"]; !ok {
		c.Props["width"] = val.Int(int64(busWidth))
	}

	if _, ok := c.Props["atomic"]; !ok {
		v := false
		if c.Props["width"].(val.Int) > val.Int(busWidth) {
			v = true
		}
		c.Props["atomic"] = val.Bool(v)
	}
}

func fillPropertiesFunc(f *Element) {
	return
}

func fillPropertiesMask(m *Element) {
	if _, ok := m.Props["width"]; !ok {
		m.Props["width"] = val.Int(int64(busWidth))
	}

	if _, ok := m.Props["atomic"]; !ok {
		v := false
		if m.Props["width"].(val.Int) > val.Int(busWidth) {
			v = true
		}
		m.Props["atomic"] = val.Bool(v)
	}
}

func fillPropertiesParam(p *Element) {
	if _, ok := p.Props["width"]; !ok {
		p.Props["width"] = val.Int(int64(busWidth))
	}
}
