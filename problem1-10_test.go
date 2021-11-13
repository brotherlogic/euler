package main

import (
	"context"
	"testing"
)

func TestSolve1(t *testing.T) {
	s := InitTestServer()
	sol := s.solve1(context.Background(), 10)

	if sol != 23 {
		t.Errorf("Wrong answer 23 vs %v", sol)
	}
}
