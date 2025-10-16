package types

var busWidth int64

// bw is the busWidth.
//
// Having bus width as a global package variable helps keeping access make
// functions easier to call, as all of them require information on bus width.
func Init(bw int64) {
	busWidth = bw
}
