package reg

import (
	"sort"
	"strings"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"
)

func assignGlobalAccessAddresses(blk *fn.Block, baseAddr int64) {
	if blk.IsArray {
		blk.AddrSpace = types.SingleRange{
			Start: baseAddr,
			End:   baseAddr + blk.Count*blk.Sizes.Aligned - 1,
		}
	} else {
		blk.AddrSpace = types.SingleRange{
			Start: baseAddr,
			End:   baseAddr + blk.Sizes.Aligned - 1,
		}
	}

	if len(blk.Subblocks) == 0 {
		return
	}

	sortFunc := func(i, j int) bool {
		sizei := blk.Subblocks[i].Sizes.Aligned
		sizej := blk.Subblocks[j].Sizes.Aligned

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

	subblockBaseAddr := blk.AddrSpace.End + 1
	// Iterate subblocks in decreasing size order.
	for i := len(blk.Subblocks) - 1; i >= 0; i-- {
		sb := blk.Subblocks[i]
		subblockBaseAddr -= sb.Count * sb.Sizes.Aligned
		assignGlobalAccessAddresses(sb, subblockBaseAddr)
	}
}
