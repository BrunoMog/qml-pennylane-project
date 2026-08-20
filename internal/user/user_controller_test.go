package user

import (
	"errors"
	"testing"

	"github.com/google/uuid"
)

type fakeUserUseCase struct {
	createUserCalled  bool
	getUserByIDCalled bool
	setRoleCalled     bool

	createdName string
	userID      uuid.UUID
	targetUser  *User
	newRole     Role

	createUserResult *User
	createUserErr    error
	foundUser        *User
	getUserErr       error
	setRoleErr       error
}

func (f *fakeUserUseCase) CreateUser(name string) (*User, error) {
	f.createUserCalled = true
	f.createdName = name
	return f.createUserResult, f.createUserErr
}

func (f *fakeUserUseCase) GetUserByID(id uuid.UUID) (*User, error) {
	f.getUserByIDCalled = true
	f.userID = id
	return f.foundUser, f.getUserErr
}

func (f *fakeUserUseCase) SetUserRole(user *User, newRole Role) error {
	f.setRoleCalled = true
	f.targetUser = user
	f.newRole = newRole
	return f.setRoleErr
}

func TestUserController_CreateUser(t *testing.T) {
	expectedErr := errors.New("create failed")

	fake := &fakeUserUseCase{
		createUserErr: expectedErr,
	}
	controller := NewUserController(fake)

	err := controller.CreateUser("Alice")

	if !errors.Is(err, expectedErr) {
		t.Fatalf("error = %v, want %v", err, expectedErr)
	}

	if !fake.createUserCalled {
		t.Fatal("expected CreateUser to be called")
	}

	if fake.createdName != "Alice" {
		t.Fatalf("name = %q, want %q", fake.createdName, "Alice")
	}
}

func TestUserController_CreateUser_Success(t *testing.T) {
	fake := &fakeUserUseCase{}
	controller := NewUserController(fake)

	err := controller.CreateUser("Alice")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !fake.createUserCalled {
		t.Fatal("expected CreateUser to be called")
	}
}
