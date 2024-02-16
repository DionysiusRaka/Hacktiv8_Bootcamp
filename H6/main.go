package main

import (
	"Hactiv8_Bootcamp/H6/goodies"
)

// func callbackExp(x []int, callback func(int) bool) int {
// 	var totalEvenNumber int

// 	for _, v := range x {
// 		fmt.Println("INI V ", v)
// 		if callback(v) {
// 			totalEvenNumber++
// 		}
// 	}
// 	return totalEvenNumber
// }

// func extractLastDigit(nums ...int) func() []int {
// 	newFunc := func() []int {
// 		var result []int
// 		for _, nums := range nums {
// 			lasNum := nums % 10
// 			if lasNum != 0 {
// 				result = append(result, lasNum)
// 			}
// 		}
// 		return result
// 	}
// 	return newFunc
// }

func main() {
	// var number = 0.123
	// var reflectNumber = reflect.ValueOf(number)
	// fmt.Println("Var type : ", reflectNumber.Type())
	// dumbFunction := extractLastDigit(23, 6, 889, 1230, 10102)
	// sl := []int{23, 21, 1, 3}
	// fmt.Println(dumbFunction())

	// callbackExp := callbackExp(sl, func(x int) bool {
	// 	return x%2 == 0
	// })
	// fmt.Println(callbackExp)

	goodies.Ref2Struct()
}
