package main

import (
	"fmt"
)

func primerDiagrama() int {
	var sum int = 0
	for i := 1; i <= 100; i++ {

		sum += 5 * i

	}
	return sum
}
func segundoDiagrama() int {
	var sum int
	for i := 2; i <= 100; i += 2 {
		sum += i
	}
	return sum
}
func tercerDiagrama() []int {
	var sum []int
	for i := 1; i <= 300; i++ {
		if i%2 == 0 {
			sum = append(sum, i)
		}
	}
	return sum
}
func main() {

	fmt.Println(primerDiagrama())
	fmt.Println(segundoDiagrama())
	fmt.Println(tercerDiagrama())
}
