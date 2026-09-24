package user

import "strings"

type Role string

const (
	RoleOwner Role = "owner"
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
	RoleGuest Role = "guest"
)

const maxRoleLength = 20

func ParseRole(roleStr string) (Role, error) {
	if len(roleStr) > maxRoleLength {
		return "", ErrInvalidParseRole
	}

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
		return "", ErrInvalidParseRole
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

func (r Role) String() string {
	return string(r)
}
