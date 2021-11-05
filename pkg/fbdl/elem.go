package fbdl

type Block struct {
	Name      string
	Count     int64
	Sizes     Sizes
	AddrSpace AddrSpace

	// Properties
	Doc string

	// Elements
	Blocks []Block
	//Configs  []Config
	//Funcs    []Func
	//Masks    []Mask
	Statuses []Status
}

func (b *Block) addStatus(s Status) {
	b.Statuses = append(b.Statuses, s)
}

func (b *Block) hasElement(name string) bool {
	for i, _ := range b.Statuses {
		if b.Statuses[i].Name == name {
			return true
		}
	}

	return false
}

func (b *Block) IsArray() bool {
	return b.AddrSpace.IsArray()
}

type Status struct {
	Name   string
	Count  int64
	Access Access

	// Properties
	Atomic  bool
	Default string
	Doc     string
	Groups  []string
	Once    bool
	Width   int64
}

func (s *Status) IsArray() bool {
	return s.Access.IsArray()
}
