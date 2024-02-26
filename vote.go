package engine

import (
	"fmt"
	"sort"
	"time"

	"github.com/xh3b4sd/tracer"
)

func (e *Engine) Vote(tow int, use string, amo int, bal int) error {
	var err error

	{
		err = e.verifyVote(tow, use, amo, bal)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// record time of latest deposit
	{
		e.tim = time.Now().UTC()
	}

	// record user vote for tower
	{
		e.gam += amo
		e.dep[tow] += 1
		e.tow[tow] += amo
		e.use[tow][use] += amo
	}

	return nil
}

func (e *Engine) verifyVote(tow int, use string, amo int, bal int) error {
	if e.Finish() == 1 {
		return fmt.Errorf("game is finished")
	}

	if tow != 1 && tow != 2 && tow != 3 {
		return fmt.Errorf("tower must be 1, 2 or 3")
	}

	if use == "" {
		return fmt.Errorf("use must not be empty")
	}

	if amo <= 0 {
		return fmt.Errorf("amo must not be empty")
	}

	if amo > 100 {
		return fmt.Errorf("amo must not be greater 100")
	}

	if (bal - (e.use[tow][use] + amo)) < 0 {
		return fmt.Errorf("insufficient balance")
	}

	// leading tower must not be higher than 3 deposits
	var lea []int
	{
		lea = e.leadingTower()
	}

	if lea[0] == tow && (e.dep[lea[0]]-e.dep[lea[1]]) >= 3 {
		return fmt.Errorf("leading tower has maximum lead")
	}

	return nil
}

func (e *Engine) leadingTower() []int {
	type wrp struct {
		dep int
		tow int
	}

	var srt []wrp
	{
		srt = []wrp{
			{dep: e.dep[1], tow: 1},
			{dep: e.dep[2], tow: 2},
			{dep: e.dep[3], tow: 3},
		}
	}

	{
		sort.SliceStable(srt, func(i, j int) bool { return srt[i].dep > srt[j].dep })
	}

	var tow []int
	for _, x := range srt {
		tow = append(tow, x.tow)
	}

	return tow
}
