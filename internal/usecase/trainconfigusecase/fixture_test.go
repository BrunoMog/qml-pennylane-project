package trainconfigusecase

import (
	"pennylane_project_backend/internal/domain/trainconfig"
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type testFixture struct {
	t               *testing.T
	service         *TrainConfigService
	userRepo        *testkit.MockUserRepository
	trainConfigRepo *testkit.MockTrainConfigRepository
}

func newTestFixture(t *testing.T) *testFixture {
	t.Helper()
	userRepo := testkit.NewMockUserRepository()
	trainConfigRepo := testkit.NewMockTrainConfigRepository()
	service := NewTrainConfigService(trainConfigRepo, userRepo)

	return &testFixture{
		t:               t,
		service:         service,
		userRepo:        userRepo,
		trainConfigRepo: trainConfigRepo,
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

func (f *testFixture) createTrainConfig(ownerID uuid.UUID) *trainconfig.TrainConfig {
	makeTrainConfig := testkit.DefaultTrainConfig()
	trainConfig := makeTrainConfig(ownerID)
	err := f.trainConfigRepo.Save(trainConfig)
	require.NoError(f.t, err)
	return trainConfig
}
