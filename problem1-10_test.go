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

func TestSolve4(t *testing.T) {
	s := InitTestServer()
	sol := s.solve4(context.Background(), 2)

	if sol != 9009 {
		t.Errorf("Wrong answer: %v vs %v", 9009, sol)
	}
}

func TestSolve5(t *testing.T) {
	s := InitTestServer()
	sol := s.solve5(context.Background(), 10)

	if sol != 2520 {
		t.Errorf("Wrong answer: %v vs %v", 2520, sol)
	}
}

func TestSolve6(t *testing.T) {
	s := InitTestServer()
	sol := s.solve6(context.Background(), 10)

	if sol != 2640 {
		t.Errorf("Wrong answer: %v vs %v", 2640, sol)
	}
}

func TestSolve7(t *testing.T) {
	s := InitTestServer()
	sol := s.solve7(context.Background(), 6)

	if sol != 13 {
		t.Errorf("Wrong answer: %v vs %v", 13, sol)
	}
}

func TestSolve8(t *testing.T) {
	s := InitTestServer()
	sol := s.solve8(context.Background(), 4)

	if sol != 5832 {
		t.Errorf("Wrong answer: %v vs %v", 5832, sol)
	}
}

func TestSolve9(t *testing.T) {
	s := InitTestServer()
	sol := s.solve9(context.Background(), 3+4+5)

	found := false
	for _, so := range sol {
		if so == 3*4*5 {
			found = true
		}
	}

	if !found {
		t.Errorf("Could not find solution: %v", sol)
	}
}

func TestSolve10(t *testing.T) {
	s := InitTestServer()
	sol := s.solve10(context.Background(), 10)

	if sol != 17 {
		t.Errorf("Wrong answer: %v vs %v", 17, sol)
	}
}
