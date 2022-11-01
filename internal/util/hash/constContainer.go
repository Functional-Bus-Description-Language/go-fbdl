package hash

import (
	"bytes"
	"hash/adler32"
	"sort"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
	"golang.org/x/exp/maps"
)

func hashConstContainer(c *elem.ConstContainer) uint32 {
	buf := bytes.Buffer{}

	// BoolConsts
	keys := maps.Keys(c.BoolConsts)
	sort.Strings(keys)
	for _, key := range keys {
		write(&buf, c.BoolConsts[key])
	}

	// BoolListConsts
	keys = maps.Keys(c.BoolListConsts)
	sort.Strings(keys)
	for _, key := range keys {
		list := c.BoolListConsts[key]
		for _, val := range list {
			write(&buf, val)
		}
	}

	// FloatConsts
	keys = maps.Keys(c.FloatConsts)
	sort.Strings(keys)
	for _, key := range keys {
		write(&buf, c.FloatConsts[key])
	}

	// IntConsts
	keys = maps.Keys(c.IntConsts)
	sort.Strings(keys)
	for _, key := range keys {
		write(&buf, c.IntConsts[key])
	}

	// IntListConsts
	keys = maps.Keys(c.IntListConsts)
	sort.Strings(keys)
	for _, key := range keys {
		list := c.IntListConsts[key]
		for _, val := range list {
			write(&buf, val)
		}
	}

	// StrConsts
	keys = maps.Keys(c.StrConsts)
	sort.Strings(keys)
	for _, key := range keys {
		buf.Write([]byte(c.StrConsts[key]))
	}

	return adler32.Checksum(buf.Bytes())
}
