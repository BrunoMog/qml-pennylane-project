package trainingpipeline

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
