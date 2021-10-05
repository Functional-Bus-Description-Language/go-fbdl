package fbdl

// Do not start from 0.
// Starting from greater value makes it easier to search for ids in packages dump.
var current_id uint32 = 0xFFF

func generateId() uint32 {
	current_id += 1

	return current_id
}
