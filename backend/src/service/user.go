package service

import (
	"context"
	"encoding/json"
	"fmt"

	"fixit.com/backend/internal/auth"
	"fixit.com/backend/src/models"
	"fixit.com/backend/src/models/dto"
	"fixit.com/backend/src/repo"
)

type UserSvc struct {
	googleAuth auth.GoogleAuth
	userRepo   repo.UserRepo
}

func CreateUserSvc(userRepo repo.UserRepo, googleAuth auth.GoogleAuth) *UserSvc {
	return &UserSvc{userRepo: userRepo, googleAuth: googleAuth}
}

func (s *UserSvc) CreateUser(ctx context.Context, req *dto.SignupRequest) error {
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		Email:          req.Email,
		IsThirdParty:   false,
	}

	return s.userRepo.CreateUser(ctx, user)
}

// Login is used to login a user with username and password
// This doesn't work for third party users
func (s *UserSvc) Login(ctx context.Context, req *dto.LoginRequest) (string, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", fmt.Errorf("user not found")
	}

	if user.IsThirdParty {
		return "", fmt.Errorf("third party users cannot login with username and password")
	}

	valid := auth.VerifyPassword(req.Password, user.HashedPassword)
	if !valid {
		return "", fmt.Errorf("invalid password")
	}

	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserSvc) GoogleAuthCallback(ctx context.Context, code, state string) (string, error) {
	data, err := s.googleAuth.VerifyCallBack(ctx, code)
	if err != nil {
		return "", err
	}

	fmt.Println("got data: ", data)

	userData := &dto.GoogleCallBackResponse{}
	err = json.Unmarshal([]byte(data), userData)
	if err != nil {
		return "", err
	}

	user, err := s.userRepo.GetUserByUsername(ctx, userData.Username)
	if err != nil {
		return "", err
	}

	if user == nil {
		user = &models.User{
			Username:     userData.Username,
			Email:        userData.Email,
			FirstName:    userData.FirstName,
			LastName:     userData.LastName,
			IsThirdParty: false,
		}
		err = s.userRepo.CreateUser(ctx, user)
		if err != nil {
			return "", err
		}
	}

	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
