package user

import (
	"testing"

	"uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserCreation(t *testing.T) {

	t.Run("create user with valid inputs", func(t *testing.T) {
		name, err := NewName("John Doe")
		require.NoError(t, err)
		email, err := NewEmail("john.doe@example.com")
		require.NoError(t, err)
		user, err := NewUser(name, email)

		assert.NotNil(t, user)
		assert.NoError(t, err)
		assert.Equal(t, name, user.Name())
		assert.Equal(t, email, user.Email())
		assert.Equal(t, RoleUser, user.Role())
		assert.NotEqual(t, uuid.Nil(), user.ID())
	})

	t.Run("create user with invalid name", func(t *testing.T) {
		email, err := NewEmail("john.doe@example.com")
		require.NoError(t, err)
		user, err := NewUser(Name{}, email)

		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyName, err)
	})

	t.Run("create user with invalid email", func(t *testing.T) {
		name, err := NewName("John Doe")
		require.NoError(t, err)
		user, err := NewUser(name, Email{})

		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyEmail, err)
	})
}

func TestRestoreUser(t *testing.T) {
	t.Run("restore user with valid inputs", func(t *testing.T) {
		id := uuid.New()
		name, err := NewName("John Doe")
		require.NoError(t, err)
		email, err := NewEmail("john.doe@example.com")
		require.NoError(t, err)
		role := RoleAdmin
		user, err := RestoreUser(id, name, email, role)

		assert.NotNil(t, user)
		assert.NoError(t, err)
		assert.Equal(t, name, user.Name())
		assert.Equal(t, email, user.Email())
		assert.Equal(t, RoleAdmin, user.Role())
		assert.Equal(t, id, user.ID())
	})

	t.Run("restore user with invalid id", func(t *testing.T) {
		name, err := NewName("John Doe")
		require.NoError(t, err)
		email, err := NewEmail("john.doe@example.com")
		require.NoError(t, err)
		role := RoleAdmin
		user, err := RestoreUser(uuid.Nil(), name, email, role)

		assert.Nil(t, user)
		assert.Error(t, err)
	})

	t.Run("restore user with invalid name", func(t *testing.T) {
		id := uuid.New()
		email, err := NewEmail("john.doe@example.com")
		require.NoError(t, err)
		role := RoleAdmin
		user, err := RestoreUser(id, Name{}, email, role)

		assert.Nil(t, user)
		assert.Error(t, err)
	})

	t.Run("restore user with invalid email", func(t *testing.T) {
		id := uuid.New()
		name, err := NewName("John Doe")
		require.NoError(t, err)
		role := RoleAdmin
		user, err := RestoreUser(id, name, Email{}, role)

		assert.Nil(t, user)
		assert.Error(t, err)
	})

	t.Run("restore user with invalid role", func(t *testing.T) {
		id := uuid.New()
		name, err := NewName("John Doe")
		require.NoError(t, err)
		email, err := NewEmail("john.doe@example.com")
		require.NoError(t, err)
		user, err := RestoreUser(id, name, email, Role("invalid_role"))

		assert.Nil(t, user)
		assert.Error(t, err)
	})
}

func TestSetRole(t *testing.T) {

	t.Run("valid role cases", func(t *testing.T) {
		validRoles := []struct {
			testName string
			role     Role
		}{
			{"owner role", RoleOwner},
			{"admin role", RoleAdmin},
			{"user role", RoleUser},
			{"guest role", RoleGuest},
		}

		for _, testCase := range validRoles {
			t.Run(testCase.testName, func(t *testing.T) {
				name, err := NewName("John Doe")
				require.NoError(t, err)
				email, err := NewEmail("john.doe@example.com")
				require.NoError(t, err)
				user, err := NewUser(name, email)
				require.NoError(t, err)

				err = user.SetRole(testCase.role)
				require.NoError(t, err)
				assert.Equal(t, testCase.role, user.Role())
			})
		}
	})

	t.Run("invalid role cases", func(t *testing.T) {
		name, err := NewName("John Doe")
		require.NoError(t, err)
		email, err := NewEmail("john.doe@example.com")
		require.NoError(t, err)
		user, err := NewUser(name, email)

		require.NoError(t, err)
		err = user.SetRole(Role("invalid_role"))
		require.Error(t, err)
		assert.Equal(t, ErrInvalidRole, err)
	})
}

