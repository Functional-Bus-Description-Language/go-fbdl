package reg

import (
	"math"
	"time"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"

	fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

func timestamp() *elem.Static {
	ts := elem.Static{}

	ts.Name = "TIMESTAMP"
	ts.Doc = "Bus generation timestamp."
	ts.IsArray = false
	ts.Count = 1

	width := busWidth
	// Limit timestamp width. 36 bits is enough, do not waste resources.
	if width > 36 {
		width = 36
	}

	ts.Width = width

	timestamp := val.Int(time.Now().Unix() & int64(math.Pow(2, float64(width))-1))

	dflt, err := val.BitStrFromInt(timestamp, width)
	if err != nil {
		panic("TIMESTAMP")
	}

	ts.Default = fbdlVal.MakeBitStr(dflt)

	return &ts
}

// Value generation is not yet supported.
func id() *elem.Static {
	id := elem.Static{}

	id.Name = "ID"
	id.Doc = "Bus identifier."
	id.IsArray = false
	id.Count = 1

	width := busWidth
	// Current implementaiton uses adler32 for hash, no sense to make ID wider.
	if width > 32 {
		width = 32
	}

	id.Width = width

	return &id
}
