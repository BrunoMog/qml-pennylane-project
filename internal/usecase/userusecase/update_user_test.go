package userusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		testName      string
		setup         func(fixture *testFixture) UpdateUserInput
		expectedError error
	}{
		{
			testName: "owner updates admin's name and email",
			setup: func(fixture *testFixture) UpdateUserInput {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				newName := "Updated Admin Name"
				newEmail := "updated@example.com"
				return UpdateUserInput{
					CallerID: owner.ID(),
					TargetID: admin.ID(),
					Name:     &newName,
					Email:    &newEmail,
				}
			},
			expectedError: nil,
		},
		{
			testName: "user updates self email",
			setup: func(fixture *testFixture) UpdateUserInput {
				user := fixture.createUser(user.RoleUser)
				newEmail := "updated@example.com"
				return UpdateUserInput{
					CallerID: user.ID(),
					TargetID: user.ID(),
					Email:    &newEmail,
				}
			},
			expectedError: nil,
		},
		{
			testName: "inexistent target user",
			setup: func(fixture *testFixture) UpdateUserInput {
				owner := fixture.createUser(user.RoleOwner)
				newName := "Nonexistent User"
				return UpdateUserInput{
					CallerID: owner.ID(),
					TargetID: uuid.New(),
					Name:     &newName,
				}
			},
			expectedError: &testkit.ErrUserNotFound{},
		},
		{
			testName: "inexistent caller user",
			setup: func(fixture *testFixture) UpdateUserInput {
				user := fixture.createUser(user.RoleUser)
				newEmail := "updated@example.com"
				return UpdateUserInput{
					CallerID: uuid.New(),
					TargetID: user.ID(),
					Email:    &newEmail,
				}
			},
			expectedError: &testkit.ErrUserNotFound{},
		},
		{
			testName: "no fields to update",
			setup: func(fixture *testFixture) UpdateUserInput {
				owner := fixture.createUser(user.RoleOwner)
				return UpdateUserInput{
					CallerID: owner.ID(),
					TargetID: owner.ID(),
				}
			},
			expectedError: &NoFieldsToUpdateError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			fixture := newTestFixture(t)
			input := tt.setup(fixture)
			err := fixture.service.UpdateUser(input)
			if tt.expectedError != nil {
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				updatedUser, err := fixture.userRepo.FindByID(input.TargetID)
				assert.NoError(t, err)
				if input.Name != nil {
					assert.Equal(t, *input.Name, updatedUser.Name())
				}
				if input.Email != nil {
					assert.Equal(t, *input.Email, updatedUser.Email())
				}
			}
		})
	}
}
