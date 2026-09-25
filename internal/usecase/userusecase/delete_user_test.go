package userusecase

import (
	"errors"
	"pennylane_project_backend/internal/domain/user"
	"testing"

	"uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		setup    func(fixture *testFixture) (DeleteUserInput, error)
		testName string
	}{
		{
			testName: "owner deletes admin",
			setup: func(fixture *testFixture) (DeleteUserInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				return DeleteUserInput{
					CallerID: owner.ID(),
					TargetID: admin.ID(),
				}, nil
			},
		},
		{
			testName: "user deletes self",
			setup: func(fixture *testFixture) (DeleteUserInput, error) {
				newUser := fixture.createUser(user.RoleUser)
				return DeleteUserInput{
					CallerID: newUser.ID(),
					TargetID: newUser.ID(),
				}, nil
			},
		},
		{
			testName: "inexistent target user",
			setup: func(fixture *testFixture) (DeleteUserInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				return DeleteUserInput{
					CallerID: owner.ID(),
					TargetID: uuid.New(),
				}, user.ErrUserNotFound
			},
		},
		{
			testName: "inexistent caller user",
			setup: func(fixture *testFixture) (DeleteUserInput, error) {
				newUser := fixture.createUser(user.RoleUser)
				return DeleteUserInput{
					CallerID: uuid.New(),
					TargetID: newUser.ID(),
				}, user.ErrUserNotFound
			},
		},
		{
			testName: "admin tries to delete owner",
			setup: func(fixture *testFixture) (DeleteUserInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				return DeleteUserInput{
					CallerID: admin.ID(),
					TargetID: owner.ID(),
				}, ErrPermissionDenied
			},
		},
		{
			testName: "owner try to delete self",
			setup: func(fixture *testFixture) (DeleteUserInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				return DeleteUserInput{
					CallerID: owner.ID(),
					TargetID: owner.ID(),
				}, ErrPermissionDenied
			},
		},
		{
			testName: "fail to retrieve caller user",
			setup: func(fixture *testFixture) (DeleteUserInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				dbErr := errors.New("database unavailable")
				fixture.userRepo.FindByIDErr = dbErr
				return DeleteUserInput{
					CallerID: owner.ID(),
					TargetID: admin.ID(),
				}, dbErr
			},
		},
		{
			testName: "fail to retrieve target user",
			setup: func(fixture *testFixture) (DeleteUserInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				dbErr := errors.New("database unavailable")
				fixture.userRepo.FindByIDErr = dbErr
				return DeleteUserInput{
					CallerID: owner.ID(),
					TargetID: admin.ID(),
				}, dbErr
			},
		},
		{
			testName: "fail to delete user",
			setup: func(fixture *testFixture) (DeleteUserInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				dbErr := errors.New("database unavailable")
				fixture.userRepo.DeleteByIDErr = dbErr
				return DeleteUserInput{
					CallerID: owner.ID(),
					TargetID: admin.ID(),
				}, dbErr
			},
		},
		{
			testName: "nil caller ID",
			setup: func(fixture *testFixture) (DeleteUserInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				return DeleteUserInput{
					CallerID: uuid.Nil(),
					TargetID: owner.ID(),
				}, ErrNilID
			},
		},
		{
			testName: "nil target ID",
			setup: func(fixture *testFixture) (DeleteUserInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				return DeleteUserInput{
					CallerID: owner.ID(),
					TargetID: uuid.Nil(),
				}, ErrNilID
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			fixture := newTestFixture(t)

			input, expectedErr := tt.setup(fixture)

			receivedErr := fixture.service.DeleteUser(input)
			if expectedErr != nil {
				assert.ErrorIs(t, receivedErr, expectedErr)
			} else {
				assert.NoError(t, receivedErr)
				exists, err := fixture.userRepo.ExistsByID(input.TargetID)
				require.NoError(t, err)
				assert.False(t, exists)
			}
		})
	}

}
