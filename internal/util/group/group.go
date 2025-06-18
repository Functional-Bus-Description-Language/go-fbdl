package group

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func HasFunctionality(grp *fn.Group, name string) bool {
	for i := range grp.Configs {
		if grp.Configs[i].Name == name {
			return true
		}
	}
	for i := range grp.Masks {
		if grp.Masks[i].Name == name {
			return true
		}
	}
	for i := range grp.Statics {
		if grp.Statics[i].Name == name {
			return true
		}
	}
	for i := range grp.Statuses {
		if grp.Statuses[i].Name == name {
			return true
		}
	}

	return false
}

func IsEmpty(grp fn.Group) bool {
	if len(grp.Configs) > 0 ||
		len(grp.Irqs) > 0 ||
		len(grp.Masks) > 0 ||
		len(grp.Params) > 0 ||
		len(grp.Returns) > 0 ||
		len(grp.Statics) > 0 ||
		len(grp.Statuses) > 0 {
		return false
	}

	return true
}

func cantAddErrMsg(g *fn.Group, f1, f2 fn.Functionality) error {
	return fmt.Errorf(
		"cannot add %s '%s' to group '%s', group already has %s '%s'",
		f1.Type(), f1.GetName(), g.Name, f2.Type(), f2.GetName(),
	)
}

func AddConfig(grp *fn.Group, c *fn.Config) error {
	if len(grp.Irqs) > 0 {
		return cantAddErrMsg(grp, c, grp.Irqs[0])
	}
	if len(grp.Params) > 0 {
		return cantAddErrMsg(grp, c, grp.Params[0])
	}
	if len(grp.Returns) > 0 {
		return cantAddErrMsg(grp, c, grp.Returns[0])
	}

	grp.Configs = append(grp.Configs, c)

	return nil
}

func AddIrq(grp *fn.Group, i *fn.Irq) error {
	if grp.Virtual {
		return fmt.Errorf(
			"cannot add irq '%s' to virtual group '%s', virtual irq group makes no sense",
			i.Name, grp.Name,
		)
	}

	if len(grp.Configs) > 0 {
		return cantAddErrMsg(grp, i, grp.Configs[0])
	}
	if len(grp.Masks) > 0 {
		return cantAddErrMsg(grp, i, grp.Masks[0])
	}
	if len(grp.Params) > 0 {
		return cantAddErrMsg(grp, i, grp.Params[0])
	}
	if len(grp.Returns) > 0 {
		return cantAddErrMsg(grp, i, grp.Returns[0])
	}
	if len(grp.Statics) > 0 {
		return cantAddErrMsg(grp, i, grp.Statics[0])
	}
	if len(grp.Statuses) > 0 {
		return cantAddErrMsg(grp, i, grp.Statuses[0])
	}

	grp.Irqs = append(grp.Irqs, i)

	return nil
}

func AddMask(grp *fn.Group, m *fn.Mask) error {
	if len(grp.Irqs) > 0 {
		return cantAddErrMsg(grp, m, grp.Irqs[0])
	}
	if len(grp.Params) > 0 {
		return cantAddErrMsg(grp, m, grp.Params[0])
	}
	if len(grp.Returns) > 0 {
		return cantAddErrMsg(grp, m, grp.Returns[0])
	}

	grp.Masks = append(grp.Masks, m)

	return nil
}

func AddParam(grp *fn.Group, p *fn.Param) error {
	if grp.Virtual {
		return fmt.Errorf(
			"cannot add param '%s' to virtual group '%s', virtual param group makes no sense",
			p.Name, grp.Name,
		)
	}

	if len(grp.Configs) > 0 {
		return cantAddErrMsg(grp, p, grp.Configs[0])
	}
	if len(grp.Irqs) > 0 {
		return cantAddErrMsg(grp, p, grp.Irqs[0])
	}
	if len(grp.Masks) > 0 {
		return cantAddErrMsg(grp, p, grp.Masks[0])
	}
	if len(grp.Returns) > 0 {
		return cantAddErrMsg(grp, p, grp.Returns[0])
	}
	if len(grp.Statics) > 0 {
		return cantAddErrMsg(grp, p, grp.Statics[0])
	}
	if len(grp.Statuses) > 0 {
		return cantAddErrMsg(grp, p, grp.Statuses[0])
	}

	grp.Params = append(grp.Params, p)

	return nil
}

func AddReturn(grp *fn.Group, r *fn.Return) error {
	if grp.Virtual {
		return fmt.Errorf(
			"cannot add return '%s' to virtual group '%s', virtual return group makes no sense",
			r.Name, grp.Name,
		)
	}

	if len(grp.Configs) > 0 {
		return cantAddErrMsg(grp, r, grp.Configs[0])
	}
	if len(grp.Irqs) > 0 {
		return cantAddErrMsg(grp, r, grp.Irqs[0])
	}
	if len(grp.Masks) > 0 {
		return cantAddErrMsg(grp, r, grp.Masks[0])
	}
	if len(grp.Params) > 0 {
		return cantAddErrMsg(grp, r, grp.Params[0])
	}
	if len(grp.Statics) > 0 {
		return cantAddErrMsg(grp, r, grp.Statics[0])
	}
	if len(grp.Statuses) > 0 {
		return cantAddErrMsg(grp, r, grp.Statuses[0])
	}

	grp.Returns = append(grp.Returns, r)

	return nil
}

func AddStatic(grp *fn.Group, s *fn.Static) error {
	if len(grp.Irqs) > 0 {
		return cantAddErrMsg(grp, s, grp.Irqs[0])
	}
	if len(grp.Params) > 0 {
		return cantAddErrMsg(grp, s, grp.Params[0])
	}
	if len(grp.Returns) > 0 {
		return cantAddErrMsg(grp, s, grp.Returns[0])
	}

	grp.Statics = append(grp.Statics, s)

	return nil
}

func AddStatus(grp *fn.Group, s *fn.Status) error {
	if len(grp.Irqs) > 0 {
		return cantAddErrMsg(grp, s, grp.Irqs[0])
	}
	if len(grp.Params) > 0 {
		return cantAddErrMsg(grp, s, grp.Params[0])
	}
	if len(grp.Returns) > 0 {
		return cantAddErrMsg(grp, s, grp.Returns[0])
	}

	grp.Statuses = append(grp.Statuses, s)

	return nil
}
