package main

import "fmt"
import "strconv"

// no encoding
func reverseBytes(s string) string {
    r := make([]byte, len(s))
    for i := 0; i < len(s); i++ {
        r[i] = s[len(s)-1-i]
    }
    return string(r)
}

func main() {
    min3digit := 100
    max3digit := 999

    minproduct := min3digit * min3digit
    maxproduct := max3digit * max3digit

    fmt.Println(minproduct, maxproduct)

    for i:=maxproduct; i>minproduct; i-- {
        s := fmt.Sprintf("%#v", i)
        if s[:len(s)/2] == reverseBytes(s[len(s)/2:]) {
            for j:=min3digit; j<max3digit; j++ {
                if i%j == 0 && len(strconv.FormatInt(int64(i/j), 10)) == 3 {
                        fmt.Println(i, j, i/j)
                        return
                }
            }
        }
    }
}
