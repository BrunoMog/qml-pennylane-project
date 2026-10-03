package vqcconfigusecase

import (
	"errors"
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/domain/vqc"
	"pennylane_project_backend/internal/domain/vqcconfig"
	"pennylane_project_backend/internal/usecase/apperrors"
	"strings"
	"testing"

	"uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateVQCConfig(t *testing.T) {
	tests := []struct {
		setup    func(f *testFixture) (UpdateVQCConfigInput, error)
		testName string
	}{
		{
			testName: "update VQCConfig successfully",
			setup: func(f *testFixture) (UpdateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				newName := "Updated Config Name"
				newDescription := "Updated Description"
				vqc := validVQCDTO()
				return UpdateVQCConfigInput{
					Name:        &newName,
					Description: &newDescription,
					CallerID:    user.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
					VQCDTO:      &vqc,
				}, nil
			},
		},
		{
			testName: "inexistent caller user",
			setup: func(f *testFixture) (UpdateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				newName := "Updated Config Name"
				return UpdateVQCConfigInput{
					Name:        &newName,
					Description: nil,
					CallerID:    uuid.New(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, apperrors.ErrPermissionDenied
			},
		},
		{
			testName: "inexistent VQCConfig",
			setup: func(f *testFixture) (UpdateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				newName := "Updated Config Name"
				return UpdateVQCConfigInput{
					Name:        &newName,
					Description: nil,
					CallerID:    user.ID(),
					VQCConfigID: uuid.New(),
				}, vqcconfig.ErrVQCConfigNotFound
			},
		},
		{
			testName: "permission denied due to unauthorized user",
			setup: func(f *testFixture) (UpdateVQCConfigInput, error) {
				owner := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(owner.ID())
				unauthorizedUser := f.createUser(user.RoleUser)
				newName := "Updated Config Name"
				return UpdateVQCConfigInput{
					Name:        &newName,
					Description: nil,
					CallerID:    unauthorizedUser.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, apperrors.ErrPermissionDenied
			},
		},
		{
			testName: "no fields to update",
			setup: func(f *testFixture) (UpdateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				return UpdateVQCConfigInput{
					Name:        nil,
					Description: nil,
					CallerID:    user.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, ErrNoFieldsToUpdate
			},
		},
		{
			testName: "try to update with invalid name",
			setup: func(f *testFixture) (UpdateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				invalidName := ""
				return UpdateVQCConfigInput{
					Name:        &invalidName,
					Description: nil,
					CallerID:    user.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, vqcconfig.ErrInvalidName
			},
		},
		{
			testName: "try to update with invalid description",
			setup: func(f *testFixture) (UpdateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				invalidDescription := strings.Repeat("a", 501) // Exceeding max length
				return UpdateVQCConfigInput{
					Name:        nil,
					Description: &invalidDescription,
					CallerID:    user.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, vqcconfig.ErrInvalidDescription
			},
		},
		{
			testName: "try to update with invalid VQC",
			setup: func(f *testFixture) (UpdateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				invalidVQCDTO := invalidVQCDTO()
				return UpdateVQCConfigInput{
					Name:        nil,
					Description: nil,
					CallerID:    user.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
					VQCDTO:      &invalidVQCDTO,
				}, vqc.ErrInvalidQubit
			},
		},
		{
			testName: "try to update with existing name",
			setup: func(f *testFixture) (UpdateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				vqcConfig1 := f.createVQCConfig(user.ID())
				vqcConfig2 := f.createVQCConfig(user.ID())
				newName := vqcConfig2.Name().String()
				return UpdateVQCConfigInput{
					Name:        &newName,
					Description: nil,
					CallerID:    user.ID(),
					VQCConfigID: vqcConfig1.VQCConfigID(),
				}, ErrVQCConfigNameAlreadyExists
			},
		},
		{
			testName: "cosmetic update name",
			setup: func(f *testFixture) (UpdateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				name := strings.ToUpper(vqcConfig.Name().String())
				return UpdateVQCConfigInput{
					Name:        &name,
					CallerID:    user.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, nil
			},
		},
		{
			testName: "idepotent update",
			setup: func(f *testFixture) (UpdateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				name := vqcConfig.Name().String()
				description := vqcConfig.Description().String()
				vqcDTO := buildVQCDTOFromVQC(vqcConfig.VQC())
				return UpdateVQCConfigInput{
					Name:        &name,
					Description: &description,
					VQCDTO:      &vqcDTO,
					CallerID:    user.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, nil
			},
		},
		{
			testName: "nil caller ID",
			setup: func(f *testFixture) (UpdateVQCConfigInput, error) {
				newUser := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(newUser.ID())
				newName := "Updated Config Name"
				return UpdateVQCConfigInput{
					Name:        &newName,
					Description: nil,
					CallerID:    uuid.Nil(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, user.ErrNilUserID
			},
		},
		{
			testName: "nil VQCConfig ID",
			setup: func(f *testFixture) (UpdateVQCConfigInput, error) {
				newUser := f.createUser(user.RoleUser)
				newName := "Updated Config Name"
				return UpdateVQCConfigInput{
					Name:        &newName,
					Description: nil,
					CallerID:    newUser.ID(),
					VQCConfigID: uuid.Nil(),
				}, vqcconfig.ErrNilVQCConfigID
			},
		},
		{
			testName: "fail to find VQCConfig due to repository error",
			setup: func(f *testFixture) (UpdateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				dbErr := errors.New("database error")
				f.vqcConfigRepo.ErrFindByID = dbErr
				newName := "Updated Config Name"
				return UpdateVQCConfigInput{
					Name:        &newName,
					Description: nil,
					CallerID:    user.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, dbErr
			},
		},
		{
			testName: "fail to check name existence due to repository error",
			setup: func(f *testFixture) (UpdateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				dbErr := errors.New("database error")
				f.vqcConfigRepo.ErrExistsByName = dbErr
				newName := "Updated Config Name"
				return UpdateVQCConfigInput{
					Name:        &newName,
					Description: nil,
					CallerID:    user.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, dbErr
			},
		},
		{
			testName: "fail to save VQCConfig due to repository error",
			setup: func(f *testFixture) (UpdateVQCConfigInput, error) {
				user := f.createUser(user.RoleUser)
				vqcConfig := f.createVQCConfig(user.ID())
				dbErr := errors.New("database error")
				f.vqcConfigRepo.ErrSave = dbErr
				newName := "Updated Config Name"
				return UpdateVQCConfigInput{
					Name:        &newName,
					Description: nil,
					CallerID:    user.ID(),
					VQCConfigID: vqcConfig.VQCConfigID(),
				}, dbErr
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := newTestFixture(t)
			input, expectedErr := tt.setup(f)
			receivedErr := f.service.UpdateVQCConfig(input)
			if expectedErr != nil {
				assert.ErrorIs(t, receivedErr, expectedErr)
			} else {
				assert.NoError(t, receivedErr)
				vqcConfig, err := f.vqcConfigRepo.FindByID(input.VQCConfigID)
				require.NoError(t, err)
				if input.Name != nil {
					assert.Equal(t, *input.Name, vqcConfig.Name().String())
				}
				if input.Description != nil {
					assert.Equal(t, *input.Description, vqcConfig.Description().String())
				}
			}
		})
	}
}
