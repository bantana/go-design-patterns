package singleton

import "testing"

func TestSingleton(t *testing.T) {
	ins1 := New()
	ins2 := New()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}
