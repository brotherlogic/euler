package main

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n < 2 {
		return false
	}
	for i := 3; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func primeFactors(n int) []int {
	var nums []int

	for i := 2; i < n; i++ {
		if n%i == 0 && isPrime(i) {
			nums = append(nums, i)
		}
	}

	return nums
}
