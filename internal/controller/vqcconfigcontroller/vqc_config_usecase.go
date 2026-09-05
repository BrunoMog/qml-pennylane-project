package vqcconfigcontroller

import (
	"pennylane_project_backend/internal/usecase/vqcconfigusecase"
)

type VQCConfigUseCase interface {
	CreateVQCConfig(input vqcconfigusecase.CreateVQCConfigInput) (*vqcconfigusecase.CreateVQCConfigOutput, error)
	DeleteVQCConfig(input vqcconfigusecase.DeleteVQCConfigInput) error
	UpdateVQCConfig(input vqcconfigusecase.UpdateVQCConfigInput) error
	LoadVQCConfig(input vqcconfigusecase.LoadVQCConfigInput) (*vqcconfigusecase.LoadVQCConfigOutput, error)
	LoadAllVQCConfigs(input vqcconfigusecase.LoadAllVQCConfigsInput) (*vqcconfigusecase.LoadAllVQCConfigsOutput, error)
}
