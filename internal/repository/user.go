package repository

import (
	"context"
	"ecommerce-ums/internal/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertNewUser(ctx context.Context, user *models.User) error { 
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserbyUsername(ctx context.Context, username string, role string) (models.User, error) {
	var (
		user models.User 
		err error
	)

	sql := r.DB.Where("username = ?", username)
	if role != "" {
		sql = sql.Where("role = ?", role)
	}
	err = sql.First(&user).Error
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}
func (r *UserRepository) InsertNewUserSession(ctx context.Context, session *models.UserSession) error {
	return r.DB.Create(session).Error
}

func (r *UserRepository) GetUserSessionByToken(ctx context.Context, token string) (models.UserSession, error) {
	var (
		session models.UserSession
		err     error
	)
	err = r.DB.Where("token = ?", token).First(&session).Error
	if err != nil {
		return session, err
	}
	if session.ID == 0 {
		return session, errors.New("session not found")
	}
	return session, nil
}

func (r *UserRepository) GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (models.UserSession, error) {
	var (
		session models.UserSession
		err     error
	)
	err = r.DB.Where("refresh_token = ?", refreshToken).First(&session).Error
	if err != nil {
		return session, err
	}
	if session.ID == 0 {
		return session, errors.New("session not found")
	}
	return session, nil
}

func (r *UserRepository) UpdateTokenByRefreshToken(ctx context.Context, token string, refreshToken string) error {
	return r.DB.Exec("UPDATE user_sessions SET token = ? WHERE refresh_token = ?", token, refreshToken).Error
}

func (r *UserRepository) DeleteUserSession(ctx context.Context, token string) error {
	return r.DB.Exec("DELETE FROM user_sessions WHERE token = ?", token).Error
}

