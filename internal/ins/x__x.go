package ins

import (
	"math"
	"time"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

func x_timestamp_x() *Element {
	width := busWidth
	// Limit timestamp width. 36 bits is enough, do not waste resources.
	if width > 36 {
		width = 36
	}

	timestamp := val.Int(time.Now().Unix() & int64(math.Pow(2, float64(width))-1))

	dflt, err := val.BitStrFromInt(timestamp, width)
	if err != nil {
		panic("X_TIMESTAMP_X")
	}

	return &Element{
		Name:  "X_TIMESTAMP_X",
		Doc:   "Bus generation timestamp.",
		Type:  "status",
		Count: int64(1),
		Props: map[string]val.Value{
			"atomic":  val.Bool(true),
			"default": dflt,
			"width":   val.Int(width),
		},
	}
}

// Value generation is not yet supported.
func x_uuid_x() *Element {
	dflt, err := val.BitStrFromInt(val.Int(0xDEADBEEF), busWidth)
	if err != nil {
		panic("X_UUID_X")
	}

	return &Element{
		Name:  "X_UUID_X",
		Doc:   "Bus universally unique identifier.",
		Type:  "status",
		Count: int64(1),
		Props: map[string]val.Value{
			"atomic":  val.Bool(true),
			"default": dflt,
			"width":   val.Int(busWidth),
		},
	}
}
