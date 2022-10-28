package access

type Mask interface {
	Start() int64
	End() int64
	Width() int64
}
