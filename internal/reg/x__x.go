package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl"

	"math"
	"time"
)

func x_timestamp_x() *FunctionalElement {
	return &FunctionalElement{
		Access: MakeAccessSingle(1, busWidth),
		InsElem: &ins.Element{
			Name:     "x_timestamp_x",
			BaseType: "status",
			IsArray:  false,
			Count:    0,
			Properties: map[string]fbdl.Value{
				"width":   fbdl.Int{V: int64(busWidth)},
				"default": fbdl.Int{V: time.Now().Unix() & int64(math.Pow(2, float64(busWidth))-1)},
			},
			Constants: nil,
			Elements:  nil,
		},
	}
}

// Value generation is not yet supported.
func x_uuid_x() *FunctionalElement {
	return &FunctionalElement{
		Access: MakeAccessSingle(0, busWidth),
		InsElem: &ins.Element{
			Name:     "x_uuid_x",
			BaseType: "status",
			IsArray:  false,
			Count:    0,
			Properties: map[string]fbdl.Value{
				"width":   fbdl.Int{V: int64(busWidth)},
				"default": fbdl.Int{V: 0},
			},
			Constants: nil,
			Elements:  nil,
		},
	}
}
