package mathics

import "math"

// GCDT calculates the Greatest Common Divisor (GCD) of two integers using the Euclidean algorithm.
// It returns the GCD of the two input integers.
func GCDT(a int, b int) int {
	// Ensure the inputs are non-negative
	a = int(math.Abs(float64(a)))
	b = int(math.Abs(float64(b)))

	// Euclidean algorithm to find GCD
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

// GCDM is Greatest Common factor function for more than two number
func GCDM(i []int) int {
	l := len(i)
	if l != 0 {
		a := i[0]
		for x := 1; x < l; x++ {
			b := i[x]
			a = GCDT(a, b)
		}
		return a
	} else {
		return 0
	}
}

// LCMT is Least Common Multiple function for two number
func LCMT(a int, b int) int {
	return (a * b) / GCDT(a, b)
}

// LCMM is Least Common Multiple function for more than two number
func LCMM(i []int) int {
	if len(i) == 0 {
		return 0
	}

	result := i[0]
	for x := 1; x < len(i); x++ {
		result = LCMT(result, i[x])
	}
	return result
}

func ExtendedGCD(a int, b int) (int, int, int) {
	// Ensure the inputs are non-negative
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	old_r, r := a, b
	old_s, s := 1, 0
	old_t, t := 0, 1

	for r != 0 {
		quotient := old_r / r
		old_r, r = r, old_r-quotient*r
		old_s, s = s, old_s-quotient*s
		old_t, t = t, old_t-quotient*t
	}

	return old_r, old_s, old_t
}
