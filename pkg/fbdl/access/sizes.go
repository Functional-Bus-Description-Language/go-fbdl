package access

type Sizes interface {
	BlockAligned() int64
	Compact() int64
	Own() int64
}
