package elem

type Element interface {
	Type() string
	Name() string
	Doc() string
	IsArray() bool
	Count() int64
	Hash() int64
}

// Elem is base type for all elements. It encapsulates all common fields.
// It is public as it is used by go-fbdl internally. It should not be
// used by external code importing go-fbdl.
type Elem struct {
	name    string
	doc     string
	isArray bool
	count   int64
}

func (e Elem) Name() string  { return e.name }
func (e Elem) Doc() string   { return e.doc }
func (e Elem) IsArray() bool { return e.isArray }
func (e Elem) Count() int64  { return e.count }

func MakeElem(name, doc string, isArray bool, count int64) Elem {
	return Elem{
		name: name, doc: doc, isArray: isArray, count: count,
	}
}
