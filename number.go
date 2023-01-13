package math

// import "golang.org/x/exp/constraints"

type Number interface {
	// constraints.Integer | constraints.Float
	int16 | int32 | int64 | int | float32 | float64
}
