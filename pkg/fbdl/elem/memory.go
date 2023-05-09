package elem

type Memory struct {
	Elem

	Access          string
	ByteWriteEnable bool
	ReadLatency     int64
	Size            int64
	Width           int64
}
