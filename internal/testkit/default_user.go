package testkit

import (
	"pennylane_project_backend/internal/domain/user"
	"strconv"
)

func DefaultUser() func() *user.User {
	count := 0
	return func() *user.User {
		count++
		name, err := user.NewName("Test User " + strconv.Itoa(count))
		if err != nil {
			panic(err)
		}
		email, err := user.NewEmail("testuser" + strconv.Itoa(count) + "@example.com")
		if err != nil {
			panic(err)
		}
		u := user.NewUser(name, email)
		return u
	}
}
