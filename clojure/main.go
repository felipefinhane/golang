package main

import (
	"fmt"
)

func main() {
	fmt.Println("Clojure")
	a := contador()
	fmt.Println("A ->", a())
	fmt.Println("A ->", a())
	fmt.Println("A ->", a())

	fmt.Println("")

	b := contador()
	fmt.Println("B ->", b())
	fmt.Println("B ->", b())
	fmt.Println("B ->", b())

	fmt.Println("")

	fmt.Println("A ->", a())
	fmt.Println("A ->", a())
	fmt.Println("A ->", a())

	fmt.Println("")

	fmt.Println("B ->", b())
	fmt.Println("B ->", b())
	fmt.Println("B ->", b())

}

func contador() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
