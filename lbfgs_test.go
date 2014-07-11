package lbfgs

import (
	"fmt"
	"testing"
)

func TestOptimizer(t *testing.T) {
	optimizer := NewOptimizer(10, 2)

	x := NewVector(2)
	g := NewVector(2)
	x.SetValues([]float64{1, 1})


	for k := 0; k < 20; k++ {
		fmt.Println("======> k = ", k)
		fmt.Println("x = ", x)
		g.SetValues([]float64{4*x.Get(0)*x.Get(0)*x.Get(0) + 1, 2*x.Get(1)})
		fmt.Println("g = ", g)
		delta := optimizer.GetDeltaX(x, g)
		x.Increment(delta, 1)
		if g.Norm() < 0.0001 {
			break
		}
	}

	fmt.Println("==== done ====")
	fmt.Println("x = ", x)
}
