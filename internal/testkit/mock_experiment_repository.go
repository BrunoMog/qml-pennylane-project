package testkit

import (
	"pennylane_project_backend/internal/domain/experiment"

	"uuid"
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

func (m *MockExperimentRepository) FindByName(ownerID uuid.UUID, name experiment.Name) (*experiment.Experiment, error) {
	for _, v := range m.experiments {
		if v.OwnerID() == ownerID && v.Name().Equals(name) {
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

func (m *MockExperimentRepository) ExistsByName(ownerID uuid.UUID, name experiment.Name) (bool, error) {
	for _, v := range m.experiments {
		if v.OwnerID() == ownerID && v.Name().Equals(name) {
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

func (m *MockExperimentRepository) CheckOwnership(ownerID, experimentID uuid.UUID) (bool, error) {
	if experiment, ok := m.experiments[experimentID]; ok {
		return experiment.OwnerID() == ownerID, nil
	}
	return false, &ErrExperimentNotFound{}
}

func (m *MockExperimentRepository) DeleteAllByOwnerID(ownerID uuid.UUID) error {
	for id, v := range m.experiments {
		if v.OwnerID() == ownerID {
			delete(m.experiments, id)
		}
	}
	return nil
}
