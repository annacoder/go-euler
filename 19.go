package main

import (
    "fmt"
    "time"
)

func main() {
    start := time.Date(1901, time.January, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2000, time.December, 31, 0, 0, 0, 0, time.UTC)

    numSundays := 0

    current := start
    for current.Before(end) {
        if (current.Weekday() == time.Sunday) {
            numSundays += 1
        }
        current = current.AddDate(0, 1, 0)
    }
    fmt.Println(numSundays)
}
