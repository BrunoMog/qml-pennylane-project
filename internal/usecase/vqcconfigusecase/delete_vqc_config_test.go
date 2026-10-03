package vqcconfigusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"testing"

	"errors"
	"pennylane_project_backend/internal/usecase/apperrors"
	"uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteVQCConfig(t *testing.T) {
	tests := []struct {
		setup    func(f *testFixture) (DeleteVQCConfigInput, error)
		testName string
	}{
		{
			testName: "delete VQCConfig successfully",
			setup: func(f *testFixture) (DeleteVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				return DeleteVQCConfigInput{
					CallerID:    user.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, nil
			},
		},
		{
			testName: "inexistent caller user",
			setup: func(f *testFixture) (DeleteVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				return DeleteVQCConfigInput{
					CallerID:    uuid.New(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, apperrors.ErrPermissionDenied
			},
		},
		{
			testName: "inexistent VQCConfig",
			setup: func(f *testFixture) (DeleteVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				return DeleteVQCConfigInput{
					CallerID:    user.ID(),
					VQCConfigID: uuid.New(),
				}, vqcconfig.ErrVQCConfigNotFound
			},
		},
		{
			testName: "permission denied due to unauthorized user",
			setup: func(f *testFixture) (DeleteVQCConfigInput, error) {
				owner := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(owner.ID())
				unauthorizedUser := f.createUser(user.RoleUser)
				return DeleteVQCConfigInput{
					CallerID:    unauthorizedUser.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, apperrors.ErrPermissionDenied
			},
		},
		{
			testName: "fail to check ownership due to repository error",
			setup: func(f *testFixture) (DeleteVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				dbErr := errors.New("database error")
				f.vqcConfigRepo.ErrCheckOwnership = dbErr
				return DeleteVQCConfigInput{
					CallerID:    user.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, dbErr
			},
		},
		{
			testName: "fail to delete VQCConfig due to repository error",
			setup: func(f *testFixture) (DeleteVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				dbErr := errors.New("database error")
				f.vqcConfigRepo.ErrDeleteByID = dbErr
				return DeleteVQCConfigInput{
					CallerID:    user.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, dbErr
			},
		},
		{
			testName: "fail to delete VQCConfig due to nil caller ID",
			setup: func(f *testFixture) (DeleteVQCConfigInput, error) {
				newUser := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(newUser.ID())
				return DeleteVQCConfigInput{
					CallerID:    uuid.Nil(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, user.ErrNilUserID
			},
		},
		{
			testName: "fail to delete VQCConfig due to nil VQCConfig ID",
			setup: func(f *testFixture) (DeleteVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				return DeleteVQCConfigInput{
					CallerID:    user.ID(),
					VQCConfigID: uuid.Nil(),
				}, vqcconfig.ErrNilVQCConfigID
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			fixture := newTestFixture(t)
			input, expectedError := tt.setup(fixture)
			receivedError := fixture.service.DeleteVQCConfig(input)
			if expectedError != nil {
				assert.ErrorIs(t, receivedError, expectedError)
			} else {
				assert.NoError(t, receivedError)
				exists, err := fixture.vqcConfigRepo.ExistsByID(input.VQCConfigID)
				require.NoError(t, err)
				assert.False(t, exists)
			}
		})
	}

}
