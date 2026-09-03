package testkit

import (
	"pennylane_project_backend/internal/domain/user"
)

type UserSeed struct {
	Name  string
	Email string
	Role  user.Role
	Ref   uint8
}

type UserSeedResult struct {
	ByRef map[uint8]*user.User
}

func SeedUsers(repo *MockUserRepository, seeds []UserSeed) (UserSeedResult, error) {
	res := UserSeedResult{ByRef: map[uint8]*user.User{}}
	for _, s := range seeds {
		u, err := user.NewUser(s.Name, s.Email)
		if err != nil {
			return UserSeedResult{}, err
		}
		u.SetRole(s.Role)
		if err := repo.Save(u); err != nil {
			return UserSeedResult{}, err
		}
		res.ByRef[s.Ref] = u
	}
	return res, nil
}
