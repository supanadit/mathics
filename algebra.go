package mathics

// Polynomial represents a polynomial expression with coefficients
type Polynomial struct {
	Coefficients []float64 // Ordered from highest degree to constant term
}

// EvaluatePolynomial calculates the value of a polynomial for a given x
func EvaluatePolynomial(p Polynomial, x float64) float64 {
	result := 0.0
	for i, coef := range p.Coefficients {
		power := len(p.Coefficients) - i - 1
		term := coef
		for j := 0; j < power; j++ {
			term *= x
		}
		result += term
	}
	return result
}

// QuadraticSolver solves quadratic equations in the form ax² + bx + c = 0
// Returns the two roots and a boolean indicating if the roots are real
func QuadraticSolver(a, b, c float64) (float64, float64, bool) {
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return 0, 0, false
	}

	sqrtDisc := discriminant
	if discriminant > 0 {
		sqrtDisc = discriminant
	}

	root1 := (-b + sqrtDisc) / (2 * a)
	root2 := (-b - sqrtDisc) / (2 * a)

	return root1, root2, true
}
