package main

import "fmt"

func Array2() {
	myName := "dionysius"

	var arr [4][5]string = [4][5]string{
		{"alvin", "arif", "reza", "rinaldi", "nina"},      //0
		{"noel", "dilla", "rosa", "juan michel", "teguh"}, //1
		{"septyan", "farras", "giyanda", "afin", "azwar"}, //2
		{"dionysius", "rifki", "raffi", "zain"},           //3
	}

	for firstIndex, firstValue := range arr {
		for secondIndex, secondValue := range firstValue {
			if secondValue == myName {
				fmt.Println(firstIndex)
				fmt.Println(firstValue)
				fmt.Println(secondIndex)
				fmt.Println(secondValue)
			}
		}
	}
}
