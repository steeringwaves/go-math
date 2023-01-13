package orientation

import (
	// "fmt"
	// "math"

	"testing"

	"github.com/stretchr/testify/assert"
)

var translatetests = []struct {
	local    Orientation[float64]
	azimuth  float64
	expected Orientation[float64]
}{
	{
		local: Orientation[float64]{
			Pitch:   1.0,
			Roll:    0.0,
			Heading: 0,
		},
		azimuth: 90.0,
		expected: Orientation[float64]{
			Pitch:   0.0,
			Roll:    1.0,
			Heading: -90,
		},
	},
	{
		local: Orientation[float64]{
			Pitch:   0.2,
			Roll:    0.1,
			Heading: 0,
		},
		azimuth: 25.0,
		expected: Orientation[float64]{
			Pitch:   0.14,
			Roll:    0.18,
			Heading: -25,
		},
	},
	{
		local: Orientation[float64]{
			Pitch:   1.1,
			Roll:    -0.6,
			Heading: 0,
		},
		azimuth: 130.3,
		expected: Orientation[float64]{
			Pitch:   -0.25,
			Roll:    1.23,
			Heading: -130.3,
		},
	},
	{
		local: Orientation[float64]{
			Pitch:   -0.6,
			Roll:    -1.5,
			Heading: 0,
		},
		azimuth: 222.1,
		expected: Orientation[float64]{
			Pitch:   -0.56,
			Roll:    1.52,
			Heading: 137.9,
		},
	},
	{
		local: Orientation[float64]{
			Pitch:   -2.2,
			Roll:    3.1,
			Heading: 0,
		},
		azimuth: 348,
		expected: Orientation[float64]{
			Pitch:   -1.51,
			Roll:    3.49,
			Heading: 12,
		},
	},
}

func TestTranslateFrom(t *testing.T) {
	t.Run("TranslateFrom", func(t *testing.T) {
		for _, tt := range translatetests {
			res := tt.local.TranslateFrom(tt.azimuth)
			assert.InDelta(t, tt.expected.Pitch, res.Pitch, 0.01)
			assert.InDelta(t, tt.expected.Roll, res.Roll, 0.01)
			assert.InDelta(t, tt.expected.Heading, res.Heading, 0.01)
		}
	})
}
