package reg

// Address represents block internal address.
type address struct {
	value     int64 // Current address value
	byteWidth int64 // Register byte width, busWidth / 8
}

func makeAddr(value int64, busWidth int64) address {
	return address{value: value, byteWidth: busWidth / 8}
}

func (addr *address) inc(value int64) {
	// Use below when byte addressing will be the only addressing.
	//addr.value += value * addr.byteWidth
	addr.value += 1 * value
}
