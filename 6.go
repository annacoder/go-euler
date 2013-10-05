package main

import "fmt"

func main() {
    num := 100
    sum_of_squares, squares_of_sum := 0,0
    for i:=1; i<=num; i++ {
        sum_of_squares += i*i
        squares_of_sum += i
    }
    squares_of_sum *= squares_of_sum
    fmt.Println(squares_of_sum - sum_of_squares)
}
