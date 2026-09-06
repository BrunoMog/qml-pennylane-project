package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteVQCConfig(t *testing.T) {
	tests := []struct {
		expectedError error
		setup         func(f *testFixture) DeleteVQCConfigInput
		testName      string
	}{
		{
			testName: "delete VQCConfig successfully",
			setup: func(f *testFixture) DeleteVQCConfigInput {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				return DeleteVQCConfigInput{
					CallerID:    user.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}
			},
			expectedError: nil,
		},
		{
			testName: "inexistent caller user",
			setup: func(f *testFixture) DeleteVQCConfigInput {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				return DeleteVQCConfigInput{
					CallerID:    uuid.New(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "inexistent VQCConfig",
			setup: func(f *testFixture) DeleteVQCConfigInput {
				user := f.createUser(user.RoleUser)
				return DeleteVQCConfigInput{
					CallerID:    user.ID(),
					VQCConfigID: uuid.New(),
				}
			},
			expectedError: &testkit.ErrVQCConfigNotFound{},
		},
		{
			testName: "unauthorized user",
			setup: func(f *testFixture) DeleteVQCConfigInput {
				owner := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(owner.ID())
				unauthorizedUser := f.createUser(user.RoleUser)
				return DeleteVQCConfigInput{
					CallerID:    unauthorizedUser.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}
			},
			expectedError: &UnauthorizedError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			fixture := newTestFixture(t)
			input := tt.setup(fixture)
			err := fixture.service.DeleteVQCConfig(input)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				exists, err := fixture.vqcConfigRepo.ExistsByID(input.VQCConfigID)
				assert.NoError(t, err)
				assert.False(t, exists)
			}
		})
	}

}
