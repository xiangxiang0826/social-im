package randUtils

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Logf("%d", RandInt(0, 10))
	}
}


func TestRandStr(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	a :=RandString(3)
	rand.Seed(time.Now().UnixNano())
		t.Logf("%s", a+""+RandString(6))

}
