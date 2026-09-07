package userusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		expectedError error
		setup         func(fixture *testFixture) CreateUserInput
		testName      string
	}{
		{
			testName: "create user successfully",
			setup: func(f *testFixture) CreateUserInput {
				return CreateUserInput{
					Name:  "John Doe",
					Email: "john.doe@example.com",
				}
			},
			expectedError: nil,
		},
		{
			testName: "create user with existing email",
			setup: func(f *testFixture) CreateUserInput {
				userEmail, err := user.NewEmail("jane.doe@example.com")
				require.NoError(t, err)
				u := f.createUser(user.RoleUser)
				u.SetEmail(userEmail)

				return CreateUserInput{
					Name:  "Jhon Doe",
					Email: userEmail.Value(),
				}
			},
			expectedError: &EmailAlreadyExistsError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			fixture := newTestFixture(t)
			input := tt.setup(fixture)
			user, err := fixture.service.CreateUser(input)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, input.Name, user.Name)
				assert.Equal(t, input.Email, user.Email)
				assert.Equal(t, user.Role, user.Role)
			}
		})
	}
}
