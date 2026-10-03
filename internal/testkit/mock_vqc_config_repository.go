package testkit

import (
	"pennylane_project_backend/internal/domain/vqcconfig"

	"uuid"
)

type MockVQCConfigRepository struct {
	vqcConfigs map[uuid.UUID]*vqcconfig.VQCConfig

	ErrSave               error
	ErrFindByID           error
	ErrFindByName         error
	ErrExistsByID         error
	ErrExistsByName       error
	ErrFindAllByOwnerID   error
	ErrDeleteByID         error
	ErrCheckOwnership     error
	ErrDeleteAllByOwnerID error
}

func NewMockVQCConfigRepository() *MockVQCConfigRepository {
	return &MockVQCConfigRepository{
		vqcConfigs: make(map[uuid.UUID]*vqcconfig.VQCConfig),
	}
}

func (r *MockVQCConfigRepository) Save(vqcConfig *vqcconfig.VQCConfig) error {
	if r.ErrSave != nil {
		return r.ErrSave
	}

	r.vqcConfigs[vqcConfig.VQCConfigID()] = vqcConfig
	return nil
}

func (r *MockVQCConfigRepository) FindByID(id uuid.UUID) (*vqcconfig.VQCConfig, error) {
	if r.ErrFindByID != nil {
		return nil, r.ErrFindByID
	}

	if v, ok := r.vqcConfigs[id]; ok {
		copiedVQCConfig := *v
		return &copiedVQCConfig, nil
	}
	return nil, vqcconfig.ErrVQCConfigNotFound
}

func (r *MockVQCConfigRepository) FindByName(ownerID uuid.UUID, name vqcconfig.Name) (*vqcconfig.VQCConfig, error) {
	if r.ErrFindByName != nil {
		return nil, r.ErrFindByName
	}

	for _, v := range r.vqcConfigs {
		if v.OwnerID() == ownerID && v.Name().Equals(name) {
			copiedVQCConfig := *v
			return &copiedVQCConfig, nil
		}
	}
	return nil, vqcconfig.ErrVQCConfigNotFound
}

func (r *MockVQCConfigRepository) ExistsByID(id uuid.UUID) (bool, error) {
	if r.ErrExistsByID != nil {
		return false, r.ErrExistsByID
	}

	_, exists := r.vqcConfigs[id]
	return exists, nil
}

func (r *MockVQCConfigRepository) ExistsByName(ownerID uuid.UUID, name vqcconfig.Name) (bool, error) {
	if r.ErrExistsByName != nil {
		return false, r.ErrExistsByName
	}

	for _, v := range r.vqcConfigs {
		if v.OwnerID() == ownerID && v.Name().Equals(name) {
			return true, nil
		}
	}
	return false, nil
}

func (r *MockVQCConfigRepository) FindAllByOwnerID(ownerID uuid.UUID) ([]*vqcconfig.VQCConfig, error) {
	if r.ErrFindAllByOwnerID != nil {
		return nil, r.ErrFindAllByOwnerID
	}

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
	if r.ErrDeleteByID != nil {
		return r.ErrDeleteByID
	}

	if _, ok := r.vqcConfigs[id]; ok {
		delete(r.vqcConfigs, id)
		return nil
	}
	return vqcconfig.ErrVQCConfigNotFound
}

func (r *MockVQCConfigRepository) CheckOwnership(ownerID uuid.UUID, vqcConfigID uuid.UUID) (bool, error) {
	if r.ErrCheckOwnership != nil {
		return false, r.ErrCheckOwnership
	}

	vqcConfig, exists := r.vqcConfigs[vqcConfigID]
	if !exists {
		return false, vqcconfig.ErrVQCConfigNotFound
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
