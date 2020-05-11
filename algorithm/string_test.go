package algorithm

import (
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestAlgorithms(t *testing.T) {
	t.Run("numberToWords", func(t *testing.T) {
		testCases := []struct {
			number   int
			expected string
		}{
			{
				12,
				"Twelve",
			},
		}

		for _, tc := range testCases {
			words := numberToWords(tc.number)
			assert.Equal(t, words, tc.expected)
		}
	})
}
