package engine

import "time"

func (e *Engine) Finish() int {
	var sin time.Duration
	{
		sin = time.Now().UTC().Sub(e.tim)
	}

	// the game is finished
	if sin >= e.dwn {
		return 1
	}

	// the game is not finished
	return 0
}
