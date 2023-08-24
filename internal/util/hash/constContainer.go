package hash

import (
	"bytes"
	"hash/adler32"
	"sort"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/cnst"
	"golang.org/x/exp/maps"
)

func hashConstContainer(c *cnst.Container) uint32 {
	buf := bytes.Buffer{}

	// Bools
	keys := maps.Keys(c.Bools)
	sort.Strings(keys)
	for _, key := range keys {
		write(&buf, c.Bools[key])
	}

	// BoolLists
	keys = maps.Keys(c.BoolLists)
	sort.Strings(keys)
	for _, key := range keys {
		list := c.BoolLists[key]
		for _, val := range list {
			write(&buf, val)
		}
	}

	// Floats
	keys = maps.Keys(c.Floats)
	sort.Strings(keys)
	for _, key := range keys {
		write(&buf, c.Floats[key])
	}

	// Ints
	keys = maps.Keys(c.Ints)
	sort.Strings(keys)
	for _, key := range keys {
		write(&buf, c.Ints[key])
	}

	// IntLists
	keys = maps.Keys(c.IntLists)
	sort.Strings(keys)
	for _, key := range keys {
		list := c.IntLists[key]
		for _, val := range list {
			write(&buf, val)
		}
	}

	// StrConsts
	keys = maps.Keys(c.Strings)
	sort.Strings(keys)
	for _, key := range keys {
		buf.Write([]byte(c.Strings[key]))
	}

	return adler32.Checksum(buf.Bytes())
}
