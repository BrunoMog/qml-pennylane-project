package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"strings"
	"testing"

	"errors"
	"uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateVQCConfig(t *testing.T) {
	tests := []struct {
		setup    func(f *testFixture) (CreateVQCConfigInput, error)
		testName string
	}{
		{
			testName: "create VQCConfig successfully",
			setup: func(f *testFixture) (CreateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				return CreateVQCConfigInput{
					Name:        "Test Config",
					Description: "This is a test VQCConfig",
					CallerID:    user.ID(),
					VQCDTO:      validVQCDTO(),
				}, nil
			},
		},
		{
			testName: "fail to create VQCConfig due to invalid name",
			setup: func(f *testFixture) (CreateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				return CreateVQCConfigInput{
					Name:        "",
					Description: "This is a test VQCConfig",
					CallerID:    user.ID(),
					VQCDTO:      validVQCDTO(),
				}, vqcconfig.ErrInvalidName
			},
		},
		{
			testName: "fail to create VQCConfig due to invalid description",
			setup: func(f *testFixture) (CreateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				return CreateVQCConfigInput{
					Name:        "Test Config",
					Description: strings.Repeat("a", 501), // Exceeding max length
					CallerID:    user.ID(),
					VQCDTO:      validVQCDTO(),
				}, vqcconfig.ErrInvalidDescription
			},
		},
		{
			testName: "fail to create VQCConfig due to invalid VQC",
			setup: func(f *testFixture) (CreateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				return CreateVQCConfigInput{
					Name:        "Test Config",
					Description: "This is a test VQCConfig",
					CallerID:    user.ID(),
					VQCDTO:      invalidVQCDTO(),
				}, vqc.ErrInvalidQubit
			},
		},
		{
			testName: "fail to create VQCConfig due to non-existent user",
			setup: func(f *testFixture) (CreateVQCConfigInput, error) {
				return CreateVQCConfigInput{
					Name:        "Test Config",
					Description: "This is a test VQCConfig",
					CallerID:    uuid.New(),
					VQCDTO:      validVQCDTO(),
				}, user.ErrUserNotFound
			},
		},
		{
			testName: "fail to create VQCConfig due to duplicate name",
			setup: func(f *testFixture) (CreateVQCConfigInput, error) {
				newUser := f.createUser(user.RoleUser)
				newVQCConfig := f.createVQCConfig(newUser.ID())
				name, err := vqcconfig.NewName("Test Config")
				require.NoError(t, err)
				newVQCConfig.SetName(name)
				return CreateVQCConfigInput{
					Name:        name.String(),
					Description: "This is a test VQCConfig",
					CallerID:    newUser.ID(),
					VQCDTO:      validVQCDTO(),
				}, ErrVQCConfigNameAlreadyExists
			},
		},
		{
			testName: "fail to create VQCConfig due to nil caller ID",
			setup: func(f *testFixture) (CreateVQCConfigInput, error) {
				return CreateVQCConfigInput{
					Name:        "Test Config",
					Description: "This is a test VQCConfig",
					CallerID:    uuid.Nil(),
					VQCDTO:      validVQCDTO(),
				}, user.ErrNilUserID
			},
		},
		{
			testName: "fail to check user existence due to repository error",
			setup: func(f *testFixture) (CreateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				dbErr := errors.New("database error")
				f.userRepo.ErrExistsByID = dbErr
				return CreateVQCConfigInput{
					Name:        "Test Config",
					Description: "This is a test VQCConfig",
					CallerID:    user.ID(),
					VQCDTO:      validVQCDTO(),
				}, dbErr
			},
		},
		{
			testName: "fail to check VQCConfig name existence due to repository error",
			setup: func(f *testFixture) (CreateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				dbErr := errors.New("database error")
				f.vqcConfigRepo.ErrExistsByName = dbErr
				return CreateVQCConfigInput{
					Name:        "Test Config",
					Description: "This is a test VQCConfig",
					CallerID:    user.ID(),
					VQCDTO:      validVQCDTO(),
				}, dbErr
			},
		},
		{
			testName: "fail to save VQCConfig due to repository error",
			setup: func(f *testFixture) (CreateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				dbErr := errors.New("database error")
				f.vqcConfigRepo.ErrSave = dbErr
				return CreateVQCConfigInput{
					Name:        "Test Config",
					Description: "This is a test VQCConfig",
					CallerID:    user.ID(),
					VQCDTO:      validVQCDTO(),
				}, dbErr
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := newTestFixture(t)
			input, expectedError := tt.setup(f)

			output, receivedError := f.service.CreateVQCConfig(input)

			if expectedError != nil {
				assert.ErrorIs(t, receivedError, expectedError)
				assert.Nil(t, output)
			} else {
				assert.NoError(t, receivedError)
				assert.NotNil(t, output)
				assert.Equal(t, input.Name, output.Name)
				assert.Equal(t, input.Description, output.Description)
				assert.NotEqual(t, uuid.Nil(), output.VQCConfigID)
			}
		})
	}

}
