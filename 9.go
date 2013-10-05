package main

import "fmt"

func main() {
    sum := 1000
    for i:=1; i <= sum; i++ {
        for j:=i+1; i+j <= sum; j++ {
            for k:=j+1; i+j+k <= sum; k++ {
                if i+j+k == sum  && (i*i+j*j) == k*k {
                    fmt.Println(i, j, k, i*j*k)
                }
            }
        }
    }
}
