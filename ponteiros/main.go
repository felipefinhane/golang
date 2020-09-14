package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	x := 10
	fmt.Println("x ->", x)

	fmt.Println("")

	y := &x
	fmt.Println("Y (&x) ->", y)
	fmt.Println("*Y (&x) ->", *y)

	fmt.Println("")

	*y = 15
	fmt.Println("x ->", x)
	fmt.Println("Y (&x) ->", y)
	fmt.Println("*Y (&x) ->", *y)

	fmt.Println("")
	fmt.Println("Valor de X ->", x)
	fmt.Println("")
	incrementoNormal(x)
	fmt.Println("[Incremento Normal] Valor de X ->", x)
	fmt.Println("Valor de X ->", x)
	fmt.Println("")
	incrementoPonteiro(&x)
	fmt.Println("[Incremento Ponteiro] Valor de X ->", x)
	fmt.Println("Valor de X ->", x)
}

func incrementoNormal(valor int) {
	valor++
	fmt.Println("Na funcão valor ->", valor)
}

func incrementoPonteiro(valor *int) {
	*valor++
	fmt.Println("Na funcão valor ->", *valor)
}
