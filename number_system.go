package mathics

import (
	"errors"
)

// NumberType represents a type of number system
type NumberType struct {
	Name        string
	Description string
}

// Predefined number systems
var (
	WholeNumbers = NumberType{
		Name:        "Whole Numbers",
		Description: "A number with no fractional or decimal part and cannot be negative",
	}
	NaturalNumbers = NumberType{
		Name:        "Natural Numbers",
		Description: "Whole numbers from 1 and up or some teacher say these are all the \"counting number\"",
	}
	Integers = NumberType{
		Name:        "Integers",
		Description: "All whole numbers (including positive and negative whole numbers)",
	}
	RationalNumbers = NumberType{
		Name:        "Rational Numbers",
		Description: "Any number that can be expressed as the quotient or fraction of two integers",
	}
	IrrationalNumbers = NumberType{
		Name:        "Irrational Numbers",
		Description: "Any real number that cannot be expressed as a ratio of integers",
	}
	RealNumbers = NumberType{
		Name:        "Real Numbers",
		Description: "All rational and irrational numbers",
	}
)

// CheckNumberSystem determines the type(s) of the given number
func CheckNumberSystem(i interface{}) (n []NumberType, e error) {
	switch v := i.(type) {
	case int:
		if IsWholeNumbers(v) {
			n = append(n, WholeNumbers)
		}
		if IsNaturalNumbers(v) {
			n = append(n, NaturalNumbers)
		}
		n = append(n, Integers)
	case float64:
		if IsRationalNumber(v) {
			n = append(n, RationalNumbers)
		} else {
			n = append(n, IrrationalNumbers)
		}
		n = append(n, RealNumbers)
	default:
		e = errors.New("your input is not a type of number")
	}
	return n, e
}

// IsWholeNumbers checks if the input is a whole number
func IsWholeNumbers(i int) bool {
	return i >= 0
}

// IsNaturalNumbers checks if the input is a natural number
func IsNaturalNumbers(i int) bool {
	return i > 0
}

// IsRationalNumber checks if the input is a rational number
func IsRationalNumber(f float64) bool {
	return f == float64(int(f))
}
