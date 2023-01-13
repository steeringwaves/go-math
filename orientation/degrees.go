package orientation

import (
	"math"

	"github.com/steeringwaves/go-math/clamp"
	"github.com/steeringwaves/go-math/wrap"
)

// ScaleDegrees will scale degrees to either -180,180 or 0-360 scales
func ScaleDegrees[T int16 | int32 | int64 | int | float32 | float64](degrees T, use360Scale bool) T {
	var scaledDeg float64

	if float64(degrees) < 0 {
		scaledDeg = math.Mod(float64(degrees), -360)
	} else {
		scaledDeg = math.Mod(float64(degrees), 360)
	}

	if !use360Scale {
		if (scaledDeg < -180) || (scaledDeg > 180) {
			if scaledDeg > 180 {
				scaledDeg -= 360
			} else if scaledDeg < -180 {
				scaledDeg += 360
			}
		}

		if scaledDeg == -180 {
			scaledDeg += 0.01
		} else if scaledDeg == 180 {
			scaledDeg -= 0.01
		}
	} else {
		if (scaledDeg < 0) || (scaledDeg > 360) {
			if scaledDeg < 0 {
				scaledDeg += 360
			} else if scaledDeg > 360 {
				scaledDeg -= 360
			}
		}

		if scaledDeg == 360 {
			scaledDeg -= 0.01
		}
	}

	return T(scaledDeg)
}

func DegreeWrapDelta[T int16 | int32 | int64 | int | float32 | float64](a, b T) T {
	if b >= 0 {
		return wrap.Wrap[T](wrap.Wrap[T](b, T(-180), T(180))-wrap.Wrap[T](a, T(-180), T(180))-T(360), T(-180), T(180))
	}

	return wrap.Wrap[T](wrap.Wrap[T](b, T(-180), T(180))-wrap.Wrap[T](a, T(-180), T(180))+T(360), T(-180), T(180))
}

func DegreeDelta[T int16 | int32 | int64 | int | float32 | float64](a, b, min, max T) T {
	return b - a
}

func DegreeClampDelta[T int16 | int32 | int64 | int | float32 | float64](a, b, min, max T) T {
	return clamp.Clamp[T](b, min, max) - clamp.Clamp[T](a, min, max)
}

func DegreeWrapCompare[T int16 | int32 | int64 | int | float32 | float64](a, b T, tolerance float64) bool {
	return (math.Abs(float64(DegreeWrapDelta[T](a, b))) <= tolerance)
}

func DegreeCompare[T int16 | int32 | int64 | int | float32 | float64](a, b, min, max T, tolerance float64) bool {
	return (math.Abs(float64(DegreeDelta[T](a, b, min, max))) <= tolerance)
}

func DegreeClampCompare[T int16 | int32 | int64 | int | float32 | float64](a, b, min, max T, tolerance float64) bool {
	return (math.Abs(float64(DegreeClampDelta[T](a, b, min, max))) <= tolerance)
}
