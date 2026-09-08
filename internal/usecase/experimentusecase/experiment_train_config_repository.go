package experimentusecase

import (
	"github.com/google/uuid"
)

type TrainConfigRepository interface {
	CheckOwnership(userID uuid.UUID, trainConfigID uuid.UUID) (bool, error)
}
