package orientation

import (
	// "fmt"
	// "math"

	"testing"

	"github.com/stretchr/testify/assert"
)

var azimuthattests = []struct {
	local           Orientation[float64]
	worldAzimuth    float64
	expectedAzimuth float64
}{
	{
		local: Orientation[float64]{
			Heading: 90.0,
		},
		worldAzimuth:    90.0,
		expectedAzimuth: 0.0,
	},
	{
		local: Orientation[float64]{
			Heading: 90.0,
		},
		worldAzimuth:    91.0,
		expectedAzimuth: 1.0,
	},
	{
		local: Orientation[float64]{
			Heading: 90.0,
		},
		worldAzimuth:    0.0,
		expectedAzimuth: -90.0,
	},
	{
		local: Orientation[float64]{
			Heading: -90.0,
		},
		worldAzimuth:    0.0,
		expectedAzimuth: 90.0,
	},
	{
		local: Orientation[float64]{
			Heading: 120.0,
		},
		worldAzimuth:    90.0,
		expectedAzimuth: -30.0,
	},
}

func TestAzimuthAt(t *testing.T) {
	t.Run("AzimuthAt", func(t *testing.T) {
		for _, tt := range azimuthattests {
			res := tt.local.AzimuthAt(tt.worldAzimuth)
			assert.InDelta(t, tt.expectedAzimuth, res, 0.01)
		}
	})
}

var worldazimuthattests = []struct {
	local                Orientation[float64]
	azimuth              float64
	expectedWorldAzimuth float64
}{
	{
		local: Orientation[float64]{
			Heading: 90.0,
		},
		azimuth:              90.0,
		expectedWorldAzimuth: 180.0,
	},
	{
		local: Orientation[float64]{
			Heading: 90.0,
		},
		azimuth:              0.0,
		expectedWorldAzimuth: 90.0,
	},
	{
		local: Orientation[float64]{
			Heading: -90.0,
		},
		azimuth:              0.0,
		expectedWorldAzimuth: -90.0,
	},
	{
		local: Orientation[float64]{
			Heading: 120.0,
		},
		azimuth:              90.0,
		expectedWorldAzimuth: -150.0,
	},
}

func TestWorldAzimuthAt(t *testing.T) {
	t.Run("WorldAzimuthAt", func(t *testing.T) {
		for _, tt := range worldazimuthattests {
			res := tt.local.WorldAzimuthAt(tt.azimuth)
			assert.InDelta(t, tt.expectedWorldAzimuth, res, 0.01)
		}
	})
}
