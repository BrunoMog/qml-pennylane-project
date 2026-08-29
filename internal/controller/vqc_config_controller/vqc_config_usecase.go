package vqc_config_usecase

import (
	"pennylane_project_backend/internal/usecase/vqc_config_usecase"
)

type VQCConfigUseCase interface {
	CreateVQCConfig(input vqc_config_usecase.CreateVQCConfigInput) (*vqc_config_usecase.CreateVQCConfigOutput, error)
	DeleteVQCConfig(input vqc_config_usecase.DeleteVQCConfigInput) error
	UpdateVQCConfig(input vqc_config_usecase.UpdateVQCConfigInput) error
	GetVQCConfig(input vqc_config_usecase.GetVQCConfigInput) (*vqc_config_usecase.GetVQCConfigOutput, error)
	ListVQCConfigs(input vqc_config_usecase.ListVQCConfigsInput) ([]*vqc_config_usecase.GetVQCConfigOutput, error)
}
