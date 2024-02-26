package engine

func (e *Engine) Balance(use string) int {
	var bal int
	{
		bal = e.bal[use]
	}

	// The minimum deposit into the game is 1000 tokens by default. We do not know
	// in advance which users are going to play the game. That means we cannot
	// fetch user balances prematurely. Our assumption is that any user making a
	// vote has deposited into the game at least once, providing every user with a
	// minimum balance of 1000.
	if bal == 0 {
		return 1000
	}

	return bal
}
