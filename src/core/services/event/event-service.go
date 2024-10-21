package services

import (
	"const/core/orm/models"
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type EventoService struct {
	db *sql.DB
}

func NewEventoService(db *sql.DB) *EventoService {
	return &EventoService{db: db}
}

func (s *EventoService) CreateEvento(ctx context.Context, evento *models.Evento) error {
	return evento.Insert(ctx, s.db, boil.Infer())
}

func (s *EventoService) GetEventoByID(ctx context.Context, id int) (*models.Evento, error) {
	return models.FindEvento(ctx, s.db, id)
}

func (s *EventoService) GetAllEventos(ctx context.Context) (models.EventoSlice, error) {
	return models.Eventos().All(ctx, s.db)
}

func (s *EventoService) UpdateEvento(ctx context.Context, evento *models.Evento) error {
	_, err := evento.Update(ctx, s.db, boil.Infer())
	return err
}

func (s *EventoService) DeleteEvento(ctx context.Context, id int) error {
	evento, err := s.GetEventoByID(ctx, id)
	if err != nil {
		return err
	}
	_, err = evento.Delete(ctx, s.db)
	return err
}

func (s *EventoService) GetEventosByTenant(ctx context.Context, tenantID int) (models.EventoSlice, error) {
	return models.Eventos(models.EventoWhere.Tenantid.EQ(tenantID)).All(ctx, s.db)
}
