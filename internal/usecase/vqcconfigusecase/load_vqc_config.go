package vqcconfigusecase

import (
	"errors"
	"fmt"
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"pennylane_project_backend/internal/usecase/apperrors"
	"time"

	"uuid"
)

type LoadVQCConfigByIDInput struct {
	VQCConfigID uuid.UUID
	CallerID    uuid.UUID
}

type LoadVQCConfigByNameInput struct {
	VQCConfigName string
	CallerID      uuid.UUID
}

type LoadVQCConfigOutput struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	VQCDTO      VQCDTO
	OwnerID     uuid.UUID
	VQCConfigID uuid.UUID
}

func (s *VQCConfigService) LoadVQCConfigByID(input LoadVQCConfigByIDInput) (*LoadVQCConfigOutput, error) {
	if input.CallerID == uuid.Nil() {
		return nil, user.ErrNilUserID
	}

	if input.VQCConfigID == uuid.Nil() {
		return nil, vqcconfig.ErrNilVQCConfigID
	}

	config, err := s.vqcConfigRepository.FindByID(input.VQCConfigID)
	if err != nil {
		if errors.Is(err, vqcconfig.ErrVQCConfigNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("vqcconfigusecase: find vqc config by ID: %w", err)
	}

	if !canLoadVQCConfig(input.CallerID, config) {
		return nil, apperrors.NewPermissionDeniedError(
			"user does not own the VQC config",
			"load VQC config",
			input.CallerID,
		)
	}

	output := toLoadVQCConfigOutput(config)
	return &output, nil
}

func canLoadVQCConfig(callerID uuid.UUID, config *vqcconfig.VQCConfig) bool {
	return callerID == config.OwnerID()
}

func toLoadVQCConfigOutput(config *vqcconfig.VQCConfig) LoadVQCConfigOutput {
	return LoadVQCConfigOutput{
		Name:        config.Name().String(),
		Description: config.Description().String(),
		VQCDTO:      buildVQCDTOFromVQC(config.VQC()),
		CreatedAt:   config.CreatedAt(),
		UpdatedAt:   config.UpdatedAt(),
		OwnerID:     config.OwnerID(),
		VQCConfigID: config.VQCConfigID(),
	}
}

func (s *VQCConfigService) LoadVQCConfigByName(input LoadVQCConfigByNameInput) (*LoadVQCConfigOutput, error) {
	if input.CallerID == uuid.Nil() {
		return nil, user.ErrNilUserID
	}

	name, err := vqcconfig.NewName(input.VQCConfigName)
	if err != nil {
		return nil, err
	}

	config, err := s.vqcConfigRepository.FindByName(input.CallerID, name)
	if err != nil {
		if errors.Is(err, vqcconfig.ErrVQCConfigNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("vqcconfigusecase: find vqc config by name: %w", err)
	}

	output := toLoadVQCConfigOutput(config)
	return &output, nil
}

type LoadAllVQCConfigsInput struct {
	CallerID uuid.UUID
}

type LoadAllVQCConfigsOutput struct {
	VQCConfigs []LoadVQCConfigOutput
}

func (s *VQCConfigService) LoadAllVQCConfigs(input LoadAllVQCConfigsInput) (*LoadAllVQCConfigsOutput, error) {
	configs, err := s.vqcConfigRepository.FindAllByOwnerID(input.CallerID)
	if err != nil {
		return nil, fmt.Errorf("vqcconfigusecase: find all vqc configs by owner ID: %w", err)
	}

	output := &LoadAllVQCConfigsOutput{
		VQCConfigs: make([]LoadVQCConfigOutput, len(configs)),
	}
	for i, config := range configs {
		output.VQCConfigs[i] = toLoadVQCConfigOutput(config)
	}
	return output, nil
}
