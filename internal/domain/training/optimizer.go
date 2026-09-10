package training

import "strings"

type Optimizer interface {
	Name() OptimizerName
	Equals(other Optimizer) bool

	isOptimizer()
}

type OptimizerName string

const (
	OptimizerNameAdam             OptimizerName = "adam"
	OptimizerNameRMSProp          OptimizerName = "rmsprop"
	OptimizerNameNesterovMomentum OptimizerName = "nesterov_momentum"
	OptimizerNameGradientDescent  OptimizerName = "gradient_descent"
)

func (o OptimizerName) Value() string {
	return string(o)
}

func ParseOptimizerName(optName string) (OptimizerName, error) {
	switch strings.ToLower(strings.TrimSpace(optName)) {
	case "adam":
		return OptimizerNameAdam, nil
	case "rmsprop":
		return OptimizerNameRMSProp, nil
	case "nesterov_momentum":
		return OptimizerNameNesterovMomentum, nil
	case "gradient_descent":
		return OptimizerNameGradientDescent, nil
	default:
		return "", &InvalidOptimizerNameError{OptimizerName: optName}
	}
}
