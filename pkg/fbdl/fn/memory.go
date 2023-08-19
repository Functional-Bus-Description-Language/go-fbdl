package fn

type Memory struct {
	Func

	Access          string
	ByteWriteEnable bool
	ReadLatency     int64
	Size            int64
	Width           int64
}
