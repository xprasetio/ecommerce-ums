package services

import (
	"context"
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/interfaces"
	"ecommerce-ums/internal/models"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)


type UserService struct {
	UserRepo interfaces.IUserRepository
}

func (s *UserService) Register (ctx context.Context, req *models.User, role string) (*models.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	req.Password = string(hashPassword)
	req.Role = role

	err = s.UserRepo.InsertNewUser(ctx, req)
	if err != nil {
		return nil, err
	}
	respn := req
	respn.Password = ""
	return respn, nil
	 
}
func (s *UserService) Login(ctx context.Context, req models.LoginRequest, role string) (models.LoginResponse, error) {
	var (
		response models.LoginResponse
		now      = time.Now()
	)

	userDetail, err := s.UserRepo.GetUserbyUsername(ctx, req.Username, role)
	if err != nil {
		return response, errors.Wrap(err, "failed to get user by username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(req.Password)); err != nil {
		return response, errors.Wrap(err, "failed to compare password")
	}

	token, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "token", userDetail.Email, now)
	if err != nil {
		return response, errors.Wrap(err, "failed to generate token")
	}

	refreshToken, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "refresh_token", userDetail.Email, now)
	if err != nil {
		return response, errors.Wrap(err, "failed to generate refresh token")
	}

	userSession := &models.UserSession{
		UserID:              userDetail.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(helpers.MapTypeToken["token"]),
		RefreshTokenExpired: now.Add(helpers.MapTypeToken["refresh_token"]),
	}
	err = s.UserRepo.InsertNewUserSession(ctx, userSession)
	if err != nil {
		return response, errors.Wrap(err, "failed to insert new session")
	}

	response.UserID = userDetail.ID
	response.Username = userDetail.Username
	response.FullName = userDetail.FullName
	response.Email = userDetail.Email
	response.Token = token
	response.RefreshToken = refreshToken

	return response, nil
}

func (s *UserService) GetProfile(ctx context.Context, username string) (models.User, error) {
	var (
		resp models.User
		err  error
	)

	resp, err = s.UserRepo.GetUserbyUsername(ctx, username, "")
	if err != nil {
		return resp, errors.Wrap(err, "failed to get user by username")
	}
	resp.Password = ""
	return resp, nil
}
func (s *UserService) Logout(ctx context.Context, token string) error {
	return s.UserRepo.DeleteUserSession(ctx, token)
}
