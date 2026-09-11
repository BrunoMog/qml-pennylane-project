package experimentusecase

import (
	"pennylane_project_backend/internal/domain/experiment"
	"uuid"
)

type UpdateExperimentInput struct {
	CallerID      uuid.UUID
	ExperimentID  uuid.UUID
	Name          *string
	Description   *string
	TrainConfigID *uuid.UUID
	VQCConfigID   *uuid.UUID
}

func (u *ExperimentService) UpdateExperiment(input UpdateExperimentInput) error {
	if input.Name == nil && input.Description == nil && input.TrainConfigID == nil && input.VQCConfigID == nil {
		return &NoFieldsToUpdateError{}
	}

	exp, err := u.experimentRepository.FindByID(input.ExperimentID)
	if err != nil {
		return err
	}

	if !canUpdateExperiment(input.CallerID, exp) {
		return &UnauthorizedError{}
	}

	var needToSave bool

	if input.Name != nil {
		name, err := experiment.NewName(*input.Name)
		if err != nil {
			return err
		}

		if exp.Name().Value() != name.Value() {
			if !name.Equals(exp.Name()) {
				exists, err := u.experimentRepository.ExistsByName(input.CallerID, name)
				if err != nil {
					return err
				}
				if exists {
					return &ExperimentNameAlreadyExistsError{}
				}
			}
			exp.SetName(name)
			needToSave = true
		}
	}

	if input.Description != nil {
		description, err := experiment.NewDescription(*input.Description)
		if err != nil {
			return err
		}
		if exp.Description() != description {
			exp.SetDescription(description)
			needToSave = true
		}
	}

	if input.TrainConfigID != nil && *input.TrainConfigID != exp.TrainConfigID() {
		checkTrainConfigOwnership, err := u.trainConfigRepository.CheckOwnership(input.CallerID, *input.TrainConfigID)
		if err != nil {
			return err
		}
		if !checkTrainConfigOwnership {
			return &UnauthorizedError{}
		}
		exp.SetTrainConfigID(*input.TrainConfigID)
		needToSave = true
	}

	if input.VQCConfigID != nil && *input.VQCConfigID != exp.VQCConfigID() {
		checkVQCConfigOwnership, err := u.vqcConfigRepository.CheckOwnership(input.CallerID, *input.VQCConfigID)
		if err != nil {
			return err
		}
		if !checkVQCConfigOwnership {
			return &UnauthorizedError{}
		}
		exp.SetVQCConfigID(*input.VQCConfigID)
		needToSave = true
	}

	if needToSave {
		err = u.experimentRepository.Save(exp)
		if err != nil {
			return err
		}
	}

	return nil
}

func canUpdateExperiment(callerID uuid.UUID, exp *experiment.Experiment) bool {
	return callerID == exp.OwnerID()
}
