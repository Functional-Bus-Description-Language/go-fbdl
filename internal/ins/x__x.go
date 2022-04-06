package ins

import (
	"math"
	"time"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

// If zero is true, then the timestamp will eqaul zero.
// If zero is false, then the timestamp  will be the bus generation timestamp.
func timestamp(zero bool) *Element {
	width := busWidth
	// Limit timestamp width. 36 bits is enough, do not waste resources.
	if width > 36 {
		width = 36
	}

	var timestamp val.Int
	if zero {
		timestamp = val.Int(0)
	} else {
		timestamp = val.Int(time.Now().Unix() & int64(math.Pow(2, float64(width))-1))
	}

	dflt, err := val.BitStrFromInt(timestamp, width)
	if err != nil {
		panic("TIMESTAMP")
	}

	return &Element{
		Name:  "TIMESTAMP",
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
func id() *Element {
	dflt, err := val.BitStrFromInt(val.Int(0xDEADBEEF), busWidth)
	if err != nil {
		panic("ID")
	}

	return &Element{
		Name:  "ID",
		Doc:   "Bus identifier.",
		Type:  "status",
		Count: int64(1),
		Props: map[string]val.Value{
			"atomic":  val.Bool(true),
			"default": dflt,
			"width":   val.Int(busWidth),
		},
	}
}
