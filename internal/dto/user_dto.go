package dto

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=admin user manager"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" binding:"omitempty,min=3"`
	Email string `json:"email" binding:"omitempty,email"`
	Role  string `json:"role" binding:"omitempty,oneof=admin user manager"`
}
