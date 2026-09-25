package userusecase

import (
	"errors"
	"pennylane_project_backend/internal/domain/user"
	"testing"

	"uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChangeOwner(t *testing.T) {
	tests := []struct {
		setup    func(fixture *testFixture) (ChangeOwnerInput, error)
		testName string
	}{
		{
			testName: "owner swaps ownership with admin",
			setup: func(fixture *testFixture) (ChangeOwnerInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				return ChangeOwnerInput{
					CallerID: owner.ID(),
					TargetID: admin.ID(),
				}, nil
			},
		},
		{
			testName: "inexistent caller user",
			setup: func(fixture *testFixture) (ChangeOwnerInput, error) {
				newUser := fixture.createUser(user.RoleUser)
				return ChangeOwnerInput{
					CallerID: uuid.New(),
					TargetID: newUser.ID(),
				}, user.ErrUserNotFound
			},
		},
		{
			testName: "inexistent target user",
			setup: func(fixture *testFixture) (ChangeOwnerInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				return ChangeOwnerInput{
					CallerID: owner.ID(),
					TargetID: uuid.New(),
				}, user.ErrUserNotFound
			},
		},
		{
			testName: "admin tries to change owner role",
			setup: func(fixture *testFixture) (ChangeOwnerInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				return ChangeOwnerInput{
					CallerID: admin.ID(),
					TargetID: owner.ID(),
				}, ErrPermissionDenied
			},
		},
		{
			testName: "fail to retrieve caller user",
			setup: func(fixture *testFixture) (ChangeOwnerInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				dbErr := errors.New("database unavailable")
				fixture.userRepo.FindByIDErr = dbErr
				return ChangeOwnerInput{
					CallerID: owner.ID(),
					TargetID: admin.ID(),
				}, dbErr
			},
		},
		{
			testName: "fail to retrieve target user",
			setup: func(fixture *testFixture) (ChangeOwnerInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				dbErr := errors.New("database unavailable")
				fixture.userRepo.FindByIDErr = dbErr
				return ChangeOwnerInput{
					CallerID: owner.ID(),
					TargetID: admin.ID(),
				}, dbErr
			},
		},
		{
			testName: "fail to change owner in repository",
			setup: func(fixture *testFixture) (ChangeOwnerInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				dbErr := errors.New("database unavailable")
				fixture.userRepo.ChangeOwnerErr = dbErr
				return ChangeOwnerInput{
					CallerID: owner.ID(),
					TargetID: admin.ID(),
				}, dbErr
			},
		},
		{
			testName: "idempotent change owner (caller is target)",
			setup: func(fixture *testFixture) (ChangeOwnerInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				return ChangeOwnerInput{
					CallerID: owner.ID(),
					TargetID: owner.ID(),
				}, nil
			},
		},
		{
			testName: "nil caller ID",
			setup: func(fixture *testFixture) (ChangeOwnerInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				return ChangeOwnerInput{
					CallerID: uuid.Nil(),
					TargetID: owner.ID(),
				}, ErrNilID
			},
		},
		{
			testName: "nil target ID",
			setup: func(fixture *testFixture) (ChangeOwnerInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				return ChangeOwnerInput{
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
			receivedErr := fixture.service.ChangeOwner(input)
			if expectedErr != nil {
				assert.ErrorIs(t, receivedErr, expectedErr)
			} else {
				assert.NoError(t, receivedErr)
				target, err := fixture.userRepo.FindByID(input.TargetID)
				require.NoError(t, err)
				assert.Equal(t, user.RoleOwner, target.Role())
			}
		})
	}
}
