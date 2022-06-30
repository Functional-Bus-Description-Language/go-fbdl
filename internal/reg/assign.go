package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
	"sort"
	"strings"
)

func assignGlobalAccessAddresses(bus *elem.Block, baseAddr int64) {
	// Currently there is only Block Align strategy.
	// In the future there may also be Compact and Full Align.

	assignGlobalAccessAddressesBlockAlign(bus, baseAddr)
}

func assignGlobalAccessAddressesBlockAlign(block *elem.Block, baseAddr int64) {
	if block.IsArray() {
		block.AddrSpace = access.MakeAddrSpaceArray(
			baseAddr, int64(block.Count()), block.Sizes.BlockAligned,
		)
	} else {
		block.AddrSpace = access.MakeAddrSpaceSingle(
			baseAddr, baseAddr+block.Sizes.BlockAligned-1,
		)
	}

	if len(block.Subblocks) == 0 {
		return
	}

	sortFunc := func(i, j int) bool {
		sizei := block.Subblocks[i].Sizes.BlockAligned
		sizej := block.Subblocks[j].Sizes.BlockAligned

		namei := block.Subblocks[i].Name()
		namej := block.Subblocks[j].Name()

		if sizei < sizej {
			return true
		} else if sizei > sizej {
			return false
		} else {
			if strings.Compare(namei, namej) < 0 {
				return true
			} else {
				return false
			}
		}
	}
	sort.Slice(block.Subblocks, sortFunc)

	subblockBaseAddr := block.AddrSpace.End() + 1
	// Iterate subblocks in decreasing size order.
	for i := len(block.Subblocks) - 1; i >= 0; i-- {
		sb := block.Subblocks[i]
		subblockBaseAddr -= sb.Count() * sb.Sizes.BlockAligned
		assignGlobalAccessAddressesBlockAlign(sb, subblockBaseAddr)
	}
}
