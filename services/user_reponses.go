package services

import "github.com/vmkevv/duiztapi/ent"

// RegisterRes register response data
type RegisterRes struct {
	User  *ent.User `json:"user"`
	Token string    `json:"token"`
}

// RegisterReq register request data
type RegisterReq struct {
	Name  string `json:"name" validate:"required,gte=2"`
	Email string `json:"email" validate:"required,email"`
}

// SendEmailRes send email service response data
type SendEmailRes struct {
	Message string `json:"message"`
}

// SendEmailReq send email service request data
type SendEmailReq struct {
	Email string `json:"email" validate:"required,email"`
}

// LoginReq email token to verify a user account
type LoginReq struct {
	Token string `json:"token"`
}

// LoginRes data to return to user logged
type LoginRes struct {
	User  *ent.User `json:"user"`
	Token string    `json:"token"`
}
