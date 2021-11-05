package reg

import (
	"sort"
	"strings"
)

func assignGlobalAccessAddresses(bus *Block, baseAddr uint) {
	// Currently there is only Block Align strategy.
	// In the future there may also be Compact and Full Align.

	assignGlobalAccessAddressesBlockAlign(bus, baseAddr)
}

func assignGlobalAccessAddressesBlockAlign(block *Block, baseAddr uint) {
	if block.IsArray() {
		block.AddrSpace = AddrSpaceArray{
			start:     baseAddr,
			count:     uint(block.Count),
			BlockSize: block.Sizes.BlockAligned,
		}
	} else {
		block.AddrSpace = AddrSpaceSingle{
			start: baseAddr,
			end:   baseAddr + block.Sizes.BlockAligned - 1,
		}
	}

	if len(block.Blocks) == 0 {
		return
	}

	sortFunc := func(i, j int) bool {
		sizei := block.Blocks[i].Sizes.BlockAligned
		sizej := block.Blocks[j].Sizes.BlockAligned

		namei := block.Blocks[i].Name
		namej := block.Blocks[j].Name

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
	sort.Slice(block.Blocks, sortFunc)

	subblockBaseAddr := block.AddrSpace.End() + 1
	for i, _ := range block.Blocks {
		sb := &block.Blocks[i]
		subblockBaseAddr -= uint(sb.Count) * sb.Sizes.BlockAligned
		assignGlobalAccessAddressesBlockAlign(sb, subblockBaseAddr)
	}
}
