package main

import "fmt"

var f map[int]int

func cachedFactorial(n int) int {
    if v, ok := f[n]; ok {
        return v
    }
    f[n] = factorial(n)
    return f[n]
}

func factorial(n int) int {
    if (n <=1) { return 1 }
    return n * factorial(n-1)
}

func checkCurious(i int) bool {
    orig := i
    factSum := 0
    for i > 0 {
        n := i % 10
        i = i/10
        factSum += cachedFactorial(n)
    }
    if factSum == orig {
        return true
    }
    return false
}

func main() {
    f = make(map[int]int, 0)
    sum := 0
    for i:=3 ; i<cachedFactorial(9) ; i++ {
        if checkCurious(i) {
            sum += i
        }
    }
    fmt.Println(sum)
}
