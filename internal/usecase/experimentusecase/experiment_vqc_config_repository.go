package experimentusecase

import (
	"github.com/google/uuid"
)

type VQCConfigRepository interface {
	CheckOwnership(userID uuid.UUID, vqcConfigID uuid.UUID) (bool, error)
}
