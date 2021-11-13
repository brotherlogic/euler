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

func TestSolve2(t *testing.T) {
	s := InitTestServer()
	sol := s.solve2(context.Background(), 100)

	if sol != 2+8+34 {
		t.Errorf("Wong answer: %v vs %v", 2+8+34, sol)
	}
}

func TestSolve3(t *testing.T) {
	s := InitTestServer()
	sol := s.solve3(context.Background(), 13195)

	if sol != 29 {
		t.Errorf("Wrong answer: %v vs %v", 29, sol)
	}
}
