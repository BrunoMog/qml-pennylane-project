package userusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"pennylane_project_backend/internal/testkit"
	"testing"

	"uuid"

	"github.com/stretchr/testify/assert"
)

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		expectedError error
		setup         func(fixture *testFixture) UpdateUserInput
		testName      string
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
		{
			testName: "try to update email with existing email",
			setup: func(fixture *testFixture) UpdateUserInput {
				owner := fixture.createUser(user.RoleOwner)
				existingUser := fixture.createUser(user.RoleUser)
				existingEmail := existingUser.Email().Value()
				return UpdateUserInput{
					CallerID: owner.ID(),
					TargetID: owner.ID(),
					Email:    &existingEmail,
				}
			},
			expectedError: &EmailAlreadyExistsError{},
		},
		{
			testName: "unauthorized user tries to update another user",
			setup: func(fixture *testFixture) UpdateUserInput {
				user1 := fixture.createUser(user.RoleUser)
				user2 := fixture.createUser(user.RoleUser)
				newName := "Updated Name"
				return UpdateUserInput{
					CallerID: user1.ID(),
					TargetID: user2.ID(),
					Name:     &newName,
				}
			},
			expectedError: &UnauthorizedError{},
		},
		{
			testName: "idepotent update (same name and email)",
			setup: func(fixture *testFixture) UpdateUserInput {
				user := fixture.createUser(user.RoleUser)
				name := user.Name().Value()
				email := user.Email().Value()
				return UpdateUserInput{
					CallerID: user.ID(),
					TargetID: user.ID(),
					Name:     &name,
					Email:    &email,
				}
			},
			expectedError: nil,
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
					assert.Equal(t, *input.Name, updatedUser.Name().Value())
				}
				if input.Email != nil {
					assert.Equal(t, *input.Email, updatedUser.Email().Value())
				}
			}
		})
	}
}
