package wrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapfloat64(t *testing.T) {
	t.Run("Wrap[float64]", func(t *testing.T) {
		cases := []struct {
			x   float64
			min float64
			max float64
			exp float64
		}{
			{45, -180, 180, 45},
			{-45, -180, 180, -45},
			{191, -180, 180, -169},
			{-191, -180, 180, 169},
			{-170, -217, 217, -170},
			{215, -217, 217, 215},
			{-215, -217, 217, -215},
			{-230, -217, 217, 204},
		}

		for _, c := range cases {
			got := Wrap[float64](c.x, c.min, c.max)
			assert.Equal(t, got, c.exp)
		}
	})
}

func TestWrapint(t *testing.T) {
	t.Run("Wrap[int]", func(t *testing.T) {
		cases := []struct {
			x   int
			min int
			max int
			exp int
		}{
			{45, -180, 180, 45},
			{-45, -180, 180, -45},
			{191, -180, 180, -169},
			{-191, -180, 180, 169},
			{-170, -217, 217, -170},
			{215, -217, 217, 215},
			{-215, -217, 217, -215},
			{-230, -217, 217, 204},
		}

		for _, c := range cases {
			got := Wrap[int](c.x, c.min, c.max)
			assert.Equal(t, got, c.exp)
		}
	})
}
