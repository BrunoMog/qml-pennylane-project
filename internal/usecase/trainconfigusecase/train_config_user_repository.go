package trainconfigusecase

import (
	"github.com/google/uuid"
)

type UserRepository interface {
	ExistsByID(userID uuid.UUID) (bool, error)
}
