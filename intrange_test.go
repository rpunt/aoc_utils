package utils

import (
	"testing"
)

// Contains
func TestContains(t *testing.T){
	r1 := Intrange{1,5}
	r2 := Intrange{3,7}
	r3 := Intrange{2,4}
	if got := r1.Contains(r2); got != false {
		t.Errorf("Contains: r1 does not contain r2")
	}
	if got := r1.Contains(r3); got != true {
		t.Errorf("Contains: r1 does contain r3")
	}
}

// Equals
func TestEquals(t *testing.T) {
	r1 := Intrange{1,5}
	r2 := Intrange{1,5}
	r3 := Intrange{1,3}
	r4 := Intrange{3,7}
	if got := r1.Equals(r2); got != true {
		t.Errorf("Equals: r1 does equal r2")
	}
	if got := r1.Equals(r3); got != false {
		t.Errorf("Equals: r1 does not equal r3")
	}
	if got := r1.Equals(r4); got != false {
		t.Errorf("Equals: r1 does not equal r4")
	}
}

// Overlaps
func TestOverlaps(t *testing.T) {
	r1 := Intrange{1,5}
	r2 := Intrange{0,3}
	r3 := Intrange{3,7}
	r4 := Intrange{6,9}
	r5 := Intrange{-5,0}
	r6 := Intrange{83,91}
	r7 := Intrange{82,97}
	r8 := Intrange{82,97}
	if got := r1.Overlaps(r2); got != true {
		t.Errorf("Overlaps: r1 does overlap r2")
	}
	if got := r1.Overlaps(r3); got != true {
		t.Errorf("Overlaps: r1 does not overlap r3")
	}
	if got := r1.Overlaps(r4); got != false {
		t.Errorf("Overlaps: r1 does not overlap r4")
	}
	if got := r1.Overlaps(r5); got != false {
		t.Errorf("Overlaps: r1 does not overlap r5")
	}
	if got := r6.Overlaps(r7); got != true {
		t.Errorf("Overlaps: r6 does overlap r7")
	}
	// equality is technically also overlap
	if got := r7.Overlaps(r8); got != true {
		t.Errorf("Overlaps: r7 does overlap r8")
	}
}
