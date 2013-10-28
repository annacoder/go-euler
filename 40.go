package main

import (
    "fmt"
    "strconv"
)

func main() {
    maxNum := 1000000
    digitsSeen := 0
    nextDigit := 1
    result := 1
    for i := 1; i <= maxNum; i++ {
        n := strconv.Itoa(i)
        if digitsSeen + len(n) >= nextDigit {
            num, _ := strconv.Atoi(string(n[(nextDigit-digitsSeen)-1]))
            result *= num
            nextDigit *= 10
        }
        digitsSeen += len(n)
    }
    fmt.Println(result)
}
