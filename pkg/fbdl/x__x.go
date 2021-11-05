package fbdl

import (
	_ "math"
	_ "time"
)

func x_timestamp_x() Status {
	return Status{
		Name:    "x_timestamp_x",
		Access:  MakeAccessSingle(1, busWidth),
		Width:   int64(busWidth),
		Default: "implement me",
		//"default": val.Int(time.Now().Unix() & int64(math.Pow(2, float64(busWidth))-1)),
	}
}

// Value generation is not yet supported.
func x_uuid_x() Status {
	return Status{
		Name:    "x_uuid_x",
		Access:  MakeAccessSingle(0, busWidth),
		Width:   int64(busWidth),
		Default: "implement me",
		//"default": val.Int(time.Now().Unix() & int64(math.Pow(2, float64(busWidth))-1)),
	}
}
