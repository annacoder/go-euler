package main

import "fmt"

type Factor struct {
    factors map[int]map[int]int
    primes []int
}

func NewFactor() *Factor {
    return &Factor{make(map[int]map[int]int, 0), make([]int, 0)}
}

func (f *Factor) getFactors(n int) map[int]int {
    if n == 1 { return map[int]int{} }

    if _, ok:= f.factors[n]; ok {
        return f.factors[n]
    }

    isPrime := true
    primeFactor := 0
    for _, p := range f.primes {
        if n % p == 0 {
            primeFactor = p
            isPrime = false
            break
        }
    }

    if isPrime {
        f.primes = append(f.primes, n)
        f.factors[n] = map[int]int{n : 1}
        return f.factors[n]
    } else {
        f := f.getFactors(n / primeFactor)
        r := make(map[int]int, 0)
        for k,v := range f { r[k] = v }
        if _, ok := r[primeFactor]; ok {
            r[primeFactor] += 1
        } else {
            r[primeFactor] = 1
        }
        return r
    }
}

func (f *Factor) setFactors(n int, factorsList ...map[int]int) {
    if _, ok := f.factors[n]; !ok {
        f.factors[n] = make(map[int]int, 0)
    } else {
        return
    }

    for _, factors := range factorsList {
        for k,v := range factors {
            if _, ok := f.factors[n][k]; ok {
                f.factors[n][k] += v
            } else {
                f.factors[n][k]= v
            }
        }
    }
}

func (f *Factor) numFactors(n int) int {
    factors := f.factors[n]
    num := 1
    for _,v := range factors {
        num *= (1+v)
    }
    return num
}

func main() {
    f := NewFactor()
    maxFactors := 500
    odd := 1
    checkNumFactors := func (f *Factor, i,odd int) bool {
        f.setFactors(i*odd, f.getFactors(i), f.getFactors(odd))
        numFactors := f.numFactors(i*odd)
        if numFactors > maxFactors { return true }
        return false
    }

    for i:=1; ; i++ {
        if checkNumFactors(f, i, odd) { fmt.Println(i*odd); break }
        odd += 2
        if checkNumFactors(f, i, odd) { fmt.Println(i*odd); break }
    }
}
