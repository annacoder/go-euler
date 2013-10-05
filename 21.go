package main

import "fmt"
import "math"

type Factor struct {
    factors map[int]map[int]int
    primes []int
}

func NewFactor() *Factor {
    return &Factor{make(map[int]map[int]int, 0), make([]int, 0)}
}

func (f *Factor) getFactors(n int) map[int]int {
    if n == 1 { return map[int]int{1: 1} }

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
        f.factors[n] = map[int]int{n : 1} //, 1: 1}
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


func (f *Factor) listFactors(n int) []int {
    factorsList := make([][]int, 0)
    for  k, v:= range(f.getFactors(n)) {
        factors := make([]int, 0)
        for i:=1; i <= v; i++ {
            factors = append(factors, int(math.Pow(float64(k), float64(i))))
        }
        factorsList = append(factorsList, factors)
    }

    factors := []int{1}
    for ; len(factorsList) != 0 ; {
        factorsGroup := factorsList[0]
        factorsList = factorsList[1:]
        tmpFactors := make([]int, len(factors))
        copy(tmpFactors, factors)
        for _, factor := range(factorsGroup) {
            for _, i := range(tmpFactors) {
                factors = append(factors, i * factor)
            }
        }
    }

    return factors
}

func (f *Factor) sumFactors(n int) int {

    factors := f.listFactors(n)
    sum := 0
    for _, v := range factors {
        if(v != n) {
            sum += v
        }
    }
    return sum
}

func main() {
    f := NewFactor()
    maxNum := 10000
    D := map[int]int{}
    amicable := map[int]int{}

    for i:=1; i < maxNum; i++ {
        D[i] = f.sumFactors(i)
    }

    for a, b := range(D) {
        if b1, ok := D[b]; ok && b1 == a && b!=a{
            if _, ok := amicable[a]; !ok {
                amicable[a] = 1
            }
            if _, ok := amicable[b]; !ok {
                amicable[b] = 1
            }
        }
    }
    sum := 0
    for k, _ := range(amicable) {
        sum += k
    }
    fmt.Println(sum)
}
