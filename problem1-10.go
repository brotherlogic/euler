package main

import "golang.org/x/net/context"

func (s *Server) solve1(ctx context.Context, max int) int {
	sum := 0
	for i := 3; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}
