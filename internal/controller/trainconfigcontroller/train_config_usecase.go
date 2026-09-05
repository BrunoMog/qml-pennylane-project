package trainconfigcontroller

import (
	"pennylane_project_backend/internal/usecase/trainconfigusecase"
)

type TrainConfigUseCase interface {
	CreateTrainConfig(input trainconfigusecase.CreateTrainConfigInput) (*trainconfigusecase.CreateTrainConfigOutput, error)
	DeleteTrainConfig(input trainconfigusecase.DeleteTrainConfigInput) error
	UpdateTrainConfig(input trainconfigusecase.UpdateTrainConfigInput) error
	LoadTrainConfig(input trainconfigusecase.LoadTrainConfigInput) (*trainconfigusecase.LoadTrainConfigOutput, error)
	LoadAllTrainConfigs(input trainconfigusecase.LoadAllTrainConfigsInput) (*trainconfigusecase.LoadAllTrainConfigsOutput, error)
}
