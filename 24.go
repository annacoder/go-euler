package main

import "fmt"
import "strconv"

func factorial(n int) int {
    if n == 1 { return 1 }
    return n * factorial(n-1)
}

func getPerm(digits []int, nthPerm int) []int{

    numDigitsLeft := len(digits)
    if numDigitsLeft == 1 {
        return digits
    }

    permsPerDigit := factorial(numDigitsLeft-1)
    index := nthPerm / permsPerDigit
    chosenDigit := digits[index]
    remaining := make([]int, 0)
    for i, v := range digits {
        if i != index {
            remaining = append(remaining, v)
        }
    }
    suffix := getPerm(remaining, nthPerm - (permsPerDigit * index))
    result := []int{chosenDigit}
    return append(result, suffix...)
}

func main() {

    const numDigits int = 10

    digits := make([]int, numDigits)

    for i:=0; i<numDigits; i++ {
        digits[i] = i
    }

    nthPerm := 1000000
    var buf []byte
    for _, v := range getPerm(digits, nthPerm-1) {
        buf = strconv.AppendInt(buf, int64(v), 10)
    }
    fmt.Println(string(buf))
}
