package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type testFixture struct {
	t             *testing.T
	service       *VQCConfigService
	userRepo      *testkit.MockUserRepository
	vqcConfigRepo *testkit.MockVQCConfigRepository
}

func newTestFixture(t *testing.T) *testFixture {
	t.Helper()
	userRepo := testkit.NewMockUserRepository()
	vqcConfigRepo := testkit.NewMockVQCConfigRepository()
	service := NewVQCConfigService(vqcConfigRepo, userRepo)

	return &testFixture{
		t:             t,
		service:       service,
		userRepo:      userRepo,
		vqcConfigRepo: vqcConfigRepo,
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

func (f *testFixture) createVQCConfig(ownerID uuid.UUID) *vqcconfig.VQCConfig {
	makeVQCConfig := testkit.DefaultVQCConfig()
	vqcConfig := makeVQCConfig(ownerID)
	err := f.vqcConfigRepo.Save(vqcConfig)
	require.NoError(f.t, err)
	return vqcConfig
}
