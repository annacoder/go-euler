package main

import "fmt"

func main() {
    input := 600851475143
    num := input
    maxPrime := 0
    for i:=2;i<input/2; i++ {
        if (i > num) { break }
        if num%i ==0 {
            maxPrime = i
            num /= i
        }
    }

    fmt.Println(maxPrime)
}
