package main

import (
	"math"
	"strconv"

	"golang.org/x/net/context"
)

func (s *Server) solve1(ctx context.Context, max int) int {
	sum := 0
	for i := 3; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

func (s *Server) solve2(ctx context.Context, max int) int {
	s1 := 1
	s2 := 2
	sum := 2

	for s2 < max {
		s1, s2 = s2, s1+s2
		if s2%2 == 0 && s2 < max {
			sum += s2
		}
	}

	return sum
}

func (s *Server) solve3(ctx context.Context, val int64) int64 {
	nums := primeFactors(val)
	return nums[len(nums)-1]
}

func (s *Server) solve4(ctx context.Context, val int64) int64 {
	start := math.Pow10(int(val - 1))
	end := math.Pow10(int(val))

	largest := int64(-1)
	for num1 := start; num1 < end; num1++ {
		for num2 := start; num2 < end; num2++ {
			res := int64(num1 * num2)
			if res > largest && isPalindrome(res) {
				largest = res
			}
		}
	}

	return largest
}

func (s *Server) solve5(ctx context.Context, val int64) int64 {
	mval := int64(1)
	for i := int64(1); i <= val; i++ {
		mval *= i
	}
	for i := int64(1); i < mval; i++ {
		found := true
		for j := int64(2); j < val; j++ {
			if i%j != 0 {
				found = false
				break
			}
		}

		if found {
			return i
		}
	}

	return -1
}

func (s *Server) solve6(ctx context.Context, val int64) int64 {
	ssq := int64(0)
	sum := int64(0)
	for i := int64(1); i <= val; i++ {
		ssq += i * i
		sum += i
	}

	return sum*sum - ssq
}

func (s *Server) solve7(ctx context.Context, val int64) int64 {
	primes := getPrimes(21474836)
	return primes[val-1]
}

func (s *Server) solve8(ctx context.Context, digs int) int64 {
	num := "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"

	bProd := int64(1)
	for i := 0; i < len(num)-digs; i++ {
		substr := num[i : i+digs]
		prod := int64(1)
		for _, r := range substr {
			num, err := strconv.ParseInt(string(r), 10, 16)
			if err != nil {
				return -1
			}
			prod *= num
		}

		if prod > bProd {
			bProd = prod
		}
	}
	return bProd
}
