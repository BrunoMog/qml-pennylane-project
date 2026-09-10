package training

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOptimizerName(t *testing.T) {
	testCases := []struct {
		expectedError error
		testName      string
		input         string
		expectedName  OptimizerName
	}{
		{testName: "Valid Adam", input: "adam", expectedName: OptimizerNameAdam, expectedError: nil},
		{testName: "Valid RMSProp", input: "rmsprop", expectedName: OptimizerNameRMSProp, expectedError: nil},
		{testName: "Valid Nesterov Momentum", input: "nesterov_momentum", expectedName: OptimizerNameNesterovMomentum, expectedError: nil},
		{testName: "Valid Gradient Descent", input: "gradient_descent", expectedName: OptimizerNameGradientDescent, expectedError: nil},
		{testName: "Invalid Optimizer Name", input: "invalid_optimizer", expectedName: "", expectedError: &InvalidOptimizerNameError{OptimizerName: "invalid_optimizer"}},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			name, err := ParseOptimizerName(tc.input)
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedName, name)
			}
		})
	}
}
