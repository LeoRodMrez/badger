package helper

import (
	"math/rand"
	"time"

	"github.com/LeoRodMrez/badger/constants"
)

// GetRandInt returns a random integer given a min and a max range
func GetRandInt(min, max int) int {
	if willReturnZeroValue() {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// GetRandomIntPtr returns a random integer pointer given a min and a max range
func GetRandomIntPtr(min, max int) *int {
	if willReturnZeroValue() {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(max-min+1) + min
	return &n
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

// GetRandString return a random string
func GetRandString() string {
	if willReturnZeroValue() {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	randLength := rand.Intn(30-1+1) + 1
	b := make([]byte, randLength)
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = constants.Charset[seededRand.Intn(len(constants.Charset))]
	}
	return string(b)
}

// GetRandBoolPtr returns a random bool pointer
func GetRandBoolPtr() *bool {
	if willReturnZeroValue() {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 1 {
		a := true
		return &a
	} else {
		a := false
		return &a
	}
}

// GetRandStringPtr return a random string pointer
func GetRandStringPtr() *string {
	if willReturnZeroValue() {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	randLength := rand.Intn(30-1+1) + 1
	b := make([]byte, randLength)
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = constants.Charset[seededRand.Intn(len(constants.Charset))]
	}
	a := string(b)
	return &a
}

// willReturnZeroValue decides whether the rand function will return the zero value or generate a random one
func willReturnZeroValue() bool {
	if rand.Intn(3) == 1 {
		return true
	} else {
		return false
	}
}
