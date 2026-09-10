package training

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLearningType(t *testing.T) {
	testCases := []struct {
		expectedErr error
		testName    string
		input       string
		expected    LearningType
	}{
		{testName: "Valid learning type: supervised", input: "supervised", expected: LearningTypeSupervised, expectedErr: nil},
		{testName: "Invalid learning type", input: "invalid_learning_type", expected: "", expectedErr: &InvalidLearningTypeError{learningType: LearningType("invalid_learning_type")}},
		{testName: "Valid learning type with leading/trailing spaces", input: "  supervised  ", expected: LearningTypeSupervised, expectedErr: nil},
		{testName: "Valid learning type with mixed case", input: "SuPeRvIsEd", expected: LearningTypeSupervised, expectedErr: nil},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result, err := ParseLearningType(tc.input)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.IsType(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestIsValidLearningType(t *testing.T) {
	testCases := []struct {
		testName string
		input    LearningType
		expected bool
	}{
		{"Valid learning type: supervised", LearningTypeSupervised, true},
		{"Invalid learning type", LearningType("invalid_learning_type"), false},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.input.isValid()
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestParseLearningTask(t *testing.T) {
	testCases := []struct {
		expectedErr error
		testName    string
		input       string
		expected    LearningTask
	}{
		{testName: "Valid learning task: binary_classification", input: "binary_classification", expected: LearningTaskBinaryClassification, expectedErr: nil},
		{testName: "Valid learning task: regression", input: "regression", expected: LearningTaskRegression, expectedErr: nil},
		{testName: "Invalid learning task", input: "invalid_learning_task", expected: "", expectedErr: &InvalidLearningTaskError{learningTask: LearningTask("invalid_learning_task")}},
		{testName: "Valid learning task with leading/trailing spaces", input: "  binary_classification  ", expected: LearningTaskBinaryClassification, expectedErr: nil},
		{testName: "Valid learning task with mixed case", input: "BiNaRy_ClAsSiFiCaTiOn", expected: LearningTaskBinaryClassification, expectedErr: nil},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result, err := ParseLearningTask(tc.input)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.IsType(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestIsValidLearningTask(t *testing.T) {
	testCases := []struct {
		testName string
		input    LearningTask
		expected bool
	}{
		{"Valid learning task: binary_classification", LearningTaskBinaryClassification, true},
		{"Valid learning task: regression", LearningTaskRegression, true},
		{"Invalid learning task", LearningTask("invalid_learning_task"), false},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.input.isValid()
			assert.Equal(t, tc.expected, result)
		})
	}
}
