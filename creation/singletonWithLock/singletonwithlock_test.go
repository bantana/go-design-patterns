package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	ins1 := New()
	ins2 := New()
	ins1.AddOne()
	ins2.AddOne()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}
