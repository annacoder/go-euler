package main

import "fmt"
import "math/big"

func main() {
     maxA, maxB := 100, 100

     seq := make(map[string]int, 0)
     for i:= 2; i <= maxA; i++ {
         for j:=2; j <= maxB; j++ {
              z := big.NewInt(0)
              z.Exp(big.NewInt(int64(i)), big.NewInt(int64(j)), big.NewInt(0))
              seq[z.String()] = 1
          }
     }
     fmt.Println(len(seq))
}

