package engine

import (
	"fmt"

	"github.com/xh3b4sd/tracer"
)

func (e *Engine) Deposit(tnx string, use string, bal int) error {
	var err error

	{
		err = e.verifyDeposit(tnx, use, bal)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		e.bal[use] += bal
		e.tnx[tnx] = bal
	}

	return nil
}

func (e *Engine) verifyDeposit(tnx string, use string, bal int) error {
	if e.Finish() == 1 {
		return fmt.Errorf("game is finished")
	}

	{
		_, exi := e.tnx[tnx]
		if exi {
			return fmt.Errorf("tnx must not exist")
		}
	}

	if use == "" {
		return fmt.Errorf("use must not be empty")
	}

	if bal <= 0 {
		return fmt.Errorf("bal must not be empty")
	}

	return nil
}
