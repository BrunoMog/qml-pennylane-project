package user

import "strings"

type Role string

const (
	RoleOwner Role = "owner"
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
	RoleGuest Role = "guest"
)

func ParseRole(roleStr string) (Role, error) {
	switch strings.ToLower(strings.TrimSpace(roleStr)) {
	case "owner":
		return RoleOwner, nil
	case "admin":
		return RoleAdmin, nil
	case "user":
		return RoleUser, nil
	case "guest":
		return RoleGuest, nil
	default:
		return "", &InvalidRoleError{Role(roleStr)}
	}
}

func (r Role) IsValidRole() bool {
	switch r {
	case RoleOwner, RoleAdmin, RoleUser, RoleGuest:
		return true
	default:
		return false
	}
}

func (r Role) Value() string {
	return string(r)
}
