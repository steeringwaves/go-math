package clamp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClampfloat64(t *testing.T) {
	t.Run("Clamp[float64]", func(t *testing.T) {
		cases := []struct {
			x   float64
			min float64
			max float64
			exp float64
		}{
			{45, -180, 180, 45},
			{-45, -180, 180, -45},
			{191, -180, 180, 180},
			{-191, -180, 180, -180},
			{-170, -217, 217, -170},
			{215, -217, 217, 215},
			{-215, -217, 217, -215},
			{-230, -217, 217, -217},
		}

		for _, c := range cases {
			got := Clamp[float64](c.x, c.min, c.max)
			assert.Equal(t, got, c.exp)
		}
	})
}

func TestClampint(t *testing.T) {
	t.Run("Clamp[int]", func(t *testing.T) {
		cases := []struct {
			x   int
			min int
			max int
			exp int
		}{
			{45, -180, 180, 45},
			{-45, -180, 180, -45},
			{191, -180, 180, 180},
			{-191, -180, 180, -180},
			{-170, -217, 217, -170},
			{215, -217, 217, 215},
			{-215, -217, 217, -215},
			{-230, -217, 217, -217},
		}

		for _, c := range cases {
			got := Clamp[int](c.x, c.min, c.max)
			assert.Equal(t, got, c.exp)
		}
	})
}
