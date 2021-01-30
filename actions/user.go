package actions

import (
	"context"

	"github.com/vmkevv/duiztapi/ent"
)

// UserActions does all user related logic
type UserActions struct {
	ctx context.Context
	ent *ent.Client
}

// SetupUserActions returns a new instace of UserActions
func SetupUserActions(ctx context.Context, ent *ent.Client) UserActions {
	return UserActions{ctx, ent}
}

// Register register a new user in database
func (ua UserActions) Register(name, email string) (*ent.User, error) {
	return nil, nil
}

// SendEmailToken sends email magic link to login the system
func (ua UserActions) SendEmailToken(email string) error {
	return nil
}

// GenerateToken generates a token based on user ID
func (ua UserActions) GenerateToken(ID int) (string, error) {
	return "", nil
}

// ExistsEmail check if exists an account with the email provided
func (ua UserActions) ExistsEmail(email string) bool {
	return false
}

// Login validates a JWT token and returns the user credentials
func (ua UserActions) Login(token string) (*ent.User, error) {
	return nil, nil
}

// CheckToken verifies if a token is valid
func (ua UserActions) CheckEmailToken(token string) error {
	return nil
}
