package routes

import (
	"github.com/EAFIT-ST0257-Sistemas-Operativos/comunicacion-entre-procesos-y-consistencia-jprieto/airportReservation/cmd/api/handler"
	"github.com/EAFIT-ST0257-Sistemas-Operativos/comunicacion-entre-procesos-y-consistencia-jprieto/airportReservation/internal/domain"
	"github.com/EAFIT-ST0257-Sistemas-Operativos/comunicacion-entre-procesos-y-consistencia-jprieto/airportReservation/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
}

func NewRouter(r *gin.Engine) Router {
	return &router{r: r}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildPing()
	r.buildTickets()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildPing() {
	r.r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "pong")
	},
	)
}

func (r *router) buildTickets() {
	ticket := map[string]domain.Chair{
		"a57": {
			ChairID: "a57",
			Price:   10000,
		},
		"a58": {
			ChairID: "a58",
			Price:   12000,
		},
		"c39": {
			ChairID: "c39",
			Price:   12000,
		},
		"c50": {
			ChairID: "c50",
			Price:   12000,
		},
		"b11": {
			ChairID: "b11",
			Price:   12000,
		},
		"r45": {
			ChairID: "r45",
			Price:   12000,
		},
		"r46": {
			ChairID: "r46",
			Price:   12000,
		},
		"r47": {
			ChairID: "r47",
			Price:   12000,
		},
	}
	repo := tickets.NewRepository(ticket)
	service := tickets.NewService(repo)
	handler := handler.NewTicket(service)
	r.rg.GET("/list", handler.GetAll())
	r.rg.GET("/list/:userID", handler.GetByUserId())
	r.rg.POST("/reserve", handler.Reserve())
	r.rg.DELETE("/delete", handler.Delete())
	r.rg.POST("/pay/:ticketID", handler.AddPayment())
}
