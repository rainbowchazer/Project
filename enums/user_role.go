package enums

type UserRole struct {
	Name            string
	UserPermissions []UserPermission
}

var (
	Admin = UserRole{
		Name:            "Администратор",
		UserPermissions: []UserPermission{AdminPermission},
	}

	SuperAdmin = UserRole{
		Name:            "Супер-администратор",
		UserPermissions: []UserPermission{AdminPermission, SuperAdminPermission},
	}
)

func (role UserRole) GetAuthorities() []string {
	authorities := make([]string, len(role.UserPermissions))
	for i, permission := range role.UserPermissions {
		authorities[i] = string(permission)
	}
	return authorities
}
