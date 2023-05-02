package main

import "time"

func main() {
	canal := make(chan int)

	go contar(canal)

	for num := range canal {
		println(num)
	}

}

func contar(canal chan int) {
	for i := 1; i <= 10; i++ {
		canal <- i
		time.Sleep(time.Second)
	}

	close(canal)
}