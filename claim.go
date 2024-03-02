package engine

import (
	"fmt"
	"math/big"
)

func (e *Engine) Claim(use string) (*big.Int, error) {
	if e.Finish() == 0 {
		return nil, fmt.Errorf("game is not finished")
	}

	var tow int
	{
		tow = e.leadingTower()[0]
	}

	return e.rewards(tow, use), nil
}

func (e *Engine) rewards(tow int, use string) *big.Int {
	var gmv *big.Int
	var twv *big.Int
	var plv *big.Int
	{
		gmv = big.NewInt(int64(e.gam))
		twv = big.NewInt(int64(e.tow[tow]))
		plv = big.NewInt(int64(e.use[tow][use]))
	}

	// Calculate the claimable share of the calling player, which is their
	// fraction of value deposited into the winning tower relative to that tower's
	// total value. If the tower value is 100 and the player deposited 10, then
	// the player's share is 10%.
	//
	//         player         tower
	//     ------------- = ------------
	//           x             100
	//
	var shr *big.Int
	{
		shr = big.NewInt(0).Div(
			big.NewInt(0).Mul(
				plv,
				big.NewInt(100),
			),
			twv,
		)
	}

	// Calculate the claimable rewards for the calling player according to their
	// claimable share. If the player's share is 10% and the game value is 300,
	// then the claimable rewards are 30.
	//
	//           x           game
	//     ------------ = -----------
	//         share         100
	//
	var rwr *big.Int
	{
		rwr = big.NewInt(0).Div(
			big.NewInt(0).Mul(
				shr,
				gmv,
			),
			big.NewInt(100),
		)
	}

	return rwr
}
