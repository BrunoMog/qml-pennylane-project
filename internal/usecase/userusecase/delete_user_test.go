package userusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		expectedError error
		setup         func(fixture *testFixture) DeleteUserInput
		testName      string
	}{
		{
			testName: "owner deletes admin",
			setup: func(fixture *testFixture) DeleteUserInput {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				return DeleteUserInput{
					CallerID: owner.ID(),
					TargetID: admin.ID(),
				}
			},
			expectedError: nil,
		},
		{
			testName: "inexistent target user",
			setup: func(fixture *testFixture) DeleteUserInput {
				owner := fixture.createUser(user.RoleOwner)
				return DeleteUserInput{
					CallerID: owner.ID(),
					TargetID: uuid.New(),
				}
			},
			expectedError: &testkit.ErrUserNotFound{},
		},
		{
			testName: "inexistent caller user",
			setup: func(fixture *testFixture) DeleteUserInput {
				user := fixture.createUser(user.RoleUser)
				return DeleteUserInput{
					CallerID: uuid.New(),
					TargetID: user.ID(),
				}
			},
			expectedError: &testkit.ErrUserNotFound{},
		},
		{
			testName: "unauthorized case: admin tries to delete owner",
			setup: func(fixture *testFixture) DeleteUserInput {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				return DeleteUserInput{
					CallerID: admin.ID(),
					TargetID: owner.ID(),
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "owner deletes self",
			setup: func(fixture *testFixture) DeleteUserInput {
				owner := fixture.createUser(user.RoleOwner)
				return DeleteUserInput{
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

			err := fixture.service.DeleteUser(input)
			if tt.expectedError != nil {
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				exists, err := fixture.userRepo.ExistsByID(input.TargetID)
				assert.NoError(t, err)
				assert.False(t, exists)
			}
		})
	}

}
