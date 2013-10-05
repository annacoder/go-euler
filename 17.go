package main

import "fmt"
import "strconv"

func main() {
    digitMap := map[int]string{
        1 : "one",
        2 : "two",
        3 : "three",
        4 : "four",
        5 : "five",
        6 : "six",
        7 : "seven",
        8 : "eight",
        9 : "nine",
        10: "ten",
        11: "eleven",
        12: "twelve",
        13: "thirteen",
        14: "fourteen",
        15: "fifteen",
        16: "sixteen",
        17: "seventeen",
        18: "eighteen",
        19: "nineteen",
    }
    tensMap := map[int]string {
        2 : "twenty",
        3 : "thirty",
        4 : "forty",
        5 : "fifty",
        6 : "sixty",
        7 : "seventy",
        8 : "eighty",
        9 : "ninety",
    }

    count := 0
    for i:=1; i<=1000; i++ {
        if i==1000 {
            count += len("onethousand")
            continue
        }

        if _, ok := digitMap[i]; ok {
            count += len(digitMap[i])
            continue
        }

        s := strconv.Itoa(i)
        tmp := 0
        if len(s) == 3 {
            tmp, _ = strconv.Atoi(s[1:])
        }
        str := ""
        if _, ok := digitMap[tmp]; ok {
            str += digitMap[int(s[0] - '0')]
            str += "hundredand"
            str += digitMap[tmp]
        } else {
            for j, v := range s {
                digit := int(v - '0')
                if digit == 0 { continue }
                if (j == 0 && len(s) == 3) || (j == len(s)-1) {
                    str += digitMap[digit]
                }
                if j == 0 && len(s) == 3 {
                    str += "hundred"
                    if tmp!=0 {
                        str += "and"
                    }
                }
                if (j == 1 && len(s) == 3) || (j == 0 && len(s) == 2) {
                    str += tensMap[digit]
                }
            }
        }
        count += len(str)
    }
    fmt.Println(count)
}
