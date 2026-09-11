package experimentusecase

import (
	"pennylane_project_backend/internal/domain/experiment"
	"pennylane_project_backend/internal/domain/trainconfig"
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"pennylane_project_backend/internal/testkit"
	"testing"
	"uuid"

	"github.com/stretchr/testify/require"
)

type testFixture struct {
	t               *testing.T
	service         *ExperimentService
	userRepo        *testkit.MockUserRepository
	trainConfigRepo *testkit.MockTrainConfigRepository
	vqcConfigRepo   *testkit.MockVQCConfigRepository
	experimentRepo  *testkit.MockExperimentRepository
	makeUser        func() *user.User
	makeTrainConfig func(ownerID uuid.UUID) *trainconfig.TrainConfig
	makeVQCConfig   func(ownerID uuid.UUID) *vqcconfig.VQCConfig
	makeExperiment  func(ownerID, vqcConfigID, trainConfigID uuid.UUID) *experiment.Experiment
}

func newTestFixture(t *testing.T) *testFixture {
	t.Helper()

	userRepo := testkit.NewMockUserRepository()
	trainConfigRepo := testkit.NewMockTrainConfigRepository()
	vqcConfigRepo := testkit.NewMockVQCConfigRepository()
	experimentRepo := testkit.NewMockExperimentRepository()

	service := NewExperimentService(experimentRepo, vqcConfigRepo, trainConfigRepo, userRepo)

	makeUser := testkit.DefaultUser()
	makeTrainConfig := testkit.DefaultTrainConfig()
	makeVQCConfig := testkit.DefaultVQCConfig()
	makeExperiment := testkit.DefaultExperiment()

	return &testFixture{
		t:               t,
		service:         service,
		userRepo:        userRepo,
		trainConfigRepo: trainConfigRepo,
		vqcConfigRepo:   vqcConfigRepo,
		experimentRepo:  experimentRepo,
		makeUser:        makeUser,
		makeTrainConfig: makeTrainConfig,
		makeVQCConfig:   makeVQCConfig,
		makeExperiment:  makeExperiment,
	}
}

func (f *testFixture) createUser() *user.User {
	f.t.Helper()
	user := f.makeUser()
	err := f.userRepo.Save(user)
	require.NoError(f.t, err)
	return user
}

func (f *testFixture) createTrainConfig(ownerID uuid.UUID) *trainconfig.TrainConfig {
	f.t.Helper()
	trainConfig := f.makeTrainConfig(ownerID)
	err := f.trainConfigRepo.Save(trainConfig)
	require.NoError(f.t, err)
	return trainConfig
}

func (f *testFixture) createVQCConfig(ownerID uuid.UUID) *vqcconfig.VQCConfig {
	f.t.Helper()
	vqcConfig := f.makeVQCConfig(ownerID)
	err := f.vqcConfigRepo.Save(vqcConfig)
	require.NoError(f.t, err)
	return vqcConfig
}

func (f *testFixture) createExperiment(ownerID, vqcConfigID, trainConfigID uuid.UUID) *experiment.Experiment {
	f.t.Helper()
	exp := f.makeExperiment(ownerID, vqcConfigID, trainConfigID)
	err := f.experimentRepo.Save(exp)
	require.NoError(f.t, err)
	return exp
}
