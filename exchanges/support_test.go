package exchange

import "testing"

func TestIsSupported(t *testing.T) {
	if ok := IsSupported("BiTStaMp"); !ok {
		t.Error("supported exchange should be valid")
