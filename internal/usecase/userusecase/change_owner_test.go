package userusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestChangeOwner(t *testing.T) {
	tests := []struct {
		testName      string
		setup         func(fixture *testFixture) ChangeOwnerInput
		expectedError error
	}{
		{
			testName: "valid case: owner swaps ownership with admin",
			setup: func(fixture *testFixture) ChangeOwnerInput {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				return ChangeOwnerInput{
					CallerID: owner.ID(),
					TargetID: admin.ID(),
				}
			},
			expectedError: nil,
		},
		{
			testName: "inexistent caller user",
			setup: func(fixture *testFixture) ChangeOwnerInput {
				newUser := fixture.createUser(user.RoleUser)
				return ChangeOwnerInput{
					CallerID: uuid.New(),
					TargetID: newUser.ID(),
				}
			},
			expectedError: &testkit.ErrUserNotFound{},
		},
		{
			testName: "inexistent target user",
			setup: func(fixture *testFixture) ChangeOwnerInput {
				owner := fixture.createUser(user.RoleOwner)
				return ChangeOwnerInput{
					CallerID: owner.ID(),
					TargetID: uuid.New(),
				}
			},
			expectedError: &testkit.ErrUserNotFound{},
		},
		{
			testName: "admin tries to change owner role",
			setup: func(fixture *testFixture) ChangeOwnerInput {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				return ChangeOwnerInput{
					CallerID: admin.ID(),
					TargetID: owner.ID(),
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "owner tries to swap ownership with self",
			setup: func(fixture *testFixture) ChangeOwnerInput {
				owner := fixture.createUser(user.RoleOwner)
				return ChangeOwnerInput{
					CallerID: owner.ID(),
					TargetID: owner.ID(),
				}
			},
			expectedError: &UnauthorizedError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			fixture := newTestFixture(t)
			input := tt.setup(fixture)
			err := fixture.service.ChangeOwner(input)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				target, err := fixture.userRepo.FindByID(input.TargetID)
				assert.NoError(t, err)
				assert.Equal(t, user.RoleOwner, target.Role())
			}
		})
	}
}
