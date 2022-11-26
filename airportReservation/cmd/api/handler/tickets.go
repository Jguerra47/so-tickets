package handler

import (
	"net/http"
	"strconv"

	"github.com/EAFIT-ST0257-Sistemas-Operativos/comunicacion-entre-procesos-y-consistencia-jprieto/airportReservation/internal/tickets"
	"github.com/EAFIT-ST0257-Sistemas-Operativos/comunicacion-entre-procesos-y-consistencia-jprieto/airportReservation/src/web"
	"github.com/gin-gonic/gin"
)

type Ticket struct {
	ticketService tickets.Service
}

func NewTicket(t tickets.Service) *Ticket {
	return &Ticket{
		ticketService: t,
	}
}

func (t *Ticket) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		chairs, err := t.ticketService.GetAll(ctx)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%v", err.Error())
			return
		}
		web.Success(ctx, http.StatusAccepted, chairs)
		return
	}
}

func (t *Ticket) GetByUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, err := strconv.Atoi(ctx.Param("userID"))
		tickets, err := t.ticketService.GetByUserId(ctx, userID)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%v", err.Error())
			return
		}
		web.Success(ctx, http.StatusAccepted, tickets)
		return
	}
}

func (t *Ticket) Reserve() gin.HandlerFunc {
	type userReservationRequest struct {
		UserID  int    `json:"user_id"`
		ChairID string `json:"chair_id"`
	}
	return func(ctx *gin.Context) {
		userRequest := userReservationRequest{}
		if err := ctx.ShouldBindJSON(&userRequest); err != nil {
			web.Error(ctx, http.StatusBadRequest, err.Error())
			return
		}
		ticket, err := t.ticketService.Reserve(ctx, userRequest.ChairID, userRequest.UserID)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%v", err.Error())
			return
		}
		web.Success(ctx, http.StatusAccepted, ticket)
		return
	}
}

func (t *Ticket) Delete() gin.HandlerFunc {
	type userDeleteRequest struct {
		UserID   int    `json:"user_id"`
		TicketID string `json:"ticket_id"`
	}
	return func(ctx *gin.Context) {
		userRequest := userDeleteRequest{}
		if err := ctx.ShouldBindJSON(&userRequest); err != nil {
			web.Error(ctx, http.StatusBadRequest, err.Error())
			return
		}
		err := t.ticketService.Delete(ctx, userRequest.TicketID, userRequest.UserID)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%v", err.Error())
			return
		}
		web.Success(ctx, http.StatusNoContent, "")
		return
	}
}

func (t *Ticket) AddPayment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ticketID := ctx.Param("ticketID")
		err := t.ticketService.AddPayment(ctx, ticketID)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%v", err.Error())
			return
		}
		web.Success(ctx, http.StatusNoContent, "")
		return
	}
}
