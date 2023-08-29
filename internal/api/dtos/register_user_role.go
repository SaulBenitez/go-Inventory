package dtos

type RegisterUserRoleRequest struct {
	UserID int64 `json:"user_id" validate:"required"`
	RoleID int64 `json:"role_id" validate:"required"`
}