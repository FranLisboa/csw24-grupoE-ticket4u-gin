package services

import (
	"const/core/orm/models"
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type FinanceService struct {
	db *sql.DB
}

func NewFinanceService(db *sql.DB) *FinanceService {
	return &FinanceService{db: db}
}

func (s *FinanceService) RecordMovement(ctx context.Context, exec boil.ContextExecutor, movement *models.Movimentofinanceiro) error {
	return movement.Insert(ctx, exec, boil.Infer())
}

func (s *FinanceService) GetUserBalance(ctx context.Context, userID int) (float64, error) {
	movements, err := models.Movimentofinanceiros(models.MovimentofinanceiroWhere.Userid.EQ(userID)).All(ctx, s.db)
	if err != nil {
		return 0, err
	}

	var balance float64
	for _, m := range movements {
		if m.Tipomovimento == "credito" {
			valorFloat, _ := m.Valor.Float64()
			balance += valorFloat
		} else if m.Tipomovimento == "debito" {
			valorFloat, _ := m.Valor.Float64()
			balance -= valorFloat

		}
	}

	return balance, nil
}
