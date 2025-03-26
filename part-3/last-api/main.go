package main

import (
	dependenciesRobot "last-api/robot/infraestructure/dependencies"
	routesRobot "last-api/robot/infraestructure/routes"
	"log"

	dependenciesRS "last-api/robot-status/infraestructure/dependencies"
	routesRS "last-api/robot-status/infraestructure/routes"

	dependenciesH "last-api/historial-de-entregas/infraestructure/dependencies"
	routesHistorial "last-api/historial-de-entregas/infraestructure/routes"

	dependenciesCircuito "last-api/circuito/infraestructure/dependencies"
	routesCircuito "last-api/circuito/infraestructure/routes"

	dependenciesDetalles "last-api/detalles-del-pedido/infraestructure/dependencies"
	routerDetalle "last-api/detalles-del-pedido/infraestructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	// Robot
	dependenciesRobot.Init()
	routesRobot.RobotRoutes(router)

	//RS
	dependenciesRS.Init()
	routesRS.RSroutes(router)

	//Historial
	dependenciesH.Init()
	routesHistorial.HistorialRoutes(router)

	//Circuito
	dependenciesCircuito.Init()
	routesCircuito.CircuitRoutes(router)

	//Detalles del pedido
	dependenciesDetalles.Init()
	routerDetalle.Detallesroutes(router)

	port := ":8082"
	log.Println("Servidor escuchando en el puerto", port)
	log.Fatal(router.Run(port))
}