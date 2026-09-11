package experimentusecase

import (
	"uuid"
)

type VQCConfigRepository interface {
	CheckOwnership(userID uuid.UUID, vqcConfigID uuid.UUID) (bool, error)
}
