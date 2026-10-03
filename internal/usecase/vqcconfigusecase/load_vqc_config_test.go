package vqcconfigusecase

import (
	"errors"
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"pennylane_project_backend/internal/usecase/apperrors"
	"testing"

	"uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadVQCConfigByID(t *testing.T) {
	tests := []struct {
		setup    func(f *testFixture) (LoadVQCConfigByIDInput, error)
		testName string
	}{
		{
			testName: "load VQCConfig by ID successfully",
			setup: func(f *testFixture) (LoadVQCConfigByIDInput, error) {
				newUser := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(newUser.ID())
				return LoadVQCConfigByIDInput{
					CallerID:    newUser.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, nil
			},
		},
		{
			testName: "inexistent VQCConfig ID",
			setup: func(f *testFixture) (LoadVQCConfigByIDInput, error) {
				newUser := f.createUser(user.RoleUser)
				return LoadVQCConfigByIDInput{
					CallerID:    newUser.ID(),
					VQCConfigID: uuid.New(),
				}, vqcconfig.ErrVQCConfigNotFound
			},
		},
		{
			testName: "user trying to load a VQCConfig they do not own",
			setup: func(f *testFixture) (LoadVQCConfigByIDInput, error) {
				owner := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(owner.ID())
				unauthorizedUser := f.createUser(user.RoleUser)
				return LoadVQCConfigByIDInput{
					CallerID:    unauthorizedUser.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, apperrors.ErrPermissionDenied
			},
		},
		{
			testName: "inexistent caller user",
			setup: func(f *testFixture) (LoadVQCConfigByIDInput, error) {
				newUser := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(newUser.ID())
				return LoadVQCConfigByIDInput{
					CallerID:    uuid.New(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, apperrors.ErrPermissionDenied
			},
		},
		{
			testName: "nil caller ID",
			setup: func(f *testFixture) (LoadVQCConfigByIDInput, error) {
				newUser := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(newUser.ID())
				return LoadVQCConfigByIDInput{
					CallerID:    uuid.Nil(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, user.ErrNilUserID
			},
		},
		{
			testName: "nil VQCConfig ID",
			setup: func(f *testFixture) (LoadVQCConfigByIDInput, error) {
				newUser := f.createUser(user.RoleUser)
				return LoadVQCConfigByIDInput{
					CallerID:    newUser.ID(),
					VQCConfigID: uuid.Nil(),
				}, vqcconfig.ErrNilVQCConfigID
			},
		},
		{
			testName: "fail to find VQCConfig by ID due to repository error",
			setup: func(f *testFixture) (LoadVQCConfigByIDInput, error) {
				newUser := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(newUser.ID())
				dbErr := errors.New("database unavailable")
				f.vqcConfigRepo.ErrFindByID = dbErr
				return LoadVQCConfigByIDInput{
					CallerID:    newUser.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, dbErr
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			fixture := newTestFixture(t)
			input, expectedErr := tt.setup(fixture)
			output, receivedErr := fixture.service.LoadVQCConfigByID(input)

			if expectedErr != nil {
				assert.ErrorIs(t, receivedErr, expectedErr)
				assert.Nil(t, output)
			} else {
				assert.NoError(t, receivedErr)
				assert.NotNil(t, output)
				assert.Equal(t, input.CallerID, output.OwnerID)
				assert.Equal(t, input.VQCConfigID, output.VQCConfigID)
			}
		})
	}
}

func TestLoadVQCConfigByName(t *testing.T) {
	tests := []struct {
		setup    func(f *testFixture) (LoadVQCConfigByNameInput, error)
		testName string
	}{
		{
			testName: "load VQCConfig by name successfully",
			setup: func(f *testFixture) (LoadVQCConfigByNameInput, error) {
				newUser := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(newUser.ID())
				return LoadVQCConfigByNameInput{
					CallerID:      newUser.ID(),
					VQCConfigName: vqcConfig.Name().String(),
				}, nil
			},
		},
		{
			testName: "inexistent VQCConfig name",
			setup: func(f *testFixture) (LoadVQCConfigByNameInput, error) {
				newUser := f.createUser(user.RoleUser)
				return LoadVQCConfigByNameInput{
					CallerID:      newUser.ID(),
					VQCConfigName: "NonExistentConfig",
				}, vqcconfig.ErrVQCConfigNotFound
			},
		},
		{
			testName: "user trying to load a VQCConfig they do not own",
			setup: func(f *testFixture) (LoadVQCConfigByNameInput, error) {
				owner := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(owner.ID())
				unauthorizedUser := f.createUser(user.RoleUser)
				return LoadVQCConfigByNameInput{
					CallerID:      unauthorizedUser.ID(),
					VQCConfigName: vqcConfig.Name().String(),
				}, vqcconfig.ErrVQCConfigNotFound
			},
		},
		{
			testName: "inexistent caller user",
			setup: func(f *testFixture) (LoadVQCConfigByNameInput, error) {
				newUser := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(newUser.ID())
				return LoadVQCConfigByNameInput{
					CallerID:      uuid.New(),
					VQCConfigName: vqcConfig.Name().String(),
				}, vqcconfig.ErrVQCConfigNotFound
			},
		},
		{
			testName: "nil caller ID",
			setup: func(f *testFixture) (LoadVQCConfigByNameInput, error) {
				newUser := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(newUser.ID())
				return LoadVQCConfigByNameInput{
					CallerID:      uuid.Nil(),
					VQCConfigName: vqcConfig.Name().String(),
				}, user.ErrNilUserID
			},
		},
		{
			testName: "invalid VQCConfig name",
			setup: func(f *testFixture) (LoadVQCConfigByNameInput, error) {
				newUser := f.createUser(user.RoleUser)
				return LoadVQCConfigByNameInput{
					CallerID:      newUser.ID(),
					VQCConfigName: "",
				}, vqcconfig.ErrInvalidName
			},
		},
		{
			testName: "fail to find VQCConfig by name due to repository error",
			setup: func(f *testFixture) (LoadVQCConfigByNameInput, error) {
				newUser := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(newUser.ID())
				dbErr := errors.New("database unavailable")
				f.vqcConfigRepo.ErrFindByName = dbErr
				return LoadVQCConfigByNameInput{
					CallerID:      newUser.ID(),
					VQCConfigName: vqcConfig.Name().String(),
				}, dbErr
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			fixture := newTestFixture(t)
			input, expectedErr := tt.setup(fixture)
			output, receivedErr := fixture.service.LoadVQCConfigByName(input)

			if expectedErr != nil {
				assert.ErrorIs(t, receivedErr, expectedErr)
				assert.Nil(t, output)
			} else {
				assert.NoError(t, receivedErr)
				assert.NotNil(t, output)
				assert.Equal(t, input.CallerID, output.OwnerID)
				assert.Equal(t, input.VQCConfigName, output.Name)
			}
		})
	}
}

func TestLoadAllVQCConfigs(t *testing.T) {
	tests := []struct {
		setup              func(f *testFixture) LoadAllVQCConfigsInput
		testName           string
		expectedSizeOutput int
	}{
		{
			testName: "load all VQCConfigs successfully",
			setup: func(f *testFixture) LoadAllVQCConfigsInput {
				user := f.createUser(user.RoleUser)
				f.createVQCConfig(user.ID())
				f.createVQCConfig(user.ID())
				return LoadAllVQCConfigsInput{
					CallerID: user.ID(),
				}
			},
			expectedSizeOutput: 2,
		},
		{
			testName: "inexistent caller user",
			setup: func(f *testFixture) LoadAllVQCConfigsInput {
				return LoadAllVQCConfigsInput{
					CallerID: uuid.New(),
				}
			},
			expectedSizeOutput: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			fixture := newTestFixture(t)
			input := tt.setup(fixture)
			output, err := fixture.service.LoadAllVQCConfigs(input)
			require.NoError(t, err)
			assert.NotNil(t, output)
			assert.Equal(t, tt.expectedSizeOutput, len(output.VQCConfigs))
		})
	}
}
