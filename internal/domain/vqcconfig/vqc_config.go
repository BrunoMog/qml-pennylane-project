package vqcconfig

import (
	"pennylane_project_backend/internal/domain/vqc"
	"time"

	"github.com/google/uuid"
)

type VQCConfig struct {
	createdAt   time.Time
	updatedAt   time.Time
	vqc         vqc.VQC
	name        Name
	description Description
	userID      uuid.UUID
	vqcID       uuid.UUID
}

func NewVQCConfig(userID uuid.UUID, name Name, description Description, vqc vqc.VQC) (*VQCConfig, error) {
	if userID == uuid.Nil {
		return nil, &InvalidOwnerIDError{}
	}
	if !vqc.IsValid() {
		return nil, &InvalidVQCError{}
	}

	return &VQCConfig{
		userID:      userID,
		vqcID:       uuid.New(),
		name:        name,
		description: description,
		vqc:         vqc,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}, nil
}

func (vqcConfig *VQCConfig) SetName(name Name) {
	vqcConfig.name = name
	vqcConfig.updatedAt = time.Now()
}

func (vqcConfig *VQCConfig) SetDescription(description Description) {
	vqcConfig.description = description
	vqcConfig.updatedAt = time.Now()
}

func (vqcConfig *VQCConfig) SetVQC(vqc vqc.VQC) {
	vqcConfig.vqc = vqc
	vqcConfig.updatedAt = time.Now()
}

func (vqcConfig *VQCConfig) Name() Name {
	return vqcConfig.name
}

func (vqcConfig *VQCConfig) Description() Description {
	return vqcConfig.description
}

func (vqcConfig *VQCConfig) OwnerID() uuid.UUID {
	return vqcConfig.userID
}

func (vqcConfig *VQCConfig) VQCConfigID() uuid.UUID {
	return vqcConfig.vqcID
}

func (vqcConfig *VQCConfig) CreatedAt() time.Time {
	return vqcConfig.createdAt
}

func (vqcConfig *VQCConfig) UpdatedAt() time.Time {
	return vqcConfig.updatedAt
}

func (vqcConfig *VQCConfig) VQC() vqc.VQC {
	return vqcConfig.vqc
}
