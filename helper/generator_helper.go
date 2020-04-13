package helper

import (
	"math/rand"
	"time"
)

// GetRandInt returns a random integer given a min and a max range
func GetRandInt(min, max int)int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min + 1) + min
}
