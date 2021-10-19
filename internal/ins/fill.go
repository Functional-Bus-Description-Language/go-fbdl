package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/val"
)

// fillProperties fills required properties that have not been set by the user.
// Some properties have default values and user is not obliged to set them explicitly.
func fillProperties(e *Element) {
	switch e.BaseType {
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
		msg = fmt.Sprintf(msg, e.BaseType)
		panic(msg)
	}
}

func fillPropertiesBlock(b *Element) {
	if _, ok := b.Properties["masters"]; !ok {
		b.Properties["masters"] = val.Int{V: 1}
	}
}

func fillPropertiesBus(b *Element) {
	if _, ok := b.Properties["masters"]; !ok {
		b.Properties["masters"] = val.Int{V: 1}
	}

	if _, ok := b.Properties["width"]; !ok {
		b.Properties["width"] = val.Int{V: int32(busWidth)}
	}
}

func fillPropertiesConfig(c *Element) {
	if _, ok := c.Properties["width"]; !ok {
		c.Properties["width"] = val.Int{V: int32(busWidth)}
	}

	if _, ok := c.Properties["atomic"]; !ok {
		v := false
		if c.Properties["width"].(val.Int).V > int32(busWidth) {
			v = true
		}
		c.Properties["atomic"] = val.Bool{V: v}
	}
}

func fillPropertiesFunc(f *Element) {
	return
}

func fillPropertiesMask(m *Element) {
	if _, ok := m.Properties["width"]; !ok {
		m.Properties["width"] = val.Int{V: int32(busWidth)}
	}

	if _, ok := m.Properties["atomic"]; !ok {
		v := false
		if m.Properties["width"].(val.Int).V > int32(busWidth) {
			v = true
		}
		m.Properties["atomic"] = val.Bool{V: v}
	}
}

func fillPropertiesParam(p *Element) {
	if _, ok := p.Properties["width"]; !ok {
		p.Properties["width"] = val.Int{V: int32(busWidth)}
	}
}
