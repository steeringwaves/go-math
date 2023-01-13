package wrap

import (
	"math"

	goMath "github.com/steeringwaves/go-math"
)

/* wrap x -> [0,max) */
func wrapMax[T goMath.Number](x, max T) T {
	/* integer math: `(max + x % max) % max` */
	return T(math.Mod(float64(max)+math.Mod(float64(x), float64(max)), float64(max)))
}

/* wrap x -> [min,max) */
func Wrap[T goMath.Number](x, min, max T) T {
	return min + wrapMax[T](x-min, max-min)
}
