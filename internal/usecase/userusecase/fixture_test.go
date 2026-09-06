package userusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/stretchr/testify/require"
)

type testFixture struct {
	t        *testing.T
	service  *UserService
	userRepo *testkit.MockUserRepository
}

func newTestFixture(t *testing.T) *testFixture {
	t.Helper()
	userRepo := testkit.NewMockUserRepository()
	service := NewUserService(userRepo)

	return &testFixture{
		t:        t,
		service:  service,
		userRepo: userRepo,
	}
}

func (f *testFixture) createUser(role user.Role) *user.User {
	makeUser := testkit.DefaultUser()
	user := makeUser()
	user.SetRole(role)
	err := f.userRepo.Save(user)
	require.NoError(f.t, err)
	return user
}
