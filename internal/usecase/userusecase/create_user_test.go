package userusecase

import (
	"errors"
	"pennylane_project_backend/internal/domain/user"
	"testing"
	"uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {

	tests := []struct {
		setup    func(fixture *testFixture) (CreateUserInput, error)
		testName string
	}{
		{
			testName: "create user successfully",
			setup: func(f *testFixture) (CreateUserInput, error) {
				return CreateUserInput{
					Name:  "John Doe",
					Email: "john.doe@example.com",
				}, nil
			},
		},
		{
			testName: "create user with invalid name",
			setup: func(f *testFixture) (CreateUserInput, error) {
				return CreateUserInput{
					Name:  "",
					Email: "john.doe@example.com",
				}, user.ErrInvalidName
			},
		},
		{
			testName: "create user with invalid email",
			setup: func(f *testFixture) (CreateUserInput, error) {
				return CreateUserInput{
					Name:  "John Doe",
					Email: "invalid-email",
				}, user.ErrInvalidEmail
			},
		},
		{
			testName: "create user with existing email",
			setup: func(f *testFixture) (CreateUserInput, error) {
				userEmail, err := user.NewEmail("jane.doe@example.com")
				require.NoError(t, err)
				u := f.createUser(user.RoleUser)
				u.SetEmail(userEmail)

				return CreateUserInput{
					Name:  "Jhon Doe",
					Email: userEmail.String(),
				}, ErrEmailAlreadyExists
			},
		},
		{
			testName: "error when checking email existence",
			setup: func(f *testFixture) (CreateUserInput, error) {
				dbErr := errors.New("database unavailable")
				f.userRepo.ExistsByEmailErr = dbErr

				return CreateUserInput{
					Name:  "John Doe",
					Email: "john.doe@example.com",
				}, dbErr
			},
		},
		{
			testName: "error when saving user",
			setup: func(f *testFixture) (CreateUserInput, error) {
				dbErr := errors.New("database unavailable")
				f.userRepo.SaveErr = dbErr

				return CreateUserInput{
					Name:  "John Doe",
					Email: "john.doe@example.com",
				}, dbErr
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			fixture := newTestFixture(t)
			input, expectedErr := tt.setup(fixture)
			newUser, receivedErr := fixture.service.CreateUser(input)

			if expectedErr != nil {
				assert.ErrorIs(t, receivedErr, expectedErr)
			} else {
				assert.NoError(t, receivedErr)
				assert.NotNil(t, newUser)
				assert.Equal(t, input.Name, newUser.Name)
				assert.Equal(t, input.Email, newUser.Email)
				assert.Equal(t, newUser.Role, newUser.Role)
				assert.NotEqual(t, uuid.Nil(), newUser.ID)
			}
		})
	}
}
