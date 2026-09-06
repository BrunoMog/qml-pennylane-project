package testkit

import (
	"pennylane_project_backend/internal/domain/user"
	"strconv"
)

func DefaultUser() func() *user.User {
	count := 0
	return func() *user.User {
		count++
		u, err := user.NewUser(
			"Test User "+strconv.Itoa(count),
			"testuser"+strconv.Itoa(count)+"@example.com",
		)
		if err != nil {
			panic(err)
		}
		return u
	}
}
