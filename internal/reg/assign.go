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

func assignGlobalAccessAddressesBlockAlign(blk *elem.Block, baseAddr int64) {
	if blk.IsArray {
		blk.AddrSpace = access.MakeAddrSpaceArray(
			baseAddr, int64(blk.Count), blk.Sizes.BlockAligned,
		)
	} else {
		blk.AddrSpace = access.MakeAddrSpaceSingle(
			baseAddr, baseAddr+blk.Sizes.BlockAligned-1,
		)
	}

	if len(blk.Subblocks) == 0 {
		return
	}

	sortFunc := func(i, j int) bool {
		sizei := blk.Subblocks[i].Sizes.BlockAligned
		sizej := blk.Subblocks[j].Sizes.BlockAligned

		namei := blk.Subblocks[i].Name
		namej := blk.Subblocks[j].Name

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
	sort.Slice(blk.Subblocks, sortFunc)

	subblockBaseAddr := blk.AddrSpace.End() + 1
	// Iterate subblocks in decreasing size order.
	for i := len(blk.Subblocks) - 1; i >= 0; i-- {
		sb := blk.Subblocks[i]
		subblockBaseAddr -= sb.Count * sb.Sizes.BlockAligned
		assignGlobalAccessAddressesBlockAlign(sb, subblockBaseAddr)
	}
}
