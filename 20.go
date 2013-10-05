package main

import "fmt"
import "math/big"

func main() {
    num := new(big.Int)
    num = num.MulRange(1, 100)
    sum := 0
    for _, v := range num.String() {
        sum += int(v - '0')
    }
    fmt.Println(sum)
}
