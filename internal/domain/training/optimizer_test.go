package training

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOptimizerName(t *testing.T) {
	testCases := []struct {
		testName      string
		input         string
		expectedName  OptimizerName
		expectedError error
	}{
		{"Valid Adam", "adam", OptimizerNameAdam, nil},
		{"Valid RMSProp", "rmsprop", OptimizerNameRMSProp, nil},
		{"Valid Nesterov Momentum", "nesterov_momentum", OptimizerNameNesterovMomentum, nil},
		{"Valid Gradient Descent", "gradient_descent", OptimizerNameGradientDescent, nil},
		{"Invalid Optimizer Name", "invalid_optimizer", "", &InvalidOptimizerNameError{OptimizerName: "invalid_optimizer"}},
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
