package main

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

	for i := int64(2); i < n; i++ {
		if n%i == 0 && isPrime(i) {
			nums = append(nums, i)
		}
	}

	return nums
}
