package services

import (
	"const/core/orm/models"
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserService struct {
	db *sql.DB
}

type UserServiceInterface interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUsers(ctx context.Context) ([]*models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, id string, user *models.User) error
}

func NewService(db *sql.DB) UserService {
	return UserService{
		db: db,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	err := user.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUsers(ctx context.Context) ([]*models.User, error) {
	users, err := models.Users().All(ctx, s.db)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	user, err := models.FindUser(ctx, s.db, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int, user *models.User) error {
	_, err := user.Update(ctx, s.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
