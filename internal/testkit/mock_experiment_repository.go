package testkit

import (
	"pennylane_project_backend/internal/domain/experiment"

	"github.com/google/uuid"
)

type MockExperimentRepository struct {
	experiments map[uuid.UUID]*experiment.Experiment
}

func NewMockExperimentRepository() *MockExperimentRepository {
	return &MockExperimentRepository{
		experiments: make(map[uuid.UUID]*experiment.Experiment),
	}
}

func (m *MockExperimentRepository) Save(experiment *experiment.Experiment) error {
	m.experiments[experiment.ExperimentID()] = experiment
	return nil
}

func (m *MockExperimentRepository) FindByID(experimentID uuid.UUID) (*experiment.Experiment, error) {
	if v, ok := m.experiments[experimentID]; ok {
		copiedExperiment := *v
		return &copiedExperiment, nil
	}
	return nil, &ErrExperimentNotFound{}
}

func (m *MockExperimentRepository) FindByName(ownerID uuid.UUID, name string) (*experiment.Experiment, error) {
	for _, v := range m.experiments {
		if v.OwnerID() == ownerID && v.Name() == name {
			copiedExperiment := *v
			return &copiedExperiment, nil
		}
	}
	return nil, &ErrExperimentNotFound{}
}

func (m *MockExperimentRepository) ExistsByID(experimentID uuid.UUID) (bool, error) {
	_, exists := m.experiments[experimentID]
	return exists, nil
}

func (m *MockExperimentRepository) ExistsByName(ownerID uuid.UUID, name string) (bool, error) {
	for _, v := range m.experiments {
		if v.OwnerID() == ownerID && v.Name() == name {
			return true, nil
		}
	}
	return false, nil
}

func (m *MockExperimentRepository) FindAllByOwnerID(ownerID uuid.UUID) ([]*experiment.Experiment, error) {
	var result []*experiment.Experiment
	for _, v := range m.experiments {
		if v.OwnerID() == ownerID {
			copiedExperiment := *v
			result = append(result, &copiedExperiment)
		}
	}
	return result, nil
}

func (m *MockExperimentRepository) DeleteByID(experimentID uuid.UUID) error {
	if _, ok := m.experiments[experimentID]; ok {
		delete(m.experiments, experimentID)
		return nil
	}
	return &ErrExperimentNotFound{}
}

func (m *MockExperimentRepository) DeleteAllByOwnerID(ownerID uuid.UUID) error {
	for id, v := range m.experiments {
		if v.OwnerID() == ownerID {
			delete(m.experiments, id)
		}
	}
	return nil
}
