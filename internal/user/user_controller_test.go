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

func TestUserController_GetUserByID(t *testing.T) {
	id := uuid.New()
	expectedUser, _ := NewUser("Alice")

	fake := &fakeUserUseCase{
		foundUser: expectedUser,
	}
	controller := NewUserController(fake)

	user, err := controller.GetUserByID(id)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user != expectedUser {
		t.Fatalf("user = %p, want %p", user, expectedUser)
	}

	if !fake.getUserByIDCalled {
		t.Fatal("expected GetUserByID to be called")
	}

	if fake.userID != id {
		t.Fatalf("id = %v, want %v", fake.userID, id)
	}
}

func TestUserController_SetUserRole(t *testing.T) {
	targetUser, _ := NewUser("Alice")
	expectedErr := errors.New("set role failed")

	fake := &fakeUserUseCase{
		setRoleErr: expectedErr,
	}
	controller := NewUserController(fake)

	err := controller.SetUserRole(targetUser, RoleAdmin)

	if !errors.Is(err, expectedErr) {
		t.Fatalf("error = %v, want %v", err, expectedErr)
	}

	if !fake.setRoleCalled {
		t.Fatal("expected SetUserRole to be called")
	}

	if fake.targetUser != targetUser {
		t.Fatal("controller passed a different user")
	}

	if fake.newRole != RoleAdmin {
		t.Fatalf("role = %v, want %v", fake.newRole, RoleAdmin)
	}
}
