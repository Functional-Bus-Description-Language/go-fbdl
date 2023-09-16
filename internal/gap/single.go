package gap

type Single struct {
	Addr      int64
	StartBit  int64
	EndBit    int64
	WriteSafe bool
}

func (s Single) isGap()       {}
func (s Single) Width() int64 { return s.EndBit - s.StartBit + 1 }
