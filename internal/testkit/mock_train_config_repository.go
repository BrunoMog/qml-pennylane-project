package testkit

import (
	"pennylane_project_backend/internal/domain/trainconfig"

	"github.com/google/uuid"
)

type MockTrainConfigRepository struct {
	trainConfigs map[uuid.UUID]*trainconfig.TrainConfig
}

func NewMockTrainConfigRepository() *MockTrainConfigRepository {
	return &MockTrainConfigRepository{
		trainConfigs: make(map[uuid.UUID]*trainconfig.TrainConfig),
	}
}

func (m *MockTrainConfigRepository) Save(trainConfig *trainconfig.TrainConfig) error {
	m.trainConfigs[trainConfig.TrainConfigID()] = trainConfig
	return nil
}

func (m *MockTrainConfigRepository) FindByID(trainConfigID uuid.UUID) (*trainconfig.TrainConfig, error) {
	if v, ok := m.trainConfigs[trainConfigID]; ok {
		copiedTrainConfig := *v
		return &copiedTrainConfig, nil
	}
	return nil, &ErrTrainConfigNotFound{}
}

func (m *MockTrainConfigRepository) FindByName(ownerID uuid.UUID, name string) (*trainconfig.TrainConfig, error) {
	for _, v := range m.trainConfigs {
		if v.OwnerID() == ownerID && v.Name() == name {
			copiedTrainConfig := *v
			return &copiedTrainConfig, nil
		}
	}
	return nil, &ErrTrainConfigNotFound{}
}

func (m *MockTrainConfigRepository) ExistsByID(trainConfigID uuid.UUID) (bool, error) {
	_, exists := m.trainConfigs[trainConfigID]
	return exists, nil
}

func (m *MockTrainConfigRepository) ExistsByName(ownerID uuid.UUID, name string) (bool, error) {
	for _, v := range m.trainConfigs {
		if v.OwnerID() == ownerID && v.Name() == name {
			return true, nil
		}
	}
	return false, nil
}

func (m *MockTrainConfigRepository) FindAllByOwnerID(ownerID uuid.UUID) ([]*trainconfig.TrainConfig, error) {
	var result []*trainconfig.TrainConfig
	for _, v := range m.trainConfigs {
		if v.OwnerID() == ownerID {
			copiedTrainConfig := *v
			result = append(result, &copiedTrainConfig)
		}
	}
	return result, nil
}

func (m *MockTrainConfigRepository) DeleteByID(trainConfigID uuid.UUID) error {
	if _, ok := m.trainConfigs[trainConfigID]; ok {
		delete(m.trainConfigs, trainConfigID)
		return nil
	}
	return &ErrTrainConfigNotFound{}
}

func (m *MockTrainConfigRepository) CheckOwnership(onwerID uuid.UUID, trainConfigID uuid.UUID) (bool, error) {
	trainConfig, exists := m.trainConfigs[trainConfigID]
	if !exists {
		return false, &ErrTrainConfigNotFound{}
	}
	return trainConfig.OwnerID() == onwerID, nil
}

func (m *MockTrainConfigRepository) DeleteAllByOwnerID(ownerID uuid.UUID) error {
	for id, v := range m.trainConfigs {
		if v.OwnerID() == ownerID {
			delete(m.trainConfigs, id)
		}
	}
	return nil
}
