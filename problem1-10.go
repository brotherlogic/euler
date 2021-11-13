package main

import (
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
