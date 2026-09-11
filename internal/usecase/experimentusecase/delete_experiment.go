package experimentusecase

import (
	"uuid"
)

type DeleteExperimentInput struct {
	ExperimentID uuid.UUID
	CallerID     uuid.UUID
}

func (s *ExperimentService) DeleteExperiment(input DeleteExperimentInput) error {
	checkOwnership, err := s.experimentRepository.CheckOwnership(input.CallerID, input.ExperimentID)
	if err != nil {
		return err
	}

	if !checkOwnership {
		return &UnauthorizedError{}
	}

	err = s.experimentRepository.DeleteByID(input.ExperimentID)
	if err != nil {
		return err
	}

	return nil
}
