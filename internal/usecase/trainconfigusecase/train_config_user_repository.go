package trainconfigusecase

import (
	"uuid"
)

type UserRepository interface {
	ExistsByID(userID uuid.UUID) (bool, error)
}
