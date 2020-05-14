package mathics

// GCDT is Greatest Common Factor function for two number
func GCDT(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
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
	l := len(i)
	a := 0
	b := 0

	for x := 0; x < (l - 1); x++ {
		a = i[x] % i[x+1]
		if a == 0 {
			i[x+1] = (i[x] * i[x+1]) / i[x+1]
		} else {
			b = i[x+1] % a
			if b == 0 {
				i[x+1] = (i[x] * i[x+1]) / a
			} else {
				i[x+1] = (i[x] * i[x+1]) / b
			}
		}
	}
	return i[l-1]
}
