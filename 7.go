package main

import "fmt"

func main() {
    nth := 10001
    primes := make([]int, 0)
    count := 2
    for {
        isPrime := true
        for _, v := range primes {
            if count % v == 0 {
                isPrime = false
                break
            }
        }
        if isPrime {
            primes = append(primes, count)
        }
        if len(primes) == nth {
            fmt.Println(count)
            break
        }
        count += 1
    }
}
