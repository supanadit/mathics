package mathics

import (
	"errors"
	"fmt"
	"reflect"
)

// WholeNumbers is type of number system
var WholeNumbers NumberType = NumberType{
	Name:        "Whole Numbers",
	Description: "A number with no fractional or decimal part and cannot be negative",
}

// NaturalNumbers is type of number system
var NaturalNumbers NumberType = NumberType{
	Name:        "Natural Numbers",
	Description: "Whole numbers from 1 and up or some teacher say these are all the \"counting number\"",
}

// Integers is type of number system
var Integers NumberType = NumberType{
	Name:        "Integers",
	Description: "All whole numbers (including positive and negative whole numbers)",
}

// NumberType is the struct to grouping some known type of number
type NumberType struct {
	Name        string
	Description string
}

// CheckNumberSystem is the function for checking which type of your number really is
func CheckNumberSystem(i interface{}) (n []NumberType, e error) {
	switch i.(type) {
	case int:
		d := int(reflect.ValueOf(i).Int())
		if IsWholeNumbers(d) {
			n = append(n, WholeNumbers)
		}

		if IsNaturalNumbers(d) {
			n = append(n, NaturalNumbers)
		}

		n = append(n, Integers)
		break
	case float64:
		fmt.Println("On working for rational numbers, irrational numbers and real numbers")
		break
	default:
		e = errors.New("your input is not a type of number")
		break
	}
	return n, e
}

// IsWholeNumbers is function for knowing whether the input is whole numbers or nor
func IsWholeNumbers(i int) bool {
	return i >= 0
}

// IsNaturalNumbers is function for knowing whether the input is natural numbers or nor
func IsNaturalNumbers(i int) bool {
	return i > 0
}
