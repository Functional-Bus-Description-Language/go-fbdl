package elem

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
