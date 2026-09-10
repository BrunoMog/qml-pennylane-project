package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEmbeddingType(t *testing.T) {
	testCases := []struct {
		expectErr error
		testName  string
		input     string
		expected  EmbeddingType
	}{
		{testName: "Valid angle embedding type", input: "angle", expected: EmbeddingTypeAngle, expectErr: nil},
		{testName: "Valid amplitude embedding type", input: "amplitude", expected: EmbeddingTypeAmplitude, expectErr: nil},
		{testName: "Invalid embedding type", input: "invalid", expected: "", expectErr: &InvalidParseEmbeddingError{"invalid"}},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result, err := ParseEmbeddingType(tc.input)
			if tc.expectErr != nil {
				assert.Error(t, err)
				assert.IsType(t, tc.expectErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}
