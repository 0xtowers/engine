package engine

import "time"

type Config struct {
	// Dwn is the countdown duration that has to pass without any further deposits
	// in order for the leading tower to win.
	Dwn time.Duration
}

type Engine struct {
	bal map[string]int
	dep map[int]int
	dwn time.Duration
	gam int
	tim time.Time
	tnx map[string]int
	tow map[int]int
	use map[int]map[string]int
}

func New(c Config) *Engine {
	if c.Dwn == 0 {
		c.Dwn = time.Hour
	}

	return &Engine{
		bal: map[string]int{},
		dep: map[int]int{},
		dwn: c.Dwn,
		tnx: map[string]int{},
		tow: map[int]int{
			1: 0,
			2: 0,
			3: 0,
		},
		use: map[int]map[string]int{
			1: {},
			2: {},
			3: {},
		},
	}
}
