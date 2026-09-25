package userusecase

import (
	"errors"
	"pennylane_project_backend/internal/domain/user"
	"testing"

	"uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChangeUserRole(t *testing.T) {
	tests := []struct {
		setup    func(fixture *testFixture) (ChangeUserRoleInput, error)
		testName string
	}{
		{
			testName: "owner changes admin to user",
			setup: func(fixture *testFixture) (ChangeUserRoleInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				return ChangeUserRoleInput{
					CallerID: owner.ID(),
					TargetID: admin.ID(),
					Role:     "user",
				}, nil
			},
		},
		{
			testName: "admin changes user to admin",
			setup: func(fixture *testFixture) (ChangeUserRoleInput, error) {
				admin := fixture.createUser(user.RoleAdmin)
				newUser := fixture.createUser(user.RoleUser)
				return ChangeUserRoleInput{
					CallerID: admin.ID(),
					TargetID: newUser.ID(),
					Role:     "admin",
				}, nil
			},
		},
		{
			testName: "try to assign invalid role",
			setup: func(fixture *testFixture) (ChangeUserRoleInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				newUser := fixture.createUser(user.RoleUser)
				return ChangeUserRoleInput{
					CallerID: owner.ID(),
					TargetID: newUser.ID(),
					Role:     "invalid_role",
				}, user.ErrInvalidParseRole
			},
		},
		{
			testName: "inexistent caller user",
			setup: func(fixture *testFixture) (ChangeUserRoleInput, error) {
				newUser := fixture.createUser(user.RoleUser)
				return ChangeUserRoleInput{
					CallerID: uuid.New(),
					TargetID: newUser.ID(),
					Role:     "admin",
				}, user.ErrUserNotFound
			},
		},
		{
			testName: "inexistent target user",
			setup: func(fixture *testFixture) (ChangeUserRoleInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				return ChangeUserRoleInput{
					CallerID: owner.ID(),
					TargetID: uuid.New(),
					Role:     "admin",
				}, user.ErrUserNotFound
			},
		},
		{
			testName: "unauthorized case: admin tries to change owner role",
			setup: func(fixture *testFixture) (ChangeUserRoleInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				return ChangeUserRoleInput{
					CallerID: admin.ID(),
					TargetID: owner.ID(),
					Role:     "admin",
				}, ErrPermissionDenied
			},
		},
		{
			testName: "owner tries to change own role to admin",
			setup: func(fixture *testFixture) (ChangeUserRoleInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				return ChangeUserRoleInput{
					CallerID: owner.ID(),
					TargetID: owner.ID(),
					Role:     "admin",
				}, ErrPermissionDenied
			},
		},
		{
			testName: "fail to retrieve caller user from repository",
			setup: func(fixture *testFixture) (ChangeUserRoleInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				newUser := fixture.createUser(user.RoleUser)
				dbErr := errors.New("database unavailable")
				fixture.userRepo.FindByIDErr = dbErr
				return ChangeUserRoleInput{
					CallerID: owner.ID(),
					TargetID: newUser.ID(),
					Role:     "admin",
				}, dbErr
			},
		},
		{
			testName: "fail to retrieve target user from repository",
			setup: func(fixture *testFixture) (ChangeUserRoleInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				newUser := fixture.createUser(user.RoleUser)
				dbErr := errors.New("database unavailable")
				fixture.userRepo.FindByIDErr = dbErr
				return ChangeUserRoleInput{
					CallerID: owner.ID(),
					TargetID: newUser.ID(),
					Role:     "admin",
				}, dbErr
			},
		},
		{
			testName: "fail to save target user to repository",
			setup: func(fixture *testFixture) (ChangeUserRoleInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				newUser := fixture.createUser(user.RoleUser)
				dbErr := errors.New("database unavailable")
				fixture.userRepo.SaveErr = dbErr
				return ChangeUserRoleInput{
					CallerID: owner.ID(),
					TargetID: newUser.ID(),
					Role:     "admin",
				}, dbErr
			},
		},
		{
			testName: "idempotent case: changing user role to the same role",
			setup: func(fixture *testFixture) (ChangeUserRoleInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				newUser := fixture.createUser(user.RoleUser)
				return ChangeUserRoleInput{
					CallerID: owner.ID(),
					TargetID: newUser.ID(),
					Role:     "user",
				}, nil
			},
		},
		{
			testName: "nil caller ID",
			setup: func(fixture *testFixture) (ChangeUserRoleInput, error) {
				newUser := fixture.createUser(user.RoleUser)
				return ChangeUserRoleInput{
					CallerID: uuid.Nil(),
					TargetID: newUser.ID(),
					Role:     "admin",
				}, ErrNilID
			},
		},
		{
			testName: "nil target ID",
			setup: func(fixture *testFixture) (ChangeUserRoleInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				return ChangeUserRoleInput{
					CallerID: owner.ID(),
					TargetID: uuid.Nil(),
					Role:     "admin",
				}, ErrNilID
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			fixture := newTestFixture(t)
			input, expectedErr := test.setup(fixture)
			receivedErr := fixture.service.ChangeUserRole(input)
			if expectedErr != nil {
				assert.ErrorIs(t, receivedErr, expectedErr)
			} else {
				assert.NoError(t, receivedErr)
				target, err := fixture.userRepo.FindByID(input.TargetID)
				require.NoError(t, err)
				assert.Equal(t, input.Role, target.Role().String())
			}
		})
	}
}
