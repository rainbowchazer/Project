package enums

type UserPermission string

const (
	AdminPermission      UserPermission = "ADMIN_PERMISSION"
	SuperAdminPermission UserPermission = "SUPER_ADMIN_PERMISSION"
)
