package main

import "fmt"

func main() {
	var numbers []int

	for i := 0; i <= 100; i++ {
		numbers = append(numbers, i)
	}

	var primeNumbers []int

	for i := 0; i < len(numbers); i++ {
		isPrime := true

		for j := 2; j <= numbers[i]/2; j++{
			if numbers[i] % j == 0{
				isPrime = false
				break
			}
		}

		if isPrime{
			primeNumbers = append(primeNumbers, numbers[i])
		}

	}

	fmt.Print("Números primos: ", primeNumbers)
}