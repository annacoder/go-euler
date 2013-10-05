package main

import "fmt"
import "math/big"

func main() {
    i := big.NewInt(2)
    j := big.NewInt(1000)
    z := new(big.Int)
    z.Exp(i, j, big.NewInt(0))
    sum := 0
    for _, v := range z.String() {
        sum += int(v - '0')
    }
    fmt.Println(sum)
}
