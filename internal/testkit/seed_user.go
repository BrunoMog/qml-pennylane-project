package testkit

import (
	"pennylane_project_backend/internal/domain/user"
)

type UserSeed struct {
	Ref   uint8
	Name  string
	Email string
	Role  user.Role
}

type SeedResult struct {
	ByRef map[uint8]*user.User
}

func SeedUsers(repo *MockUserRepository, seeds []UserSeed) (SeedResult, error) {
	res := SeedResult{ByRef: map[uint8]*user.User{}}
	for _, s := range seeds {
		u, err := user.NewUser(s.Name, s.Email)
		if err != nil {
			return SeedResult{}, err
		}
		u.SetRole(s.Role)
		if err := repo.Save(u); err != nil {
			return SeedResult{}, err
		}
		res.ByRef[s.Ref] = u
	}
	return res, nil
}
