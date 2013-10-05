package main

import "fmt"

func main() {
    nums := make([]int, 20)
    for i := range nums {
        nums[i] = i+1
    }
    for i, v := range nums {
        for j:=i+1; j < len(nums); j++ {
            if nums[j] % v == 0 {
                nums[j] = nums[j]/v
            }
        }
    }

    prod := 1
    for _, v := range nums {
        prod *= v
    }
    fmt.Println(prod)
}
