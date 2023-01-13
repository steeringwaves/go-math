package clamp

import (
	goMath "github.com/steeringwaves/go-math"
)

func Clamp[T goMath.Number](x, min, max T) T {
	switch {
	case x < min:
		return min
	case x > max:
		return max
	default:
		return x
	}
}
