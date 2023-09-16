package gap

type Array struct {
	StartAddr int64
	EndAddr   int64
	StartBit  int64
	EndBit    int64
	WriteSafe bool
}

func (a Array) isGap()       {}
func (a Array) Width() int64 { return a.EndBit - a.StartBit + 1 }
