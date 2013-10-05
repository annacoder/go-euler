package main

import (
    "fmt"
    "bufio"
    "os"
    "sort"
)

func sumChars(s string) (sum int) {
    for _, v := range s {
        sum += int(v - '@')
    }
    return
}

func main() {
    f, _ := os.Open("names.txt")
    scanner := bufio.NewScanner(f)
    names := []string{}
    total := 0
    for scanner.Scan() {
        s := scanner.Text()
        names = append(names, s)
    }
    sort.Sort(sort.StringSlice(names))
    for i, name := range names {
        rank := i+1
        score := sumChars(name)
        total += (score * rank)
    }
    fmt.Println(total)
}
