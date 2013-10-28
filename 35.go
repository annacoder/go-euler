package main

import "fmt"
import "math/big"
import "math"
import "strconv"

func primeSieve(n int) []int {
	x := big.NewInt(0)

	limit := int(math.Sqrt(float64(n)))

	for i := 4; i < n; i += 2 {
		x.SetBit(x, i, 1)
	}

	for i := 3; i <= limit; i += 2 {
		if x.Bit(i) == 0 {
			for j := i * i; j < n; j += 2 * i {
				x.SetBit(x, j, 1)
			}
		}
	}

	primes := make([]int, 0)
	for i := 2; i < n; i++ {
		if x.Bit(i) == 0 {
			primes = append(primes, i)
		}
	}

	return primes
}
func main() {

	maxNum := 1000000

	primes := make(map[string]int, 1)
	circularPrimes := make([]string, 0)

	for _, v := range primeSieve(maxNum) {
		primes[strconv.Itoa(v)] = 1
	}

	for k := range primes {
		isCircular := true
		for i := 0; i < len(k); i++ {
			k = k[1:] + string(k[0])
			if _, ok := primes[k]; !ok {
				isCircular = false
				break
			}
		}
		if isCircular {
			circularPrimes = append(circularPrimes, k)
		}
	}

	fmt.Println(len(circularPrimes))
}
