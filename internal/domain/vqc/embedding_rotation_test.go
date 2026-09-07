package vqc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		testName string
		rotation EmbeddingRotation
		expected bool
	}{
		{testName: "Valid X rotation", rotation: XRotation, expected: true},
		{testName: "Valid Y rotation", rotation: YRotation, expected: true},
		{testName: "Valid Z rotation", rotation: ZRotation, expected: true},
		{testName: "Invalid rotation", rotation: EmbeddingRotation("invalid"), expected: false},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.rotation.IsValid()
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestParseEmbeddingRotation(t *testing.T) {
	testCases := []struct {
		testName  string
		input     string
		expected  EmbeddingRotation
		expectErr error
	}{
		{testName: "Valid X rotation", input: "x", expected: XRotation, expectErr: nil},
		{testName: "Valid Y rotation", input: "y", expected: YRotation, expectErr: nil},
		{testName: "Valid Z rotation", input: "z", expected: ZRotation, expectErr: nil},
		{testName: "Invalid rotation", input: "invalid", expected: "", expectErr: &InvalidParseEmbeddingRotationError{"invalid"}},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result, err := ParseEmbeddingRotation(tc.input)
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
