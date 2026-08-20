package user

import (
	"testing"

	"github.com/google/uuid"
)

type MockUserRepository struct {
	users map[string]*User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[string]*User),
	}
}

func (m *MockUserRepository) Save(u *User) error {
	m.users[u.id.String()] = u
	return nil
}

func (m *MockUserRepository) FindByID(id uuid.UUID) (*User, error) {
	user, exists := m.users[id.String()]
	if !exists {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func TestUserService_CreateUser(t *testing.T) {
	testCases := []struct {
		name          string
		inputName     string
		expectedError error
	}{
		{"Valid name", "Alice", nil},
		{"Empty name", "", &InvalidNameError{""}},
		{"Name too long", "ThisNameIsWayTooLongToBeValid", &InvalidNameError{"ThisNameIsWayTooLongToBeValid"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewMockUserRepository()
			service := NewUserService(repo)

			_, err := service.CreateUser(tc.inputName)
			if (err != nil && tc.expectedError == nil) || (err == nil && tc.expectedError != nil) {
				t.Fatalf("expected error: %v, got: %v", tc.expectedError, err)
			}
			if err != nil && tc.expectedError != nil && err.Error() != tc.expectedError.Error() {
				t.Fatalf("expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestUserService_GetUserByID(t *testing.T) {
	testCases := []struct {
		name               string
		usersToCreate      []string
		userNameToRetrieve string
		expectedError      error
	}{
		{"retrive existing user", []string{"Alice", "Bob", "Charlie"}, "Bob", nil},
		{"retrive non-existing user", []string{"Alice", "Bob", "Charlie"}, "Dave", ErrUserNotFound},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewMockUserRepository()
			service := NewUserService(repo)
			var userToRetrieve *User

			for _, name := range tc.usersToCreate {
				user, err := service.CreateUser(name)
				if err != nil {
					t.Fatalf("failed to create user: %v", err)
				}
				if user.name == tc.userNameToRetrieve {
					userToRetrieve = user
				}
			}

			if userToRetrieve != nil {
				retrievedUser, err := service.GetUserByID(userToRetrieve.id)
				if err != nil {
					t.Fatalf("failed to retrieve user: %v", err)
				}
				if retrievedUser.id != userToRetrieve.id {
					t.Fatalf("expected user ID: %v, got: %v", userToRetrieve.id, retrievedUser.id)
				}
			} else {
				id := uuid.New()
				_, err := service.GetUserByID(id)
				if err == nil || err != tc.expectedError {
					t.Fatalf("expected error: %v, got: %v", tc.expectedError, err)
				}
			}
		})
	}
}
