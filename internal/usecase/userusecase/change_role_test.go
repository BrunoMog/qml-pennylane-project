package userusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestChangeUserRole(t *testing.T) {
	tests := []struct {
		testName      string
		setup         func(fixture *testFixture) ChangeUserRoleInput
		expectedError error
	}{
		{
			testName: "valid case: owner changes admin to user",
			setup: func(fixture *testFixture) ChangeUserRoleInput {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				return ChangeUserRoleInput{
					CallerID: owner.ID(),
					TargetID: admin.ID(),
					Role:     user.RoleUser,
				}
			},
			expectedError: nil,
		},
		{
			testName: "inexistent caller user",
			setup: func(fixture *testFixture) ChangeUserRoleInput {
				newUser := fixture.createUser(user.RoleUser)
				return ChangeUserRoleInput{
					CallerID: uuid.New(),
					TargetID: newUser.ID(),
					Role:     user.RoleAdmin,
				}
			},
			expectedError: &testkit.ErrUserNotFound{},
		},
		{
			testName: "inexistent target user",
			setup: func(fixture *testFixture) ChangeUserRoleInput {
				owner := fixture.createUser(user.RoleOwner)
				return ChangeUserRoleInput{
					CallerID: owner.ID(),
					TargetID: uuid.New(),
					Role:     user.RoleAdmin,
				}
			},
			expectedError: &testkit.ErrUserNotFound{},
		},
		{
			testName: "unauthorized case: admin tries to change owner role",
			setup: func(fixture *testFixture) ChangeUserRoleInput {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				return ChangeUserRoleInput{
					CallerID: admin.ID(),
					TargetID: owner.ID(),
					Role:     user.RoleAdmin,
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "owner tries to change own role to admin",
			setup: func(fixture *testFixture) ChangeUserRoleInput {
				owner := fixture.createUser(user.RoleOwner)
				return ChangeUserRoleInput{
					CallerID: owner.ID(),
					TargetID: owner.ID(),
					Role:     user.RoleAdmin,
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "try to assign invalid role",
			setup: func(fixture *testFixture) ChangeUserRoleInput {
				owner := fixture.createUser(user.RoleOwner)
				user := fixture.createUser(user.RoleUser)
				return ChangeUserRoleInput{
					CallerID: owner.ID(),
					TargetID: user.ID(),
					Role:     "invalid_role",
				}
			},
			expectedError: &user.InvalidRoleError{},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			fixture := newTestFixture(t)
			input := test.setup(fixture)
			err := fixture.service.ChangeUserRole(input)
			if test.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, test.expectedError, err)
			} else {
				assert.NoError(t, err)
				target, err := fixture.userRepo.FindByID(input.TargetID)
				assert.NoError(t, err)
				assert.Equal(t, input.Role, target.Role())
			}
		})
	}
}
