package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/training"
)

type OptimizerDTO struct {
	OptimizerName string
	LearningRate  float64
	Beta1         float64
	Beta2         float64
	Epsilon       float64
	Momentum      float64
	Decay         float64
}

type CrossValidationDTO struct {
	Enabled bool
	Folds   int
}

type EarlyStoppingDTO struct {
	ValidationMetric string
	Patience         int
	MinDelta         float64
	Enabled          bool
}

type TrainingDTO struct {
	LearningTask      string
	LearningType      string
	CostFunction      string
	EvaluationMetrics []string
	EarlyStopping     EarlyStoppingDTO
	Optimizer         OptimizerDTO
	CrossValidation   CrossValidationDTO
	TrainRatio        float64
	ValidationRatio   float64
	TestRatio         float64
	RandomSeed        int
	MaxEpochs         uint
	BatchSize         uint
}

func buildTrainingFromDTO(dto TrainingDTO) (training.Training, error) {
	optimizer, err := buildOptimizer(dto.Optimizer)
	if err != nil {
		return training.Training{}, err
	}

	crossValidation, err := buildCrossValidation(dto.CrossValidation)
	if err != nil {
		return training.Training{}, err
	}

	earlyStopping, err := buildEarlyStopping(dto.EarlyStopping)
	if err != nil {
		return training.Training{}, err
	}

	evalMetrics, err := buildEvalMetrics(dto.EvaluationMetrics)
	if err != nil {
		return training.Training{}, err
	}

	learningTask, err := training.ParseLearningTask(dto.LearningTask)
	if err != nil {
		return training.Training{}, err
	}

	costFunction, err := training.ParseCostFunction(dto.CostFunction)
	if err != nil {
		return training.Training{}, err
	}

	learningType, err := training.ParseLearningType(dto.LearningType)
	if err != nil {
		return training.Training{}, err
	}

	input := training.TrainingInput{
		Optimizer:         optimizer,
		CrossValidation:   crossValidation,
		EarlyStopping:     earlyStopping,
		LearningTask:      learningTask,
		CostFunction:      costFunction,
		LearningType:      learningType,
		EvaluationMetrics: evalMetrics,
		TrainRatio:        dto.TrainRatio,
		ValidationRatio:   dto.ValidationRatio,
		TestRatio:         dto.TestRatio,
		RandomSeed:        dto.RandomSeed,
		MaxEpochs:         dto.MaxEpochs,
		BatchSize:         dto.BatchSize,
	}

	return training.NewTraining(input)
}

func buildOptimizer(dto OptimizerDTO) (training.Optimizer, error) {
	optName, err := training.ParseOptimizerName(dto.OptimizerName)
	if err != nil {
		return nil, err
	}
	switch optName {
	case training.OptimizerNameAdam:
		return training.NewAdamOptimizer(dto.LearningRate, dto.Beta1, dto.Beta2, dto.Epsilon)
	case training.OptimizerNameGradientDescent:
		return training.NewGradientDescentOptimizer(dto.LearningRate)
	case training.OptimizerNameNesterovMomentum:
		return training.NewNesterovMomentumOptimizer(dto.LearningRate, dto.Momentum)
	case training.OptimizerNameRMSProp:
		return training.NewRMSPropOptimizer(dto.LearningRate, dto.Decay, dto.Epsilon)
	default:
		return nil, &UnreachableOptimizerNameError{}
	}
}

func buildCrossValidation(dto CrossValidationDTO) (training.CrossValidation, error) {
	input := training.CrossValidationInput{
		Enabled: dto.Enabled,
		Folds:   dto.Folds,
	}
	return training.NewCrossValidation(input)
}

func buildEarlyStopping(dto EarlyStoppingDTO) (training.EarlyStopping, error) {
	if !dto.Enabled {
		return training.NewEarlyStopping(training.EarlyStoppingInput{Enabled: false})
	}
	evalMetric, err := training.ParseEvalMetric(dto.ValidationMetric)
	if err != nil {
		return training.EarlyStopping{}, err
	}
	input := training.EarlyStoppingInput{
		Enabled:          dto.Enabled,
		ValidationMetric: evalMetric,
		Patience:         dto.Patience,
		MinDelta:         dto.MinDelta,
	}
	return training.NewEarlyStopping(input)
}

