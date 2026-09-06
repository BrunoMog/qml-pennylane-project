package testkit

import (
	"pennylane_project_backend/internal/domain/vqcconfig"

	"github.com/google/uuid"
)

type MockVQCConfigRepository struct {
	vqcConfigs map[uuid.UUID]*vqcconfig.VQCConfig
}

func NewMockVQCConfigRepository() *MockVQCConfigRepository {
	return &MockVQCConfigRepository{
		vqcConfigs: make(map[uuid.UUID]*vqcconfig.VQCConfig),
	}
}

func (r *MockVQCConfigRepository) Save(vqcConfig *vqcconfig.VQCConfig) error {
	r.vqcConfigs[vqcConfig.VQCConfigID()] = vqcConfig
	return nil
}

func (r *MockVQCConfigRepository) FindByID(id uuid.UUID) (*vqcconfig.VQCConfig, error) {
	if v, ok := r.vqcConfigs[id]; ok {
		copiedVQCConfig := *v
		return &copiedVQCConfig, nil
	}
	return nil, &ErrVQCConfigNotFound{}
}

func (r *MockVQCConfigRepository) FindByName(ownerID uuid.UUID, name string) (*vqcconfig.VQCConfig, error) {
	for _, v := range r.vqcConfigs {
		if v.OwnerID() == ownerID && v.Name() == name {
			copiedVQCConfig := *v
			return &copiedVQCConfig, nil
		}
	}
	return nil, &ErrVQCConfigNotFound{}
}

func (r *MockVQCConfigRepository) ExistsByID(id uuid.UUID) (bool, error) {
	_, exists := r.vqcConfigs[id]
	return exists, nil
}

func (r *MockVQCConfigRepository) ExistsByName(ownerID uuid.UUID, name string) (bool, error) {
	for _, v := range r.vqcConfigs {
		if v.OwnerID() == ownerID && v.Name() == name {
			return true, nil
		}
	}
	return false, nil
}

func (r *MockVQCConfigRepository) FindAllByOwnerID(ownerID uuid.UUID) ([]*vqcconfig.VQCConfig, error) {
	var result []*vqcconfig.VQCConfig
	for _, v := range r.vqcConfigs {
		if v.OwnerID() == ownerID {
			copiedVQCConfig := *v
			result = append(result, &copiedVQCConfig)
		}
	}
	return result, nil
}

func (r *MockVQCConfigRepository) DeleteByID(id uuid.UUID) error {
	if _, ok := r.vqcConfigs[id]; ok {
		delete(r.vqcConfigs, id)
		return nil
	}
	return &ErrVQCConfigNotFound{}
}

func (r *MockVQCConfigRepository) CheckOwnership(ownerID uuid.UUID, vqcConfigID uuid.UUID) (bool, error) {
	vqcConfig, exists := r.vqcConfigs[vqcConfigID]
	if !exists {
		return false, &ErrVQCConfigNotFound{}
	}
	return vqcConfig.OwnerID() == ownerID, nil
}

func (r *MockVQCConfigRepository) DeleteAllByOwnerID(ownerID uuid.UUID) error {
	for id, v := range r.vqcConfigs {
		if v.OwnerID() == ownerID {
			delete(r.vqcConfigs, id)
		}
	}
	return nil
}
