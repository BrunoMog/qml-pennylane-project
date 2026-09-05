package user

import (
	"strings"
)

func ParseRole(roleStr string) (Role, error) {
	switch strings.ToLower(roleStr) {
	case "user":
		return RoleUser, nil
	case "admin":
		return RoleAdmin, nil
	case "owner":
		return RoleOwner, nil
	case "guest":
		return RoleGuest, nil
	default:
		return "", &InvalidRoleError{Role(roleStr)}
	}
}
