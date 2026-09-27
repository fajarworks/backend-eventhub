package dto

import (
	"mime/multipart"
	"time"
)

type RegisterRequest struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
	PhotoProfile *multipart.FileHeader `form:"photo_profile"`
	Fullname     *string               `form:"fullname"`
	Location     *string               `form:"location"`
	JobPosition  *string               `form:"job_position"`
	Bio          *string               `form:"bio"`
}
