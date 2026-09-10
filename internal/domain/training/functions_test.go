package training

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCostFunction(t *testing.T) {
	testCases := []struct {
		expectedErr error
		testName    string
		input       string
		expected    CostFunction
	}{
		{testName: "Valid cost function: binary_cross_entropy", input: "binary_cross_entropy", expected: CostFunctionBinaryCrossEntropy, expectedErr: nil},
		{testName: "Valid cost function: mse", input: "mse", expected: CostFunctionMSE, expectedErr: nil},
		{testName: "Valid cost function: rmse", input: "rmse", expected: CostFunctionRMSE, expectedErr: nil},
		{testName: "Valid cost function: mae", input: "mae", expected: CostFunctionMAE, expectedErr: nil},
		{testName: "Invalid cost function", input: "invalid_cost_function", expected: "", expectedErr: &InvalidCostFunctionError{costFunction: CostFunction("invalid_cost_function")}},
		{testName: "Valid cost function with leading/trailing spaces", input: "  binary_cross_entropy  ", expected: CostFunctionBinaryCrossEntropy, expectedErr: nil},
		{testName: "Valid cost function with mixed case", input: "BiNaRy_CrOsS_EnTrOpY", expected: CostFunctionBinaryCrossEntropy, expectedErr: nil},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result, err := ParseCostFunction(tc.input)
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

func TestIsValidCostFunction(t *testing.T) {
	testCases := []struct {
		testName string
		input    CostFunction
		expected bool
	}{
		{"Valid cost function: binary_cross_entropy", CostFunctionBinaryCrossEntropy, true},
		{"Valid cost function: mse", CostFunctionMSE, true},
		{"Valid cost function: rmse", CostFunctionRMSE, true},
		{"Valid cost function: mae", CostFunctionMAE, true},
		{"Invalid cost function", CostFunction("invalid_cost_function"), false},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.input.isValid()
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestParseEvalMetric(t *testing.T) {
	testCases := []struct {
		expectedErr error
		testName    string
		input       string
		expected    EvalMetric
	}{
		{testName: "Valid eval metric: accuracy", input: "accuracy", expected: EvalMetricAccuracy, expectedErr: nil},
		{testName: "Valid eval metric: f1_score", input: "f1_score", expected: EvalMetricF1Score, expectedErr: nil},
		{testName: "Valid eval metric: precision", input: "precision", expected: EvalMetricPrecision, expectedErr: nil},
		{testName: "Valid eval metric: recall", input: "recall", expected: EvalMetricRecall, expectedErr: nil},
		{testName: "Valid eval metric: rmse", input: "rmse", expected: EvalMetricRMSE, expectedErr: nil},
		{testName: "Valid eval metric: mae", input: "mae", expected: EvalMetricMAE, expectedErr: nil},
		{testName: "Invalid eval metric", input: "invalid_eval_metric", expected: "", expectedErr: &InvalidEvalMetricError{evaluationMetric: EvalMetric("invalid_eval_metric")}},
		{testName: "Valid eval metric with leading/trailing spaces", input: "  accuracy  ", expected: EvalMetricAccuracy, expectedErr: nil},
		{testName: "Valid eval metric with mixed case", input: "AcCuRaCy", expected: EvalMetricAccuracy, expectedErr: nil},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result, err := ParseEvalMetric(tc.input)
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

func TestIsValidEvalMetric(t *testing.T) {
	testCases := []struct {
		testName string
		input    EvalMetric
		expected bool
	}{
		{"Valid eval metric: accuracy", EvalMetricAccuracy, true},
		{"Valid eval metric: f1_score", EvalMetricF1Score, true},
		{"Valid eval metric: precision", EvalMetricPrecision, true},
		{"Valid eval metric: recall", EvalMetricRecall, true},
		{"Valid eval metric: rmse", EvalMetricRMSE, true},
		{"Valid eval metric: mae", EvalMetricMAE, true},
		{"Invalid eval metric", EvalMetric("invalid_eval_metric"), false},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := tc.input.isValid()
			assert.Equal(t, tc.expected, result)
		})
	}
}
