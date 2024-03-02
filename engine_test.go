package engine

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Engine_leadingTower(t *testing.T) {
	testCases := []struct {
		vot [][2]int
		tow [][3]int
	}{
		// Case 000
		{
			vot: [][2]int{
				{1, 10},
				{2, 10},
				{3, 10},
			},
			tow: [][3]int{
				{1, 2, 3},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		// Case 001
		{
			vot: [][2]int{
				{1, 10},
				{2, 10},
				{1, 10},
				{3, 10},
			},
			tow: [][3]int{
				{1, 2, 3},
				{0, 0, 0},
				{1, 2, 3},
				{1, 2, 3},
			},
		},
		// Case 002
		{
			vot: [][2]int{
				{1, 10},
				{2, 10},
				{1, 10},
				{3, 10},
				{2, 10},
			},
			tow: [][3]int{
				{1, 2, 3},
				{0, 0, 0},
				{1, 2, 3},
				{1, 2, 3},
				{0, 0, 0},
			},
		},
		// Case 003
		{
			vot: [][2]int{
				{1, 10},
				{2, 10},
				{1, 10},
				{3, 10},
				{2, 10},
				{3, 10},
				{3, 10},
			},
			tow: [][3]int{
				{1, 2, 3},
				{0, 0, 0},
				{1, 2, 3},
				{1, 2, 3},
				{0, 0, 0},
				{0, 0, 0},
				{3, 1, 2},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var eng *Engine
			{
				eng = New(Config{})
			}

			var tow [][3]int
			for i, x := range tc.vot {
				{
					err := eng.Vote(x[0], fmt.Sprintf("use-%d", i), x[1], 1000)
					if err != nil {
						t.Fatal(err)
					}
				}

				{
					tow = append(tow, eng.leadingTower())
				}
			}

			if !reflect.DeepEqual(tow, tc.tow) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.tow, tow))
			}
		})
	}
}
