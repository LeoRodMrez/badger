package processor

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSniffFunction(t *testing.T) {
	testSet := []struct {
		Name          string
		GivenInput    Sniffer
		ExpectedError bool
		ErrorMessage  string
	}{
		{
			Name: "(int)",
			GivenInput: Sniffer{
				Function: func(a int) int {
					return a
				},
			},
			ExpectedError: false,
		},
		{
			Name: "(int) with explicit panic",
			GivenInput: Sniffer{
				Function: func(a int) int {
					panic("panicking")
				},
			},
			ExpectedError: true,
			ErrorMessage:  "panic found",
		},
		{
			Name: "(int,int)",
			GivenInput: Sniffer{
				Function: func(a, b int) int {
					return a + b
				},
			},
			ExpectedError: false,
		},
		{
			Name: "(int,int) with explicit panic",
			GivenInput: Sniffer{
				Function: func(a, b int) int {
					panic("panicking")
				},
			},
			ExpectedError: true,
			ErrorMessage:  "panic found",
		},
		{
			Name: "(string)",
			GivenInput: Sniffer{
				Function: func(a string) {
				},
			},
			ExpectedError: false,
		},
		{
			Name: "(string) with explicit panic",
			GivenInput: Sniffer{
				Function: func(a string) {
					panic("panicking")
				},
			},
			ExpectedError: true,
			ErrorMessage:  "panic found",
		},
		{
			Name: "(int,string)",
			GivenInput: Sniffer{
				Function: func(a int,b string) {
				},
			},
			ExpectedError: false,
		},
		{
			Name: "(int,string) with explicit panic",
			GivenInput: Sniffer{
				Function: func(a int,b string) int {
					panic("panicking")
				},
			},
			ExpectedError: true,
			ErrorMessage:  "panic found",
		},
	}

	for _, ts := range testSet {
		t.Run(fmt.Sprintf("TestSniffFunction-%v", ts.Name), func(t *testing.T) {
			response := ts.GivenInput.SniffFunction()
			if ts.ExpectedError {
				require.Error(t, response)
				require.Contains(t, response.Error(), "panic found")
			} else {
				require.NoError(t, response)
			}
		})
	}
}
