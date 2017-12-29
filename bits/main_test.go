package main

import "testing"

func TestCountOnes(t *testing.T) {
	if countOnes(0xF) != 4 {
		t.Fail()
	}
	if countOnes(0xFF) != 8 {
		t.Fail()
	}
	if countOnes(0x0) != 0 {
		t.Fail()
	}
	if countOnes(0xA0A) != 4 {
		t.Fail()
	}
}
