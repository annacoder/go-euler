package main

import "fmt"

func getDivisible(num int, c chan int) {
	if num%3 == 0 || num%5 == 0 {
		c <- num
	} else {
		c <- 0
	}
}
func main() {
    c := make(chan int, 1000)
    sum := 0
    for i:=0; i<1000; i++ {
	go getDivisible(i, c)
    }
    for i:=0; i<1000; i++ {
	select {
		case s := <-c: {
			sum+= s
		}
	}
    }
    fmt.Println(sum)
}
