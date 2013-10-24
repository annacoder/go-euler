package main

import "strconv"
import "fmt"

func isPalindrome(s string) bool {
    var a, b string
    if len(s)%2 == 0 {
        a = s[:len(s)/2]
        b = s[len(s)/2:]
    } else {
        a = s[:len(s)/2]
        b = s[(len(s)/2)+1:]
    }

    isPalindrome := true

    for i := range a {
        c := b[len(a)-1-i]
        if a[i] != c {
            isPalindrome = false
            break
        }
    }
    return isPalindrome
}

func reverse(s string) string {
    // Get Unicode code points.
    n := 0
    Rune := make([]rune, len(s))
    for _, r := range s {
        Rune[n] = r
        n++
    }
    Rune = Rune[0:n]
    // Reverse
    for i := 0; i < n/2; i++ {
        Rune[i], Rune[n-1-i] = Rune[n-1-i], Rune[i]
    }
    // Convert back to UTF-8.
    return string(Rune)
}

func is_palindrome(s string) bool {
        return s == reverse(s)
}


func main() {
    maxNum := 1000000
    sum := 0

    for i:=1; i< maxNum; i++ {
        iStr := strconv.FormatInt(int64(i), 10)
        binStr := strconv.FormatInt(int64(i), 2)
        if is_palindrome(iStr) && is_palindrome(binStr) {
            sum += i
        }
    }
    fmt.Println(sum)
}
