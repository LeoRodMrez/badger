package helper

import (
	"badger/constants"
	"math/rand"
	"time"
)

// GetRandInt returns a random integer given a min and a max range
func GetRandInt(min, max int) int {
	if willReturnZeroValue() {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// GetRandBool returns a random bool
func GetRandBool() bool {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 1 {
		return true
	} else {
		return false
	}
}

func GetRandString(length int) string {
	if willReturnZeroValue() {
		return ""
	}
	b := make([]byte, length)
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = constants.Charset[seededRand.Intn(len(constants.Charset))]
	}
	return string(b)
}

// willReturnZeroValue decides whether the rand function will return the zero value or generate a random one
func willReturnZeroValue() bool {
	if rand.Intn(3) == 1 {
		return true
	} else {
		return false
	}
}