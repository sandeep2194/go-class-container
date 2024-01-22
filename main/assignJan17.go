package main

import "fmt"

type NumFinder struct{}

func (nf NumFinder) findEven(a []int) []int {
	var evens []int
	for _, num := range a {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	return evens
}

func assignJan17() {
	nf := NumFinder{}
	sampleArray := []int{1, 2, 358, 412, 502, 6, 7, 8, 9, 10, 907, 291}
	evenNumbers := nf.findEven(sampleArray)
	fmt.Println("Even Numbers:", evenNumbers)
}
