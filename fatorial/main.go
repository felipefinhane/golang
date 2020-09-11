package main

import "fmt"

func main() {
	x := 0
	fmt.Println("Recursive ->\t", fatore(x))
	fmt.Println("Loop ->\t", fatoreWirtLoop(x))
	x = 1
	fmt.Println("Recursive ->\t", fatore(x))
	fmt.Println("Loop ->\t", fatoreWirtLoop(x))
	x = 2
	fmt.Println("Recursive ->\t", fatore(x))
	fmt.Println("Loop ->\t", fatoreWirtLoop(x))
	x = 3
	fmt.Println("Recursive ->\t", fatore(x))
	fmt.Println("Loop ->\t", fatoreWirtLoop(x))
	x = 7
	fmt.Println("Recursive ->\t", fatore(x))
	fmt.Println("Loop ->\t", fatoreWirtLoop(x))
}

func fatore(x int) int {
	if x == 0 {
		return 1
	}
	if x == 1 {
		return x
	} else {
		return x * fatore(x-1)
	}
}

func fatoreWirtLoop(x int) int {
	resultado := 1
	for x > 1 {
		resultado *= x
		x--
	}
	return resultado
}
