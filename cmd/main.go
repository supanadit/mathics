package main

import (
	"fmt"

	"github.com/supanadit/mathics"
)

func main() {
	d, _ := mathics.CheckNumberSystem(-1)
	for _, v := range d {
		fmt.Println(v.Name)
	}

	//fmt.Println(mathics.GCDT(2625, 1000))
	//fmt.Println(mathics.GCDT(2664, 999))
	fmt.Println(mathics.GCDT(12, 20))
	fmt.Println(mathics.GCDM([]int{1, 2, 3, 4, 6, 12}))
	//fmt.Println(mathics.LCMM([]int{100, 90, 80, 7}))
	// fmt.Println(mathics.LCMM([]int{1, 2, 3, 5, 6, 10, 15, 30}))
	// fmt.Println(mathics.LCMT(3, 15))
	// fmt.Println(mathics.LCMT(5, 7))
	// fmt.Println(mathics.LCMT(10, 11))
	//fmt.Println(mathics.LCMT(5, 15))

	// Polynomial example
	polynomial := mathics.Polynomial{
		Coefficients: []float64{2, -3, 1}, // represents 2x² - 3x + 1
	}

	// Evaluate polynomial at x = 2
	result := mathics.EvaluatePolynomial(polynomial, 2)
	fmt.Printf("Polynomial 2x² - 3x + 1 evaluated at x = 2: %.2f\n", result)

	// Quadratic equation example
	// Solve: 1x² + 5x + 6 = 0
	root1, root2, hasRealRoots := mathics.QuadraticSolver(1, 5, 6)
	if hasRealRoots {
		fmt.Printf("Roots of x² + 5x + 6 = 0 are: %.2f and %.2f\n", root1, root2)
	} else {
		fmt.Println("No real roots exist")
	}

	// Another quadratic example
	// Solve: 1x² + 2x + 1 = 0 (perfect square)
	root1, root2, hasRealRoots = mathics.QuadraticSolver(1, 2, 1)
	if hasRealRoots {
		fmt.Printf("Roots of x² + 2x + 1 = 0 are: %.2f and %.2f\n", root1, root2)
	} else {
		fmt.Println("No real roots exist")
	}
}
