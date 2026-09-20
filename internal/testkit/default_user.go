package testkit

import (
	"pennylane_project_backend/internal/domain/user"
	"strconv"
)

func DefaultUser() func() *user.User {
	count := 0
	return func() *user.User {
		count++
		name, err := user.NewName("Test User " + counterToLetters(count))
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

func counterToLetters(count int) string {
	// Convert a counter to a string of letters
	// For example, 1 -> "a", 2 -> "b", ..., 26 -> "z", 27 -> "aa", etc.
	if count <= 0 {
		return ""
	}
	result := ""
	for count > 0 {
		result = string(rune('a'+(count-1)%26)) + result
		count = (count - 1) / 26
	}
	return result
}