func TestSetName(t *testing.T) {
	t.Run("valid name change", func(t *testing.T) {
		name, err := NewName("John Doe")
		require.NoError(t, err)
		email, err := NewEmail("john.doe@example.com")
		require.NoError(t, err)
		user, err := NewUser(name, email)
		require.NoError(t, err)

		newName, err := NewName("Jane Doe")
		require.NoError(t, err)

		err = user.SetName(newName)
		require.NoError(t, err)
		assert.Equal(t, newName, user.Name())
	})

	t.Run("invalid name change", func(t *testing.T) {
		name, err := NewName("John Doe")
		require.NoError(t, err)
		email, err := NewEmail("john.doe@example.com")
		require.NoError(t, err)
		user, err := NewUser(name, email)
		require.NoError(t, err)

		invalidName := Name{} // Invalid name

		err = user.SetName(invalidName)
		require.Error(t, err)
		assert.Equal(t, ErrEmptyName, err)
	})
}

func TestSetEmail(t *testing.T) {
	t.Run("valid email change", func(t *testing.T) {
		name, err := NewName("John Doe")
		require.NoError(t, err)
		email, err := NewEmail("john.doe@example.com")
		require.NoError(t, err)
		user, err := NewUser(name, email)
		require.NoError(t, err)

		newEmail, err := NewEmail("jane.doe@example.com")
		require.NoError(t, err)

		err = user.SetEmail(newEmail)
		require.NoError(t, err)
		assert.Equal(t, newEmail, user.Email())
	})

	t.Run("invalid email change", func(t *testing.T) {
		name, err := NewName("John Doe")
		require.NoError(t, err)
		email, err := NewEmail("john.doe@example.com")
		require.NoError(t, err)
		user, err := NewUser(name, email)
		require.NoError(t, err)

		invalidEmail := Email{}

		err = user.SetEmail(invalidEmail)
		require.Error(t, err)
		assert.Equal(t, ErrEmptyEmail, err)
	})
}

func TestIsAdmin(t *testing.T) {
	t.Run("user is admin", func(t *testing.T) {
		name, err := NewName("Admin User")
		require.NoError(t, err)
		email, err := NewEmail("admin@example.com")
		require.NoError(t, err)
		user, err := NewUser(name, email)
		require.NoError(t, err)
		err = user.SetRole(RoleAdmin)
		require.NoError(t, err)
		assert.True(t, user.IsAdmin())
	})

	t.Run("user is not admin", func(t *testing.T) {
		name, err := NewName("Regular User")
		require.NoError(t, err)
		email, err := NewEmail("user@example.com")
		require.NoError(t, err)
		user, err := NewUser(name, email)
		require.NoError(t, err)
		assert.False(t, user.IsAdmin())
	})
}

func TestIsOwner(t *testing.T) {
	t.Run("user is owner", func(t *testing.T) {
		name, err := NewName("Owner User")
		require.NoError(t, err)
		email, err := NewEmail("owner@example.com")
		require.NoError(t, err)
		user, err := NewUser(name, email)
		require.NoError(t, err)
		err = user.SetRole(RoleOwner)
		require.NoError(t, err)
		assert.True(t, user.IsOwner())
	})

	t.Run("user is not owner", func(t *testing.T) {
		name, err := NewName("Regular User")
		require.NoError(t, err)
		email, err := NewEmail("user@example.com")
		require.NoError(t, err)
		user, err := NewUser(name, email)
		require.NoError(t, err)
		assert.False(t, user.IsOwner())
	})
}
