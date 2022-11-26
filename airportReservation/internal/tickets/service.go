package tickets

import (
	"context"
	"github.com/EAFIT-ST0257-Sistemas-Operativos/comunicacion-entre-procesos-y-consistencia-jprieto/airportReservation/internal/domain"
	"time"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Chair, error)
	Reserve(ctx context.Context, chairId string, userId int) (domain.Ticket, error)
	GetByUserId(ctx context.Context, userId int) ([]domain.Ticket, error)
	Delete(ctx context.Context, tickerId string, userId int) error
	AddPayment(ctx context.Context, tickedID string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (r *service) GetAll(ctx context.Context) ([]domain.Chair, error) {
	tickets, err := r.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *service) Reserve(ctx context.Context, chairId string, userId int) (domain.Ticket, error) {
	ticket, err := r.repository.Reserve(ctx, chairId, userId)
	go func(ticketID string, userID int) {
		time.Sleep(10 * time.Second)
		if !r.repository.HasPayment(ctx, ticketID) {
			r.repository.Delete(ctx, ticketID, userID)
		}
	}(ticket.TicketID, ticket.UserID)
	return ticket, err
}

func (r *service) GetByUserId(ctx context.Context, userId int) ([]domain.Ticket, error) {
	tickets, _ := r.repository.GetByUserId(ctx, userId)
	return tickets, nil
}

func (r *service) Delete(ctx context.Context, tickerId string, userId int) error {
	err := r.repository.Delete(ctx, tickerId, userId)
	return err
}

func (r *service) AddPayment(ctx context.Context, tickedID string) error {
	return r.repository.AddPayment(ctx, tickedID)
}
