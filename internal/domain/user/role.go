package user

type Role string

const (
	RoleOwner Role = "owner"
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
	RoleGuest Role = "guest"
)

func (r Role) IsValidRole() bool {
	switch r {
	case RoleOwner, RoleAdmin, RoleUser, RoleGuest:
		return true
	default:
		return false
	}
}
