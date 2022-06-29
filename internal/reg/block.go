package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func blkAddConfig(b *elem.Block, c *elem.Config) { b.Configs = append(b.Configs, c) }
func blkAddFunc(b *elem.Block, f *elem.Func)     { b.Funcs = append(b.Funcs, f) }
func blkAddMask(b *elem.Block, m *elem.Mask)     { b.Masks = append(b.Masks, m) }
func blkAddGroup(b *elem.Block, g elem.Group)    { b.Groups = append(b.Groups, g) }
func blkAddStatus(b *elem.Block, s *elem.Status) { b.Statuses = append(b.Statuses, s) }
func blkAddStream(b *elem.Block, s *elem.Stream) { b.Streams = append(b.Streams, s) }
func blkAddSubblock(b, sb *elem.Block)           { b.Subblocks = append(b.Subblocks, sb) }
