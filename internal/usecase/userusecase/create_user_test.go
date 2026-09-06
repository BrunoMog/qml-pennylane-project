package userusecase

import (
	"pennylane_project_backend/internal/domain/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (f *testFixture) setupEmailAlreadyExists(t *testing.T, email string) {
	user, err := user.NewUser("Existing User", email)
	require.NoError(t, err)
	f.userRepo.Save(user)
}
func TestCreateUser(t *testing.T) {
	tests := []struct {
		expectedError error
		setup         func(fixture *testFixture)
		testName      string
		userName      string
		userEmail     string
	}{
		{
			testName:      "create user successfully",
			userName:      "John Doe",
			userEmail:     "john.doe@example.com",
			expectedError: nil,
		},
		{
			testName:      "create user with empty name",
			userName:      "",
			userEmail:     "john.doe@example.com",
			expectedError: &user.InvalidNameError{},
		},
		{
			testName:      "create user with empty email",
			userName:      "John Doe",
			userEmail:     "",
			expectedError: &user.InvalidEmailError{},
		},
		{
			testName:      "email already exists",
			userName:      "Jane Doe",
			userEmail:     "jane.doe@example.com",
			expectedError: &EmailAlreadyExistsError{},
			setup: func(fixture *testFixture) {
				fixture.setupEmailAlreadyExists(t, "jane.doe@example.com")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			fixture := newTestFixture(t)
			if tt.setup != nil {
				tt.setup(fixture)

			}
			user, err := fixture.service.CreateUser(tt.userName, tt.userEmail)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.userName, user.Name)
				assert.Equal(t, tt.userEmail, user.Email)
			}
		})
	}
}
