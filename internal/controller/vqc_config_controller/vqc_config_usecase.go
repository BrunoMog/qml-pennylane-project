package vqcconfigcontroller

import (
	"pennylane_project_backend/internal/usecase/vqcconfigusecase"
)

type VQCConfigUseCase interface {
	CreateVQCConfig(input vqcconfigusecase.CreateVQCConfigInput) (*vqcconfigusecase.CreateVQCConfigOutput, error)
	DeleteVQCConfig(input vqcconfigusecase.DeleteVQCConfigInput) error
	UpdateVQCConfig(input vqcconfigusecase.UpdateVQCConfigInput) error
	GetVQCConfig(input vqcconfigusecase.LoadVQCConfigInput) (*vqcconfigusecase.GetVQCConfigOutput, error)
	ListVQCConfigs(input vqcconfigusecase.ListVQCConfigsInput) ([]*vqcconfigusecase.GetVQCConfigOutput, error)
}
