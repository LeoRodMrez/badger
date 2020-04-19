package constants

const (
	// Charset is used for random string generation
	Charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" +
		"0123456789" +
		"~=+%^*/()[]{}/!@#$?|" +
		"                    "

	// MaxUint is the maximum value for a uint variable
	MaxUint = ^uint(0)

	// MaxInt is the maximum value for an int variable
	MaxInt = int(MaxUint >> 1)

	// MinInt is the minimum value for an int variable
	MinInt = -MaxInt - 1
)
