package processor

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSniffFunction(t *testing.T) {
	testSet := []struct {
		Name          string
		GivenInput    SniffableFunc
		ExpectedError bool
		ErrorMessage  string
	}{
		{
			Name: "(int,int)",
			GivenInput: SniffableFunc{
				Function: func(a, b int) int {
					return a + b
				},
			},
			ExpectedError: false,
		},
		{
			Name: "(int,int) with explicit panic",
			GivenInput: SniffableFunc{
				Function: func(a, b int) int {
					panic("panicking")
				},
			},
			ExpectedError: true,
			ErrorMessage: "panic found",
		},
	}

	for _, ts := range testSet {
		t.Run(fmt.Sprintf("TestSniffFunction-%v", ts.Name), func(t *testing.T) {
			response := ts.GivenInput.SniffFunction()
			if ts.ExpectedError {
				require.Error(t,response)
				require.Contains(t, response.Error(), "panic found")
			} else {
				require.NoError(t,response)
			}
		})
	}
}
