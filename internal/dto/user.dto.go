package dto

import (
	"mime/multipart"
	"time"
)

type RegisterRequest struct {
	FullName string `json:"fullname" binding:"required,max=255"`
	Email    string `json:"email" binding:"required,email,max=255"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"maruf321@mail.com"`
	Password string `json:"password" binding:"required" example:"moonchaser1234"`
}

type UserResponse struct {
	Id           int        `json:"id"`
	Role         string     `json:"role"`
	Fullname     string     `json:"fullname"`
	Email        string     `json:"email"`
	PhotoProfile *string    `json:"photo_profile"`
	JobPosition  *string    `json:"job_position"`
	Location     *string    `json:"location"`
	Bio          *string    `json:"bio"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type UpdateProfileRequest struct {
	PhotoProfile *multipart.FileHeader `form:"photo_profile" binding:"omitempty"`
	Fullname     *string               `form:"fullname" binding:"omitempty,max=255"`
	Location     *string               `form:"location" binding:"omitempty,max=255"`
	JobPosition  *string               `form:"job_position" binding:"omitempty,max=100"`
	Bio          *string               `form:"bio" binding:"omitempty"`
}

type ChangePassword struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

type LoginResponse struct {
	Token        string  `json:"token"`
	Fullname     string  `json:"fullname"`
	Email        string  `json:"email"`
	PhotoProfile *string `json:"photo_profile"`
}
