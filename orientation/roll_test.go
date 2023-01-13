package orientation

import (
	// "fmt"
	// "math"

	"testing"

	"github.com/stretchr/testify/assert"
)

var rollattests = []struct {
	local        Orientation[float64]
	azimuth      float64
	expectedRoll float64
}{
	{
		local: Orientation[float64]{
			Pitch: 1.0,
			Roll:  0.0,
		},
		azimuth:      0.0,
		expectedRoll: 0.0,
	},
	{
		local: Orientation[float64]{
			Pitch: -1.0,
			Roll:  0.0,
		},
		azimuth:      0,
		expectedRoll: 0.0,
	},
	{
		local: Orientation[float64]{
			Pitch: 1.0,
			Roll:  1.0,
		},
		azimuth:      0.0,
		expectedRoll: 1,
	},
	{
		local: Orientation[float64]{
			Pitch: 1.0,
			Roll:  0.0,
		},
		azimuth:      90.0,
		expectedRoll: -1.0,
	},
	{
		local: Orientation[float64]{
			Pitch: 1.0,
			Roll:  0.0,
		},
		azimuth:      -90.0,
		expectedRoll: 1.0,
	},
	{
		local: Orientation[float64]{
			Pitch: 1.0,
			Roll:  1.0,
		},
		azimuth:      45.0,
		expectedRoll: 0,
	},
	{
		local: Orientation[float64]{
			Pitch: 0.0,
			Roll:  1.0,
		},
		azimuth:      90.0,
		expectedRoll: 0,
	},
	{
		local: Orientation[float64]{
			Pitch: 1.0,
			Roll:  0.0,
		},
		azimuth:      90.0,
		expectedRoll: -1,
	},
	{
		local: Orientation[float64]{
			Pitch: 1.0,
			Roll:  0.0,
		},
		azimuth:      -90.0,
		expectedRoll: 1,
	},
}

func TestRollAt(t *testing.T) {
	t.Run("RollAt", func(t *testing.T) {
		for _, tt := range rollattests {
			res := tt.local.RollAt(tt.azimuth)
			assert.InDelta(t, tt.expectedRoll, res, 0.01)
		}
	})
}
