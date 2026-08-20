package user

import (
	"testing"
)

func TestValidateName(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectErr bool
	}{
		{"valid name", "John Doe", false},
		{"empty name", "", true},
		{"name too long", "ThisNameIsWayTooLongToBeValid", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateName(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("validateName() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}

func TestSetRole(t *testing.T) {
	tests := []struct {
		name           string
		actor          Role
		target         Role
		newRole        Role
		wantErr        bool
		wantTargetRole Role
	}{
		{"user cannot change role", RoleUser, RoleUser, RoleAdmin, true, RoleUser},
		{"admin can change user to admin", RoleAdmin, RoleUser, RoleAdmin, false, RoleAdmin},
		{"admin cannot change owner", RoleAdmin, RoleOwner, RoleUser, true, RoleOwner},
		{"owner can promote to owner", RoleOwner, RoleAdmin, RoleOwner, false, RoleOwner},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actor, _ := NewUser("actor")
			actor.role = tc.actor
			target, _ := NewUser("target")
			target.role = tc.target

			err := actor.SetRole(target, tc.newRole)
			if (err != nil) != tc.wantErr {
				t.Fatalf("err = %v, wantErr = %v", err, tc.wantErr)
			}
			if target.role != tc.wantTargetRole {
				t.Fatalf("target role = %v, want %v", target.role, tc.wantTargetRole)
			}
		})
	}
}
