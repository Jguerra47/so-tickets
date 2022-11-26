package tickets

import (
	"context"
	"errors"
	"math/rand"

	"github.com/EAFIT-ST0257-Sistemas-Operativos/comunicacion-entre-procesos-y-consistencia-jprieto/airportReservation/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Chair, error)
	Reserve(ctx context.Context, chairId string, userId int) (domain.Ticket, error)
	GetByUserId(ctx context.Context, userId int) ([]domain.Ticket, error)
	Delete(ctx context.Context, tickerId string, userId int) error
	HasPayment(ctx context.Context, tickedID string) bool
	AddPayment(ctx context.Context, tickedID string) error
}

type repository struct {
	chairs        map[string]domain.Chair
	tickets       map[string][]domain.Ticket
	ticketsByUser map[int][]*domain.Ticket
}

func NewRepository(chair map[string]domain.Chair) Repository {
	return &repository{
		chairs:        chair,
		tickets:       map[string][]domain.Ticket{},
		ticketsByUser: map[int][]*domain.Ticket{},
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Chair, error) {
	chairs := []domain.Chair{}
	for _, chair := range r.chairs {
		chairs = append(chairs, chair)
	}
	return chairs, nil
}

func (r *repository) GetByUserId(ctx context.Context, userId int) ([]domain.Ticket, error) {
	tickets := r.ticketsByUser[userId]
	ticketsAns := []domain.Ticket{}
	for _, ticket := range tickets {
		ticketsAns = append(ticketsAns, *ticket)
	}
	return ticketsAns, nil
}

func (r *repository) Reserve(ctx context.Context, chairId string, userId int) (domain.Ticket, error) {
	chair, ok := r.chairs[chairId]
	if !ok {
		return domain.Ticket{}, errors.New("chair doesnt exist")
	}
	delete(r.chairs, chairId)
	ticket := domain.Ticket{
		TicketID: RandTicket(6),
		Chair:    chair,
		UserID:   userId,
		Payment:  "pending",
	}
	r.tickets[ticket.TicketID] = append(r.tickets[ticket.TicketID], ticket)
	r.ticketsByUser[userId] = append(r.ticketsByUser[userId], &r.tickets[ticket.TicketID][len(r.tickets[ticket.TicketID])-1])
	return ticket, nil
}

func (r *repository) Delete(ctx context.Context, tickerId string, userId int) error {
	for i, ticket := range r.ticketsByUser[userId] {
		if ticket.TicketID == tickerId {
			r.chairs[ticket.Chair.ChairID] = ticket.Chair
			r.ticketsByUser[userId] = append(r.ticketsByUser[userId][:i], r.ticketsByUser[userId][i+1:]...)
			delete(r.tickets, tickerId)
			return nil
		}
	}
	return errors.New("ticket dont exist")
}

func (r *repository) HasPayment(ctx context.Context, tickedID string) bool {
	for _, ticket := range r.tickets[tickedID] {
		if ticket.TicketID == tickedID {
			return ticket.Payment == "Accepted"
		}
	}
	return false
}
func (r *repository) AddPayment(ctx context.Context, tickedID string) error {
	for i, ticket := range r.tickets[tickedID] {
		if ticket.TicketID == tickedID {
			r.tickets[tickedID][i].Payment = "Accepted"
			return nil
		}
	}
	return nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandTicket(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
