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
	CreateUser(ctx context.Context, user *models.Usuario) error
	GetUsers(ctx context.Context) ([]*models.Usuario, error)
	GetUserByID(ctx context.Context, id string) (*models.Usuario, error)
	UpdateUser(ctx context.Context, id string, user *models.Usuario) error
	GetUserByEmail(ctx context.Context, email string) (*models.Usuario, error)
	DeleteUser(ctx context.Context, userID int) error
	CreateNotificationPreferences(ctx context.Context, preferenciasdenotificacao *models.Preferenciasdenotificacao) error
}

func NewService(db *sql.DB) UserService {
	return UserService{
		db: db,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.Usuario) error {
	err := user.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUsers(ctx context.Context) ([]*models.Usuario, error) {
	users, err := models.Usuarios().All(ctx, s.db)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*models.Usuario, error) {
	user, err := models.FindUsuario(ctx, s.db, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int, user *models.Usuario) error {
	_, err := user.Update(ctx, s.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.Usuario, error) {
	return models.Usuarios(models.UsuarioWhere.Email.EQ(email)).One(ctx, s.db)
}

func (s *UserService) DeleteUser(ctx context.Context, userID int) error {
	user, err := s.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	_, err = user.Delete(ctx, s.db)
	return err
}

func (s *UserService) CreateNotificationPreferences(ctx context.Context, preferenciasdenotificacao *models.Preferenciasdenotificacao) error {
	err := preferenciasdenotificacao.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
