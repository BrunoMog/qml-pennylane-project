package testkit

import (
	"pennylane_project_backend/internal/domain/user"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func DefaultUser(t testing.TB) func() *user.User {
	t.Helper()

	count := 0

	return func() *user.User {
		t.Helper()
		count++

		name, err := user.NewName("Test User " + counterToLetters(count))
		require.NoError(t, err)

		email, err := user.NewEmail("testuser" + strconv.Itoa(count) + "@example.com")
		require.NoError(t, err)

		u, err := user.NewUser(name, email)
		require.NoError(t, err)

		return u
	}
}

func counterToLetters(count int) string {
	// Convert a counter to a string of letters
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
