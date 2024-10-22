package services

import (
	"const/core/orm/models"
	. "const/core/services/finance"
	. "const/core/services/ticket"
	"context"
	"database/sql"
	"errors"
	"fmt"

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

type Error struct {
	Message string `json:"message"`
}
type Transaction = models.Transacao

// @summary Purchase a ticket
// @description Purchase a ticket
// @tags transactions
// @accept json
// @produce json
// @param transaction body Transaction true "Transaction object"
// @success 201 {object} Transaction
// @failure 400 {object} Error
// @failure 500 {object} Error
// @router /api/v1/transactions [post]
func (s *TransactionService) PurchaseTicket(ctx context.Context, transaction *models.Transacao) error {
	fmt.Println("entrando no service")

	tx, err := s.db.BeginTx(ctx, nil)
	fmt.Println("erro no begin", err)
	if err != nil {
		return err
	}

	ticket, err := models.FindTicket(ctx, tx, transaction.Iddoticket)
	fmt.Println("erro no findticket", ticket)
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

// @summary Request a refund
// @description Request a refund
// @tags transactions
// @accept json
// @produce json
// @param id path int true "Transaction ID"
// @success 200 {object} Transaction
// @failure 400 {object} Error
// @failure 500 {object} Error
// @router /api/v1/transactions/{id}/refund [put]
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
