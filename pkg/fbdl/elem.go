package fbdl

// Block represents block element as well as bus element.
type Block struct {
	Name      string
	IsArray   bool
	Count     int64
	Sizes     Sizes
	AddrSpace AddrSpace

	// Properties
	Doc     string
	Masters int64
	Width   int64

	// Elements
	Subblocks []Block
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

// Status represents status element.
type Status struct {
	Name    string
	IsArray bool
	Count   int64
	Access  Access

	// Properties
	Atomic  bool
	Default string
	Doc     string
	Groups  []string
	Once    bool
	Width   int64
}
