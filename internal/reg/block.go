package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func blockAddStatus(b *elem.Block, s *elem.Status) { b.Statuses = append(b.Statuses, s) }

func blockAddGroup(b *elem.Block, g elem.Group) { b.Groups = append(b.Groups, g) }
