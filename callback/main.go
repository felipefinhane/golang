package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(soma()(mySlice...))
	fmt.Println(somentePar(soma(), mySlice...))
	fmt.Println(somenteImpar(soma(), mySlice...))
}

func soma() func(x ...int) int {
	return func(x ...int) int {
		total := 0
		for _, v := range x {
			total += v
		}
		return total
	}
}

func somentePar(soma func(x ...int) int, s ...int) (total int) {
	newSlice := []int{}
	for _, v := range s {
		if v%2 == 0 {
			newSlice = append(newSlice, v)
		}
	}
	return soma(newSlice...)
}

func somenteImpar(soma func(x ...int) int, s ...int) (total int) {
	newSlice := []int{}
	for _, v := range s {
		if v%2 != 0 {
			newSlice = append(newSlice, v)
		}
	}
	return soma(newSlice...)
}
