package main

import (
	"fmt"
	"github.com/supanadit/mathics"
)

func main() {
	d := mathics.CheckNumberSystem(50.123)
	for _, v := range d {
		fmt.Println(v.Name)
	}

	//fmt.Println(mathics.GCDT(2625, 1000))
	//fmt.Println(mathics.GCDT(336, 360))
	//fmt.Println(mathics.GCDM([]int{5, 10, 15, 25}))
	fmt.Println(mathics.LCMM([]int{100, 90, 80, 7}))
	fmt.Println(mathics.LCMM([]int{5, 10, 15, 25}))
	//fmt.Println(mathics.LCDT(3, 15))
	//fmt.Println(mathics.LCDT(5, 15))
}
