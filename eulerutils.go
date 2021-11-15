package main

import (
	"math"
	"strconv"
)

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

func isPalindrome(num int64) bool {
	str := strconv.FormatInt(num, 10)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}

	return true
}

func getPrimes(max int64) []int64 {
	var primes []int64

	countarr := make([]bool, max+1)
	for i := 2; i < len(countarr); i++ {
		if !countarr[i] {
			primes = append(primes, int64(i))
			for j := i; j < len(countarr); j += i {
				countarr[j] = true
			}
		}
	}

	return primes
}
