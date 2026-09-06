package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateVQCConfig(t *testing.T) {
	tests := []struct {
		expectedError error
		setup         func(f *testFixture) UpdateVQCConfigInput
		testName      string
	}{
		{
			testName: "update VQCConfig successfully",
			setup: func(f *testFixture) UpdateVQCConfigInput {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				newName := "Updated Config Name"
				newDescription := "Updated Description"
				return UpdateVQCConfigInput{
					Name:        &newName,
					Description: &newDescription,
					CallerID:    user.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}
			},
			expectedError: nil,
		},
		{
			testName: "inexistent caller user",
			setup: func(f *testFixture) UpdateVQCConfigInput {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				newName := "Updated Config Name"
				return UpdateVQCConfigInput{
					Name:        &newName,
					Description: nil,
					CallerID:    uuid.New(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "inexistent VQCConfig",
			setup: func(f *testFixture) UpdateVQCConfigInput {
				user := f.createUser(user.RoleUser)
				newName := "Updated Config Name"
				return UpdateVQCConfigInput{
					Name:        &newName,
					Description: nil,
					CallerID:    user.ID(),
					VQCConfigID: uuid.New(),
				}
			},
			expectedError: &testkit.ErrVQCConfigNotFound{},
		},
		{
			testName: "unauthorized user",
			setup: func(f *testFixture) UpdateVQCConfigInput {
				owner := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(owner.ID())
				unauthorizedUser := f.createUser(user.RoleUser)
				newName := "Updated Config Name"
				return UpdateVQCConfigInput{
					Name:        &newName,
					Description: nil,
					CallerID:    unauthorizedUser.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "no fields to update",
			setup: func(f *testFixture) UpdateVQCConfigInput {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				return UpdateVQCConfigInput{
					Name:        nil,
					Description: nil,
					CallerID:    user.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}
			},
			expectedError: &NoFieldsToUpdateError{},
		},
		{
			testName: "try to update with existing name",
			setup: func(f *testFixture) UpdateVQCConfigInput {
				user := f.createUser(user.RoleUser)
				vqcConfig1 := f.createVQCConfig(user.ID())
				vqcConfig2 := f.createVQCConfig(user.ID())
				newName := vqcConfig2.Name()
				return UpdateVQCConfigInput{
					Name:        &newName,
					Description: nil,
					CallerID:    user.ID(),
					VQCConfigID: vqcConfig1.VQCConfigID(),
				}
			},
			expectedError: &VQCConfigNameAlreadyExistsError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := newTestFixture(t)
			input := tt.setup(f)
			err := f.service.UpdateVQCConfig(input)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				vqcConfig, err := f.vqcConfigRepo.FindByID(input.VQCConfigID)
				require.NoError(t, err)
				if input.Name != nil {
					assert.Equal(t, *input.Name, vqcConfig.Name())
				}
				if input.Description != nil {
					assert.Equal(t, *input.Description, vqcConfig.Description())
				}
			}
		})
	}
}
