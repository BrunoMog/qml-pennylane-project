package testkit

import (
	"pennylane_project_backend/internal/domain/vqc_config"

	"github.com/google/uuid"
)

type MockVQCConfigRepository struct {
	vqcConfigs map[uuid.UUID]*vqc_config.VQCConfig
}

func NewMockVQCConfigRepository() *MockVQCConfigRepository {
	return &MockVQCConfigRepository{
		vqcConfigs: make(map[uuid.UUID]*vqc_config.VQCConfig),
	}
}

func (r *MockVQCConfigRepository) Save(vqcConfig *vqc_config.VQCConfig) error {
	r.vqcConfigs[vqcConfig.GetVQCConfigID()] = vqcConfig
	return nil
}

func (r *MockVQCConfigRepository) FindByID(id uuid.UUID) (*vqc_config.VQCConfig, error) {
	if v, ok := r.vqcConfigs[id]; ok {
		copiedVQCConfig := *v
		return &copiedVQCConfig, nil
	}
	return nil, &ErrVQCConfigNotFound{}
}

func (r *MockVQCConfigRepository) FindByName(name string) (*vqc_config.VQCConfig, error) {
	for _, v := range r.vqcConfigs {
		if v.GetName() == name {
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
		if v.GetOwnerID() == ownerID && v.GetName() == name {
			return true, nil
		}
	}
	return false, nil
}

func (r *MockVQCConfigRepository) FindAllByOwnerID(ownerID uuid.UUID) ([]*vqc_config.VQCConfig, error) {
	var result []*vqc_config.VQCConfig
	for _, v := range r.vqcConfigs {
		if v.GetOwnerID() == ownerID {
			copiedVQCConfig := *v
			result = append(result, &copiedVQCConfig)
		}
	}
	return result, nil
}

func (r *MockVQCConfigRepository) DeleteByID(id uuid.UUID) error {
	delete(r.vqcConfigs, id)
	return nil
}

func (r *MockVQCConfigRepository) DeleteByName(name string) error {
	for id, v := range r.vqcConfigs {
		if v.GetName() == name {
			delete(r.vqcConfigs, id)
			return nil
		}
	}
	return &ErrVQCConfigNotFound{}
}

func (r *MockVQCConfigRepository) DeleteAllByOwnerID(ownerID uuid.UUID) error {
	for id, v := range r.vqcConfigs {
		if v.GetOwnerID() == ownerID {
			delete(r.vqcConfigs, id)
		}
	}
	return nil
}
