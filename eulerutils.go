package main

import "math"

func isPrime(n int64) bool {
	if n == 2 {
		return true
	}
	if n < 2 {
		return false
	}
	for i := int64(3); i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func primeFactors(n int64) []int64 {
	var nums []int64
	first := true

	for n%2 == 0 {
		if first {
			first = false
			nums = append(nums, 2)
		}
		n = n / 2
	}

	for i := int64(3); i < int64(math.Sqrt(float64(n))); i += 2 {
		first = true
		for n%i == 0 {
			if first {
				first = false
				nums = append(nums, i)
			}
			n = n / i
		}
	}

	if n > 2 {
		nums = append(nums, n)
	}

	return nums
}
