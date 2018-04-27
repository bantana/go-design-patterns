package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	ins1 := Get()
	ins2 := Get()
	ins1.Add()
	ins2.Add()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}
