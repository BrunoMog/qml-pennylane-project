package experimentusecase

import (
	"github.com/google/uuid"
)

type UserRepository interface {
	ExistsByID(id uuid.UUID) (bool, error)
}
