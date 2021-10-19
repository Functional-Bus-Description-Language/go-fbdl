package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/val"
)

func fillMissingProperties(e *Element) {
	switch e.BaseType {
	case "block":
		fillMissingPropertiesBlock(e)
	case "bus":
		fillMissingPropertiesBus(e)
	case "config", "status":
		fillMissingPropertiesConfig(e)
	case "func":
		fillMissingPropertiesFunc(e)
	case "mask":
		fillMissingPropertiesMask(e)
	case "param":
		fillMissingPropertiesParam(e)
	default:
		msg := `no implementation for base type '%s'`
		msg = fmt.Sprintf(msg, e.BaseType)
		panic(msg)
	}
}

func fillMissingPropertiesBlock(b *Element) {
	if _, ok := b.Properties["masters"]; !ok {
		b.Properties["masters"] = val.Int{V: 1}
	}
}

func fillMissingPropertiesBus(b *Element) {
	if _, ok := b.Properties["masters"]; !ok {
		b.Properties["masters"] = val.Int{V: 1}
	}

	if _, ok := b.Properties["width"]; !ok {
		b.Properties["width"] = val.Int{V: int32(busWidth)}
	}
}

func fillMissingPropertiesConfig(c *Element) {
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

func fillMissingPropertiesFunc(f *Element) {
	return
}

func fillMissingPropertiesMask(m *Element) {
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

func fillMissingPropertiesParam(p *Element) {
	if _, ok := p.Properties["width"]; !ok {
		p.Properties["width"] = val.Int{V: int32(busWidth)}
	}
}
