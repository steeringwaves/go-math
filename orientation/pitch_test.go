package orientation

import (
	// "fmt"
	// "math"

	"testing"

	"github.com/stretchr/testify/assert"
)

var pitchattests = []struct {
	local         Orientation[float64]
	azimuth       float64
	expectedPitch float64
}{
	{
		local: Orientation[float64]{
			Pitch: 1.0,
			Roll:  0.0,
		},
		azimuth:       0.0,
		expectedPitch: 1.0,
	},
	{
		local: Orientation[float64]{
			Pitch: -1.0,
			Roll:  0.0,
		},
		azimuth:       0,
		expectedPitch: -1.0,
	},
	{
		local: Orientation[float64]{
			Pitch: 0.0,
			Roll:  1.0,
		},
		azimuth:       90.0,
		expectedPitch: 1,
	},
	{
		local: Orientation[float64]{
			Pitch: -1.0,
			Roll:  1.0,
		},
		azimuth:       0.0,
		expectedPitch: -1.0,
	},
	{
		local: Orientation[float64]{
			Pitch: 1.0,
			Roll:  1.0,
		},
		azimuth:       45.0,
		expectedPitch: 1.414,
	},
	{
		local: Orientation[float64]{
			Pitch: 0.2,
			Roll:  0.1,
		},
		azimuth:       25.0,
		expectedPitch: 0.2235,
	},
	{
		local: Orientation[float64]{
			Pitch: 1.1,
			Roll:  -0.6,
		},
		azimuth:       130.3,
		expectedPitch: -1.169,
	},
	{
		local: Orientation[float64]{
			Pitch: -0.6,
			Roll:  -1.5,
		},
		azimuth:       222.1,
		expectedPitch: 1.4508,
	},
	{
		local: Orientation[float64]{
			Pitch: -2.2,
			Roll:  3.1,
		},
		azimuth:       348,
		expectedPitch: -2.796,
	},
}

func TestPitchAt(t *testing.T) {
	t.Run("PitchAt", func(t *testing.T) {
		for _, tt := range pitchattests {
			res := tt.local.PitchAt(tt.azimuth)
			assert.InDelta(t, tt.expectedPitch, res, 0.01)
		}
	})
}
