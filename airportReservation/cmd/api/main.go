package main

import (
	"github.com/EAFIT-ST0257-Sistemas-Operativos/comunicacion-entre-procesos-y-consistencia-jprieto/airportReservation/cmd/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router := routes.NewRouter(r)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}
