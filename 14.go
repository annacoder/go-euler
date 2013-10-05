package main

import "fmt"

type collatzLength map[int]int

func (c collatzLength) get(n int) int {
    if n == 1 {
        return 1
    }
    var prev int
    if n %2 == 0 {
        prev = n/2
    } else {
        prev = 3*n + 1
    }
    if l, ok := c[prev]; ok {
        c[n] = 1 + l
    } else {
         c[n] = 1 + c.get(prev)
    }
    return c[n]
}

func main() {
    c := make(collatzLength, 0)
    maxNum := 1000000
    maxStartingNum := 0
    maxChainLen := 0
    for i:=1; i < maxNum; i++ {
        chainLen := c.get(i)
        if chainLen > maxChainLen {
            maxChainLen = chainLen
            maxStartingNum = i
        }
    }
    fmt.Println(maxStartingNum)
}
