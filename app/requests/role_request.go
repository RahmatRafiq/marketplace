package requests

type RoleRequestAssignPermissions struct {
	PermissionIDs []uint `json:"permissions" form:"permissions" binding:"required" validate:"required"`
}
