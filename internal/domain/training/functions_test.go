package training

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCostFunction(t *testing.T) {
	testCases := []struct {
		testName    string
		input       string
		expected    CostFunction
		expectedErr error
	}{
		{"Valid cost function: binary_cross_entropy", "binary_cross_entropy", CostFunctionBinaryCrossEntropy, nil},
		{"Valid cost function: mse", "mse", CostFunctionMSE, nil},
		{"Valid cost function: rmse", "rmse", CostFunctionRMSE, nil},
		{"Valid cost function: mae", "mae", CostFunctionMAE, nil},
		{"Invalid cost function", "invalid_cost_function", "", &InvalidCostFunctionError{costFunction: CostFunction("invalid_cost_function")}},
		{"Valid cost function with leading/trailing spaces", "  binary_cross_entropy  ", CostFunctionBinaryCrossEntropy, nil},
		{"Valid cost function with mixed case", "BiNaRy_CrOsS_EnTrOpY", CostFunctionBinaryCrossEntropy, nil},
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
		testName    string
		input       string
		expected    EvalMetric
		expectedErr error
	}{
		{"Valid eval metric: accuracy", "accuracy", EvalMetricAccuracy, nil},
		{"Valid eval metric: f1_score", "f1_score", EvalMetricF1Score, nil},
		{"Valid eval metric: precision", "precision", EvalMetricPrecision, nil},
		{"Valid eval metric: recall", "recall", EvalMetricRecall, nil},
		{"Valid eval metric: rmse", "rmse", EvalMetricRMSE, nil},
		{"Valid eval metric: mae", "mae", EvalMetricMAE, nil},
		{"Invalid eval metric", "invalid_eval_metric", "", &InvalidEvalMetricError{evaluationMetric: EvalMetric("invalid_eval_metric")}},
		{"Valid eval metric with leading/trailing spaces", "  accuracy  ", EvalMetricAccuracy, nil},
		{"Valid eval metric with mixed case", "AcCuRaCy", EvalMetricAccuracy, nil},
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