func buildEvalMetrics(metrics []string) ([]training.EvalMetric, error) {
	var evalMetrics []training.EvalMetric
	for _, metric := range metrics {
		evalMetric, err := training.ParseEvalMetric(metric)
		if err != nil {
			return nil, err
		}
		evalMetrics = append(evalMetrics, evalMetric)
	}
	return evalMetrics, nil
}

func buildTrainingDTOFromTraining(trainingInstance training.Training) TrainingDTO {
	optimizer := trainingInstance.Optimizer()
	crossValidation := trainingInstance.CrossValidation()
	earlyStopping := trainingInstance.EarlyStopping()
	learningTask := trainingInstance.LearningTask()
	costFunction := trainingInstance.CostFunction()
	learningType := trainingInstance.LearningType()
	evalMetrics := trainingInstance.EvaluationMetrics()

	return TrainingDTO{
		Optimizer:         buildOptimizerDTOFromOptimizer(optimizer),
		CrossValidation:   buildCrossValidationDTOFromCrossValidation(crossValidation),
		EarlyStopping:     buildEarlyStoppingDTOFromEarlyStopping(earlyStopping),
		LearningTask:      learningTask.Value(),
		CostFunction:      costFunction.Value(),
		LearningType:      learningType.Value(),
		EvaluationMetrics: buildEvalMetricsDTOFromEvalMetrics(evalMetrics),
		TrainRatio:        trainingInstance.TrainRatio(),
		ValidationRatio:   trainingInstance.ValidationRatio(),
		TestRatio:         trainingInstance.TestRatio(),
		RandomSeed:        trainingInstance.RandomSeed(),
		MaxEpochs:         trainingInstance.MaxEpochs(),
		BatchSize:         trainingInstance.BatchSize(),
	}

}

func buildOptimizerDTOFromOptimizer(optimizer training.Optimizer) OptimizerDTO {
	switch opt := optimizer.(type) {
	case training.AdamOptimizer:
		return OptimizerDTO{
			OptimizerName: training.OptimizerNameAdam.Value(),
			LearningRate:  opt.LearningRate(),
			Beta1:         opt.Beta1(),
			Beta2:         opt.Beta2(),
			Epsilon:       opt.Epsilon(),
		}
	case training.GradientDescentOptimizer:
		return OptimizerDTO{
			OptimizerName: training.OptimizerNameGradientDescent.Value(),
			LearningRate:  opt.LearningRate(),
		}
	case training.NesterovMomentumOptimizer:
		return OptimizerDTO{
			OptimizerName: training.OptimizerNameNesterovMomentum.Value(),
			LearningRate:  opt.LearningRate(),
			Momentum:      opt.Momentum(),
		}
	case training.RMSPropOptimizer:
		return OptimizerDTO{
			OptimizerName: training.OptimizerNameRMSProp.Value(),
			LearningRate:  opt.LearningRate(),
			Decay:         opt.Decay(),
			Epsilon:       opt.Epsilon(),
		}
	default:
		return OptimizerDTO{}
	}
}

func buildCrossValidationDTOFromCrossValidation(crossValidation training.CrossValidation) CrossValidationDTO {
	return CrossValidationDTO{
		Enabled: crossValidation.Enabled(),
		Folds:   crossValidation.Folds(),
	}
}

func buildEarlyStoppingDTOFromEarlyStopping(earlyStopping training.EarlyStopping) EarlyStoppingDTO {
	return EarlyStoppingDTO{
		Enabled:          earlyStopping.Enabled(),
		ValidationMetric: earlyStopping.ValidationMetric().Value(),
		Patience:         earlyStopping.Patience(),
		MinDelta:         earlyStopping.MinDelta(),
	}
}

func buildEvalMetricsDTOFromEvalMetrics(evalMetrics []training.EvalMetric) []string {
	var metrics []string
	for _, metric := range evalMetrics {
		metrics = append(metrics, metric.Value())
	}
	return metrics
}
