package userusecase

import (
	"errors"
	"pennylane_project_backend/internal/domain/user"
	"testing"

	"uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		setup    func(fixture *testFixture) (UpdateUserInput, error)
		testName string
	}{
		{
			testName: "owner updates admin's name and email",
			setup: func(fixture *testFixture) (UpdateUserInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				admin := fixture.createUser(user.RoleAdmin)
				newName := "Updated Admin Name"
				newEmail := "updated@example.com"
				return UpdateUserInput{
					CallerID: owner.ID(),
					TargetID: admin.ID(),
					Name:     &newName,
					Email:    &newEmail,
				}, nil
			},
		},
		{
			testName: "user updates self email",
			setup: func(fixture *testFixture) (UpdateUserInput, error) {
				user := fixture.createUser(user.RoleUser)
				newEmail := "updated@example.com"
				return UpdateUserInput{
					CallerID: user.ID(),
					TargetID: user.ID(),
					Email:    &newEmail,
				}, nil
			},
		},
		{
			testName: "inexistent target user",
			setup: func(fixture *testFixture) (UpdateUserInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				newName := "Nonexistent User"
				return UpdateUserInput{
					CallerID: owner.ID(),
					TargetID: uuid.New(),
					Name:     &newName,
				}, user.ErrUserNotFound
			},
		},
		{
			testName: "inexistent caller user",
			setup: func(fixture *testFixture) (UpdateUserInput, error) {
				newUser := fixture.createUser(user.RoleUser)
				newEmail := "updated@example.com"
				return UpdateUserInput{
					CallerID: uuid.New(),
					TargetID: newUser.ID(),
					Email:    &newEmail,
				}, user.ErrUserNotFound
			},
		},
		{
			testName: "no fields to update",
			setup: func(fixture *testFixture) (UpdateUserInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				return UpdateUserInput{
					CallerID: owner.ID(),
					TargetID: owner.ID(),
				}, ErrNoFieldsToUpdate
			},
		},
		{
			testName: "invalid name format",
			setup: func(fixture *testFixture) (UpdateUserInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				invalidName := "Invalid@Name"
				return UpdateUserInput{
					CallerID: owner.ID(),
					TargetID: owner.ID(),
					Name:     &invalidName,
				}, user.ErrInvalidName
			},
		},
		{
			testName: "invalid email format",
			setup: func(fixture *testFixture) (UpdateUserInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				invalidEmail := "invalid-email"
				return UpdateUserInput{
					CallerID: owner.ID(),
					TargetID: owner.ID(),
					Email:    &invalidEmail,
				}, user.ErrInvalidEmail
			},
		},
		{
			testName: "try to update email with existing email",
			setup: func(fixture *testFixture) (UpdateUserInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				existingUser := fixture.createUser(user.RoleUser)
				existingEmail := existingUser.Email().String()
				return UpdateUserInput{
					CallerID: owner.ID(),
					TargetID: owner.ID(),
					Email:    &existingEmail,
				}, ErrEmailAlreadyExists
			},
		},
		{
			testName: "unauthorized user tries to update another user",
			setup: func(fixture *testFixture) (UpdateUserInput, error) {
				user1 := fixture.createUser(user.RoleUser)
				user2 := fixture.createUser(user.RoleUser)
				newName := "Updated Name"
				return UpdateUserInput{
					CallerID: user1.ID(),
					TargetID: user2.ID(),
					Name:     &newName,
				}, ErrPermissionDenied
			},
		},
		{
			testName: "idepotent update (same name and email)",
			setup: func(fixture *testFixture) (UpdateUserInput, error) {
				user := fixture.createUser(user.RoleUser)
				name := user.Name().String()
				email := user.Email().String()
				return UpdateUserInput{
					CallerID: user.ID(),
					TargetID: user.ID(),
					Name:     &name,
					Email:    &email,
				}, nil
			},
		},
		{
			testName: "fail to retrieve caller user",
			setup: func(fixture *testFixture) (UpdateUserInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				newUser := fixture.createUser(user.RoleUser)
				dbErr := errors.New("database unavailable")
				fixture.userRepo.FindByIDErr = dbErr
				newName := "Updated Name"
				return UpdateUserInput{
					CallerID: owner.ID(),
					TargetID: newUser.ID(),
					Name:     &newName,
				}, dbErr
			},
		},
		{
			testName: "fail to retrieve target user",
			setup: func(fixture *testFixture) (UpdateUserInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				newUser := fixture.createUser(user.RoleUser)
				dbErr := errors.New("database unavailable")
				fixture.userRepo.FindByIDErr = dbErr
				newEmail := "updated@example.com"
				return UpdateUserInput{
					CallerID: owner.ID(),
					TargetID: newUser.ID(),
					Email:    &newEmail,
				}, dbErr
			},
		},
		{
			testName: "fail to save user",
			setup: func(fixture *testFixture) (UpdateUserInput, error) {
				owner := fixture.createUser(user.RoleOwner)
				newUser := fixture.createUser(user.RoleUser)
				dbErr := errors.New("database unavailable")
				fixture.userRepo.SaveErr = dbErr
				newName := "Updated Name"
				return UpdateUserInput{
					CallerID: owner.ID(),
					TargetID: newUser.ID(),
					Name:     &newName,
				}, dbErr
			},
		},
		{
			testName: "nil caller ID",
			setup: func(fixture *testFixture) (UpdateUserInput, error) {
				newUser := fixture.createUser(user.RoleUser)
				newName := "Updated Name"
				return UpdateUserInput{
					CallerID: uuid.Nil(),
					TargetID: newUser.ID(),
					Name:     &newName,
				}, ErrNilID
			},
		},
		{
			testName: "nil target ID",
			setup: func(fixture *testFixture) (UpdateUserInput, error) {
				newUser := fixture.createUser(user.RoleUser)
				newEmail := "updated@example.com"
				return UpdateUserInput{
					CallerID: newUser.ID(),
					TargetID: uuid.Nil(),
					Email:    &newEmail,
				}, ErrNilID
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			fixture := newTestFixture(t)
			input, expectedErr := tt.setup(fixture)
			receivedErr := fixture.service.UpdateUser(input)
			if expectedErr != nil {
				assert.ErrorIs(t, receivedErr, expectedErr)
			} else {
				assert.NoError(t, receivedErr)
				updatedUser, err := fixture.userRepo.FindByID(input.TargetID)
				require.NoError(t, err)
				if input.Name != nil {
					assert.Equal(t, *input.Name, updatedUser.Name().String())
				}
				if input.Email != nil {
					assert.Equal(t, *input.Email, updatedUser.Email().String())
				}
			}
		})
	}
}
