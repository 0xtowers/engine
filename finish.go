package engine

import "time"

// TODO test
func (e *Engine) Finish() int {
	// In case the game is new and there are no votes yet the game can never be
	// finished. So for that initial case we return early here.
	if e.gam == 0 {
		return 0
	}

	// In case there is no leading tower the game cannot be ended. So in that case
	// we return early here.
	if e.lea[0] == 0 {
		return 0
	}

	var sin time.Duration
	{
		sin = e.now().Sub(e.tim)
	}

	// the game is finished
	if sin >= e.dwn {
		return 1
	}

	// the game is not finished
	return 0
}
