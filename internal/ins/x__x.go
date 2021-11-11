package ins

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"math"
	"time"
)

func x_timestamp_x() *Element {
	timestamp := val.Int(time.Now().Unix() & int64(math.Pow(2, float64(busWidth))-1))

	dflt, err := val.BitStrFromInt(timestamp, busWidth)
	if err != nil {
		panic("x_timestamp_x")
	}

	return &Element{
		Name:     "x_timestamp_x",
		BaseType: "status",
		Count:    int64(1),
		Properties: map[string]val.Value{
			"atomic":  val.Bool(false),
			"default": dflt,
			"width":   val.Int(busWidth),
		},
	}
}

// Value generation is not yet supported.
func x_uuid_x() *Element {
	dflt, err := val.BitStrFromInt(val.Int(0xDEADBEEF), busWidth)
	if err != nil {
		panic("x_uuid_x")
	}

	return &Element{
		Name:     "x_uuid_x",
		BaseType: "status",
		Count:    int64(1),
		Properties: map[string]val.Value{
			"atomic":  val.Bool(false),
			"default": dflt,
			"width":   val.Int(busWidth),
		},
	}
}
