package services

import (
	"const/core/orm/models"
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Error struct {
	Message string `json:"message"`
}

type TenantService struct {
	db *sql.DB
}

func NewTenantService(db *sql.DB) *TenantService {
	return &TenantService{db: db}
}

func (s *TenantService) CreateTenant(ctx context.Context, tenant *models.Tenant) error {
	return tenant.Insert(ctx, s.db, boil.Infer())
}

func (s *TenantService) GetTenantByID(ctx context.Context, id int) (*models.Tenant, error) {
	return models.FindTenant(ctx, s.db, id)
}

func (s *TenantService) GetAllTenants(ctx context.Context) (models.TenantSlice, error) {
	return models.Tenants().All(ctx, s.db)
}

func (s *TenantService) UpdateTenant(ctx context.Context, tenant *models.Tenant) error {
	_, err := tenant.Update(ctx, s.db, boil.Infer())
	return err
}

func (s *TenantService) DeleteTenant(ctx context.Context, id int) error {
	tenant, err := s.GetTenantByID(ctx, id)
	if err != nil {
		return err
	}
	_, err = tenant.Delete(ctx, s.db)
	return err
}
