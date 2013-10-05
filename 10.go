package main

import "fmt"

func main() {
    primes := make([]int, 0)
    primeLimit := 2000000
    primeSum := 0
    count := 2
    for {
        isPrime := true
        for _,v := range primes {
            if count%v == 0 {
                isPrime = false
                break
            }
        }
        if isPrime {
            primes = append(primes, count)
            primeSum += count
        }
        if count >= primeLimit {
            fmt.Println(primeSum)
            break
        }
        if count% 200000 == 0 {
            fmt.Println(count)
        }
        count += 1
    }
}
