package orientation

import (
	"math"
)

const (
	M_PI_180 = math.Pi / 180.0 // pi/180
	M_180_PI = 180.0 / math.Pi // 180/pi
)

type Orientation[T int16 | int32 | int64 | int | float32 | float64] struct {
	Pitch   T
	Roll    T
	Heading T // Only valid if it is magnetic heading
}
