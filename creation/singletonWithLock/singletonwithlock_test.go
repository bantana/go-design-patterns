package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	ins1.AddOne()
	ins2.AddOne()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}
