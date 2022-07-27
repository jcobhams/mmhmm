package services

import (
	"context"
	"crypto/md5"
	"fmt"
)

type (
	UserService struct {
	}
)

var _ UserServiceInterface = (*UserService)(nil)

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GenerateUserId(ctx context.Context, email string) (string, error) {
	if email == "" {
		return "", fmt.Errorf("email is required")
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(email))), nil
}
