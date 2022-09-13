package ins

import (
	"math"
	"time"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"

	fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

// If zero is true, then the timestamp will eqaul zero.
// If zero is false, then the timestamp  will be the bus generation timestamp.
func timestamp(zero bool) *elem.Status {
	ts := elem.Status{}

	ts.SetName("TIMESTAMP")
	ts.SetDoc("Bus generation timestamp.")
	ts.SetIsArray(false)
	ts.SetCount(1)
	ts.SetAtomic(true)

	width := busWidth
	// Limit timestamp width. 36 bits is enough, do not waste resources.
	if width > 36 {
		width = 36
	}

	ts.SetWidth(width)

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

	ts.SetDefault(fbdlVal.MakeBitStr(dflt))

	return &ts
}

// Value generation is not yet supported.
func id() *elem.Status {
	id := elem.Status{}

	id.SetName("ID")
	id.SetDoc("Bus identifier.")
	id.SetIsArray(false)
	id.SetCount(1)
	id.SetAtomic(true)

	width := busWidth
	// Current implementaiton uses adler32 for hash, no sense to make ID wider.
	if width > 32 {
		width = 32
	}

	id.SetWidth(width)

	return &id
}
