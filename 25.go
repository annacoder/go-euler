package main

import (
    "fmt"
    "math/big"
)

func main() {
    a := new(big.Int).SetInt64(1)
    b := new(big.Int).SetInt64(1)
    for i:=1;;i++  {

        if len(a.String()) == 1000 { fmt.Println(i); break }

        tmp := new(big.Int).Set(b)
        b = b.Add(a, b)
        a.Set(tmp)
    }
}
