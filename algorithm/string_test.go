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

func TestIP(t *testing.T) {
	t.Run("validIPv4Segment", func(t *testing.T) {
		testCases := []struct {
			segment  string
			expected bool
		}{
			{
				"0",
				true,
			},
			{
				"1",
				true,
			},
			{
				"12",
				true,
			},
			{
				"123",
				true,
			},
			{
				"255",
				true,
			},
			{
				"250",
				true,
			},
			{
				"256",
				false,
			},
			{
				"1234",
				false,
			},
			{
				"012",
				false,
			},
			{
				"00",
				false,
			},
		}

		for _, tc := range testCases {
			result := validIPv4Segment(tc.segment)
			assert.Equal(t, result, tc.expected)
		}
	})
}
