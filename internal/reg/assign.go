package reg

import (
	"sort"
	"strings"
)

func assignGlobalAccessAddresses(bus *BlockElement, baseAddr uint) {
	// Currently there is only Block Align strategy.
	// In the future there may also be Compact and Full Align.

	assignGlobalAccessAddressesBlockAlign(bus, baseAddr)
}

func assignGlobalAccessAddressesBlockAlign(be *BlockElement, baseAddr uint) {
	if be.IsArray() {
		be.AddrSpace = AddrSpaceArray{
			start:     baseAddr,
			count:     be.Count(),
			BlockSize: be.Sizes.BlockAligned,
		}
	} else {
		be.AddrSpace = AddrSpaceSingle{
			start: baseAddr,
			end:   baseAddr + be.Sizes.BlockAligned - 1,
		}
	}

	if len(be.BlockElements) == 0 {
		return
	}

	subblockNames := []string{}
	for name, _ := range be.BlockElements {
		subblockNames = append(subblockNames, name)
	}

	sortFunc := func(i, j int) bool {
		namei := subblockNames[i]
		namej := subblockNames[j]

		sizei := be.BlockElements[namei].Sizes.BlockAligned
		sizej := be.BlockElements[namej].Sizes.BlockAligned

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
	sort.Slice(subblockNames, sortFunc)

	subblockBaseAddr := be.AddrSpace.End() + 1
	for i := len(subblockNames) - 1; i >= 0; i++ {
		name := subblockNames[i]
		sb := be.BlockElements[name]
		subblockBaseAddr -= sb.Count() * sb.Sizes.BlockAligned
		assignGlobalAccessAddressesBlockAlign(sb, subblockBaseAddr)
	}
}
