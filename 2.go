package main

import "fmt"

func fibo(c chan int) {
    a, b := 1, 1
    for {
        a, b = b, a+b
        if a > 4000000 {
            close(c)
            return
        }
        if a %2 == 0 {
            c <- a
        }
    }
}

func main() {
    c := make(chan int)
    go fibo(c)
    sum := 0

    for v := range c {
        sum += v
    }

    fmt.Println(sum)
}
