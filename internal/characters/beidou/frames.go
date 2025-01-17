package beidou

import "github.com/genshinsim/gcsim/pkg/core"

func (c *char) ActionFrames(a core.ActionType, p map[string]int) (int, int) {
	switch a {
	case core.ActionAttack:
		f := 0
		switch c.NormalCounter {
		//TODO: need to add atkspd mod
		case 0:
			f = 23 //frames from keqing lib
		case 1:
			f = 43
		case 2:
			f = 68
		case 3:
			f = 44
		case 4:
			f = 68
		}
		atkspd := c.Stat(core.AtkSpd)
		if c.Core.Status.Duration("beidoua4") > 0 {
			atkspd += 0.15
		}
		f = int(float64(f) / (1 + atkspd))
		return f, f
	case core.ActionCharge:
		f := 35 //frames from keqing lib
		atkspd := c.Stat(core.AtkSpd)
		if c.Core.Status.Duration("beidoua4") > 0 {
			atkspd += 0.15
		}
		f = int(float64(f) / (1 + atkspd))
		return f, f
	case core.ActionSkill:
		return 41, 41 //ok
	case core.ActionBurst:
		return 45, 45 //ok
	default:
		c.Core.Log.NewEventBuildMsg(core.LogActionEvent, c.Index, "unknown action (invalid frames): ", a.String())
		return 0, 0
	}
}
