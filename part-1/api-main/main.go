package main

import (
	"api-main/core"
	mesasRoutes "api-main/mesa/infraestructure/routes"
	productosRoutes "api-main/producto/infraestructure/routes"
	userRoutes "api-main/users/infraestructure/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	userRoutes.SetupRoutesCount(router)
	mesasRoutes.SetUpRoutes(router)
	productosRoutes.SetUpRoutes(router)

	port := ":8080"
	log.Println("Servidor escuchando en el puerto", port)
	log.Println(core.GetDBPool())
	log.Fatal(router.Run(port))

		
}