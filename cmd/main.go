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
	//fmt.Println(mathics.GCDT(336, 360))
	//fmt.Println(mathics.GCDM([]int{5, 10, 15, 25}))
	//fmt.Println(mathics.LCMM([]int{100, 90, 80, 7}))
	fmt.Println(mathics.LCMM([]int{1, 2, 3, 5, 6, 10, 15, 30}))
	fmt.Println(mathics.LCMT(3, 15))
	fmt.Println(mathics.LCMT(5, 7))
	fmt.Println(mathics.LCMT(10, 11))
	//fmt.Println(mathics.LCMT(5, 15))
}
