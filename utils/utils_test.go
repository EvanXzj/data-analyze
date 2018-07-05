package utils

import "testing"

func TestRandomInt(t *testing.T) {
	res := RandomInt(0, 2)

	if res < 0 && res >= 2 {
		t.Error("res should between [0,2)")
	}
}
