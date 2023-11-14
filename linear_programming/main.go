package main

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/optimize"
	"gonum.org/v1/gonum/optimize/functions"
)

func main() {
	p := optimize.Problem{
		Func: functions.ExtendedRosenbrock{}.Func,
		Grad: functions.ExtendedRosenbrock{}.Grad,
	}

	x := []float64{1.3, 0.7, 0.8, 1.9, 1.2}
	result, err := optimize.Minimize(p, x, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = result.Status.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result.Status: %v\n", result.Status)
	fmt.Printf("result.X: %0.4g\n", result.X)
	fmt.Printf("result.F: %0.4g\n", result.F)
	fmt.Printf("result.Stats.FuncEvaluations: %d\n", result.Stats.FuncEvaluations)
}
