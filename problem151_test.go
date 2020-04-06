package main

import "testing"

func Test4(t *testing.T) {
	tv, s := solve151(0, []int{4})

	if tv != 0 {
		t.Errorf("Bad pulls: %v", tv)
	}

	if s != 0 {
		t.Errorf("Bad singles: %v", s)
	}
}

func Test3(t *testing.T) {
	tv, s := solve151(0, []int{3})

	if tv != 3 {
		t.Errorf("Bad pulls: %v", tv)
	}

	if s != 1 {
		t.Errorf("Bad singles: %v", s)
	}

}
