package print

import (
	"fmt"

	con "golang.org/x/exp/constraints"
)

func Identity[T any](v T) T {
	return v
}

type CT interface {
	~int | ~float32 | ~float64
}

type Bunch[T CT] []T

func Bunching[E CT, T Bunch[E]](data T) T {

	for i, val := range data {
		if i+1 < len(data) && val > data[i+1] {
			fmt.Println(val)
		}
	}

	return data
}

type idFunc[T any] func(T) T

type Pointish interface {
	~struct{ X, Y int }
}

func GetX[T con.Integer](d T) T {
	return d
}
func Equal[T comparable](x, y T) bool {
	return x == y
}
