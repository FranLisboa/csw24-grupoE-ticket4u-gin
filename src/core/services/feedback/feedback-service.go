package services

import (
	"const/core/orm/models"
	"context"
	"database/sql"
	"errors"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type FeedbackService struct {
	db *sql.DB
}

func NewFeedbackService(db *sql.DB) *FeedbackService {
	return &FeedbackService{db: db}
}

func (s *FeedbackService) CreateFeedback(ctx context.Context, feedback *models.Avaliacao) error {
	transacao, err := models.FindTransacao(ctx, s.db, feedback.Transacaoid)
	if err != nil {
		return errors.New("transação não encontrada")
	}

	if transacao.Iddocomprador != feedback.Compradorid {
		return errors.New("o comprador não está associado a esta transação")
	}

	ticket, err := models.FindTicket(ctx, s.db, transacao.Iddoticket)
	if err != nil {
		return errors.New("ingresso não encontrado na transação")
	}

	if ticket.Iddovendedor != feedback.Vendedorid {
		return errors.New("o vendedor não está associado a esta transação")
	}

	exists, err := models.Avaliacaos(
		models.AvaliacaoWhere.Transacaoid.EQ(feedback.Transacaoid),
	).Exists(ctx, s.db)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("avaliação já enviada para esta transação")
	}

	return feedback.Insert(ctx, s.db, boil.Infer())
}

func (s *FeedbackService) GetFeedbacksByVendedor(ctx context.Context, vendedorid int) (models.AvaliacaoSlice, error) {
	return models.Avaliacaos(
		models.AvaliacaoWhere.Vendedorid.EQ(vendedorid),
	).All(ctx, s.db)
}
