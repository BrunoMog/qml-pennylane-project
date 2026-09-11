package experimentusecase

import (
	"uuid"
)

type UserRepository interface {
	ExistsByID(id uuid.UUID) (bool, error)
}
