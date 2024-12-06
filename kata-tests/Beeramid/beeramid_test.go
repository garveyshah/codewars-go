package beeramid

import (
	"testing"
)

func TestBeeramid(t *testing.T) {
	t.Parallel()
	tt := []struct {
		refBonus int
		price float64
		want int

	}{
		{9, 2.0, 1},
		{21, 1.5, 3},
		{-1, 4.0, 0},
	}

	for i, tc := range tt {
		got := Beeramid(tc.refBonus, tc.price)
		if tc.want != got {
			t.Fatalf("Test[%d] failed - got= %d want= %d", i+1, got, tc.want)
		}
	}
		
}