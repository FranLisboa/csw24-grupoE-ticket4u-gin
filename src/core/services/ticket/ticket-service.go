package services

import (
	"const/core/orm/models"
	"context"
	"database/sql"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TicketService struct {
	db *sql.DB
}

func NewTicketService(db *sql.DB) *TicketService {
	return &TicketService{db: db}
}

func (s *TicketService) CreateTicket(ctx context.Context, ticket *models.Ticket) error {
	return ticket.Insert(ctx, s.db, boil.Infer())
}

func (s *TicketService) GetTicketByID(ctx context.Context, id int) (*models.Ticket, error) {
	return models.FindTicket(ctx, s.db, id)
}

func (s *TicketService) GetAvailableTicketsByEvent(ctx context.Context, eventID int) (models.TicketSlice, error) {
	return models.Tickets(
		models.TicketWhere.Eventoid.EQ(eventID),
		models.TicketWhere.Status.EQ("disponivel"),
	).All(ctx, s.db)
}

func (s *TicketService) UpdateTicketStatus(ctx context.Context, ticketID int, status string) error {
	ticket, err := s.GetTicketByID(ctx, ticketID)
	if err != nil {
		return err
	}
	ticket.Status = status
	_, err = ticket.Update(ctx, s.db, boil.Whitelist("status"))
	return err
}

func (s *TicketService) GetTicketsBySeller(ctx context.Context, sellerID int) (models.TicketSlice, error) {
	return models.Tickets(
		models.TicketWhere.Iddovendedor.EQ(sellerID),
	).All(ctx, s.db)
}

func (s *TicketService) MarkTicketAsUsed(ctx context.Context, ticketID int) error {
	ticket, err := s.GetTicketByID(ctx, ticketID)
	if err != nil {
		return err
	}
	ticket.Status = "usado"
	ticket.Usado = null.BoolFrom(true)
	_, err = ticket.Update(ctx, s.db, boil.Whitelist("status", "usado"))
	return err
}
