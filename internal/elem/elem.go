package elem

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"
)

// elem is base type for all elements.
type elem struct {
	Name    string
	Doc     string
	IsArray bool
	Count   int64
}

type Elem struct {
	elem
}

func (e *Elem) SetName(n string) { e.elem.Name = n }
func (e *Elem) Name() string     { return e.elem.Name }

func (e *Elem) SetDoc(d string) { e.elem.Doc = d }
func (e *Elem) Doc() string     { return e.elem.Doc }

func (e *Elem) SetIsArray(ia bool) { e.elem.IsArray = ia }
func (e *Elem) IsArray() bool      { return e.elem.IsArray }

func (e *Elem) SetCount(c int64) { e.elem.Count = c }
func (e *Elem) Count() int64     { return e.elem.Count }

func (e *Elem) SetElem(el Elem) {
	e.SetName(el.Name())
	e.SetDoc(el.Doc())
	e.SetIsArray(el.IsArray())
	e.SetCount(el.Count())
}

func (e *Elem) Hash() uint32 {
	buf := bytes.Buffer{}

	// Name
	buf.Write([]byte(e.Name()))

	// Doc
	buf.Write([]byte(e.Doc()))

	// IsArray
	if e.IsArray() {
		buf.WriteByte(1)
	} else {
		buf.WriteByte(0)
	}

	// Count
	err := binary.Write(&buf, binary.LittleEndian, e.Count())
	if err != nil {
		panic(err)
	}

	return adler32.Checksum(buf.Bytes())
}
