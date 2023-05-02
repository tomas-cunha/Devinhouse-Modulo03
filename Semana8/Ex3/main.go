package main

import "fmt"

func main() {
    var n int
    fmt.Print("Digite um número inteiro: ")
    fmt.Scanln(&n)

    primos := make([]bool, n+1)
    for i := 2; i <= n; i++ {
        primos[i] = true
    }

    for i := 2; i*i <= n; i++ {
        if primos[i] {
            for j := i * i; j <= n; j += i {
                primos[j] = false
            }
        }
    }

    fmt.Println("Números primos até", n, ":")
    for i := 2; i <= n; i++ {
        if primos[i] {
            fmt.Println(i)
        }
    }
}