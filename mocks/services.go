package mocks

import "github.com/vmkevv/duiztapi/ent"

// UserActions hold all user related actions, included database actions
type UserActions interface {
	Register(name, email string) (*ent.User, error)
	SendEmailToken(email string) error
	ExistsEmail(email string) bool
	Login(token string) (*ent.User, error)
}
