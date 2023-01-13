package orientation

import (
	"math"
)

// RollAt Calculates the roll angle at the target azimuth in degrees
func (o *Orientation[T]) RollAt(azimuth T) T {
	pitchRads := float64(o.Pitch) * M_PI_180
	rollRads := float64(o.Roll) * M_PI_180
	azimuthRads := ScaleDegrees[float64]((float64(azimuth))+90.0, false) * M_PI_180

	angleRads := (pitchRads * math.Cos(azimuthRads)) - ((rollRads * -1.0) * math.Sin(azimuthRads))

	return T(angleRads * M_180_PI)
}
