package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/trainconfig"

	"github.com/google/uuid"
)

type DeleteTrainConfigInput struct {
	CallerID      uuid.UUID
	TrainConfigID uuid.UUID
}

func (s *TrainConfigService) DeleteTrainConfig(input DeleteTrainConfigInput) error {
	exists, err := s.userRepository.ExistsByID(input.CallerID)
	if err != nil {
		return err
	}
	if !exists {
		return &UserNotFoundError{}
	}

	trainConfig, err := s.trainConfigRepository.FindByID(input.TrainConfigID)
	if err != nil {
		return err
	}

	if !canDeleteTrainConfig(input.CallerID, trainConfig) {
		return &UnauthorizedError{}
	}

	err = s.trainConfigRepository.DeleteByID(input.TrainConfigID)
	if err != nil {
		return err
	}

	return nil
}

func canDeleteTrainConfig(callerID uuid.UUID, trainConfig *trainconfig.TrainConfig) bool {
	return callerID == trainConfig.OwnerID()
}
