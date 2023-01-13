package orientation

import (
	"math"
)

// PitchAt Calculates the pitch angle at the target azimuth in degrees
func (o *Orientation[T]) PitchAt(azimuth T) T {
	pitchRads := float64(o.Pitch) * M_PI_180
	rollRads := float64(o.Roll) * M_PI_180
	azimuthRads := float64(azimuth) * M_PI_180

	angleRads := (pitchRads * math.Cos(azimuthRads)) - ((rollRads * -1.0) * math.Sin(azimuthRads))

	return T(angleRads * M_180_PI)
}
