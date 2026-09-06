package trainconfigusecase

import (
	"github.com/google/uuid"
)

type DeleteTrainConfigInput struct {
	CallerID      uuid.UUID
	TrainConfigID uuid.UUID
}

func (s *TrainConfigService) DeleteTrainConfig(input DeleteTrainConfigInput) error {
	checkOwnership, err := s.trainConfigRepository.CheckOwnership(input.CallerID, input.TrainConfigID)
	if err != nil {
		return err
	}

	if !checkOwnership {
		return &UnauthorizedError{}
	}

	err = s.trainConfigRepository.DeleteByID(input.TrainConfigID)
	if err != nil {
		return err
	}

	return nil
}
