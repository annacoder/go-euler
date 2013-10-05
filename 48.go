package main

import (
    "fmt"
    "math/big"
)

func main() {
    maxNum := 1000
    sum := big.NewInt(0)
    for i:=1; i <= maxNum; i++ {
        n := big.NewInt(int64(i))
        sum.Add(n.Exp(n, n, big.NewInt(0)), sum)
    }
    sumStr := sum.String()
    fmt.Println(sumStr[len(sumStr)-10:])
}
