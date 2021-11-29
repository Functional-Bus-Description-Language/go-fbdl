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
	Subblocks []*Block
	Configs   []*Config
	Funcs     []*Func
	Statuses  []*Status
	//Masks     []*Mask

	Groups []Group `json:"-"`
}

func (b *Block) addSubblock(sb *Block) {
	b.Subblocks = append(b.Subblocks, sb)
}

func (b *Block) addConfig(c *Config) {
	b.Configs = append(b.Configs, c)
}

func (b *Block) addFunc(f *Func) {
	b.Funcs = append(b.Funcs, f)
}

func (b *Block) addStatus(s *Status) {
	b.Statuses = append(b.Statuses, s)
}

func (b *Block) addGroup(g Group) {
	b.Groups = append(b.Groups, g)
}

func (b *Block) hasElement(name string) bool {
	for i, _ := range b.Statuses {
		if b.Statuses[i].Name == name {
			return true
		}
	}

	return false
}
