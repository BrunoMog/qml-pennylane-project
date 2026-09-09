package training

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLearningType(t *testing.T) {
	testCases := []struct {
		testName    string
		input       string
		expected    LearningType
		expectedErr error
	}{
		{"Valid learning type: supervised", "supervised", LearningTypeSupervised, nil},
		{"Invalid learning type", "invalid_learning_type", "", &InvalidLearningTypeError{learningType: LearningType("invalid_learning_type")}},
		{"Valid learning type with leading/trailing spaces", "  supervised  ", LearningTypeSupervised, nil},
		{"Valid learning type with mixed case", "SuPeRvIsEd", LearningTypeSupervised, nil},
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
		testName    string
		input       string
		expected    LearningTask
		expectedErr error
	}{
		{"Valid learning task: classification", "classification", LearningTaskBinaryClassification, nil},
		{"Valid learning task: regression", "regression", LearningTaskRegression, nil},
		{"Invalid learning task", "invalid_learning_task", "", &InvalidLearningTaskError{learningTask: LearningTask("invalid_learning_task")}},
		{"Valid learning task with leading/trailing spaces", "  classification  ", LearningTaskBinaryClassification, nil},
		{"Valid learning task with mixed case", "ClAsSiFiCaTiOn", LearningTaskBinaryClassification, nil},
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
		{"Valid learning task: classification", LearningTaskBinaryClassification, true},
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
