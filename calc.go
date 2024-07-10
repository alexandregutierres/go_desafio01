package main

import "fmt"

func main() {
	x := somar(1, 3, 4)

	fmt.Println(x)
}

func somar(i ...int) int {
	total := 0

	for _, v := range i {
		total += v
	}

	return total
}

func subtrair(i ...int) int {
	total := 0

	for _, v := range i {
		total = v - total
	}

	return total
}

func multiplicar(i ...int) int {
	total := 1

	for _, v := range i {
		total = v * total
	}

	return total
}

func dividir(i ...int) int {
	total := 1

	for _, v := range i {
		total = v / total
	}

	return total
}
