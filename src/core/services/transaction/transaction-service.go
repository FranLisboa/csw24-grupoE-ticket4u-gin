package services

import (
	"const/core/orm/models"
	. "const/core/services/finance"
	. "const/core/services/ticket"
	"context"
	"database/sql"
	"errors"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TransactionService struct {
	db             *sql.DB
	ticketService  *TicketService
	financeService *FinanceService
}

func NewTransactionService(db *sql.DB, ticketService *TicketService, financeService *FinanceService) *TransactionService {
	return &TransactionService{
		db:             db,
		ticketService:  ticketService,
		financeService: financeService,
	}
}

func (s *TransactionService) PurchaseTicket(ctx context.Context, transaction *models.Transacao) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	ticket, err := models.FindTicket(ctx, tx, transaction.Iddoticket)
	if err != nil {
		return err
	}

	if ticket.Status != "disponivel" {
		return errors.New("ingresso não está disponível")
	}

	ticket.Status = "vendido"
	_, err = ticket.Update(ctx, tx, boil.Whitelist("status"))
	if err != nil {
		return err
	}

	transaction.Statusdatransacao = "concluida"
	err = transaction.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return err
	}

	movement := &models.Movimentofinanceiro{
		Userid:        ticket.Iddovendedor,
		Valor:         transaction.Precodevenda,
		Tipomovimento: "credito",
		Descricao:     null.StringFrom("Venda de ingresso"),
	}
	err = s.financeService.RecordMovement(ctx, tx, movement)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *TransactionService) RequestRefund(ctx context.Context, transactionID int) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	transaction, err := models.FindTransacao(ctx, tx, transactionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("transação não encontrada")
		}
		return err
	}

	if transaction.Statusdatransacao != "concluida" {
		return errors.New("transação não é elegível para reembolso")
	}

	ticket, err := models.FindTicket(ctx, tx, transaction.Iddoticket)
	if err != nil {
		return err
	}

	transaction.Statusdatransacao = "reembolsada"
	_, err = transaction.Update(ctx, tx, boil.Whitelist("status_da_transacao"))
	if err != nil {
		return err
	}

	ticket.Status = "disponivel"
	ticket.Usado = null.BoolFrom(false)
	_, err = ticket.Update(ctx, tx, boil.Whitelist("status", "usado"))
	if err != nil {
		return err
	}

	debitMovement := &models.Movimentofinanceiro{
		Userid:        ticket.Iddovendedor,
		Valor:         transaction.Precodevenda,
		Tipomovimento: "debito",
		Descricao:     null.StringFrom("Reembolso de venda de ingresso"),
	}
	err = s.financeService.RecordMovement(ctx, tx, debitMovement)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *TransactionService) GetTransactionByID(ctx context.Context, transactionID int) (*models.Transacao, error) {
	return models.FindTransacao(ctx, s.db, transactionID)
}

func (s *TransactionService) GetTransactionsByUser(ctx context.Context, userID int) (models.TransacaoSlice, error) {
	return models.Transacaos(
		models.TransacaoWhere.Iddocomprador.EQ(userID),
	).All(ctx, s.db)
}
