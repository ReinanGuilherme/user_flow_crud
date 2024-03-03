package main

import (
	"net/http"
	"os"

	database "github.com/ReinanGuilherme/user_flow_crud/src/config"
	"github.com/ReinanGuilherme/user_flow_crud/src/routes"
	"github.com/joho/godotenv"
)

// @title USER FLOW CRUD
// @version 1.0

func main() {

	godotenv.Load()

	database.ConnectDataBase()

	r := routes.SetupRouter()

	errorServer := http.ListenAndServe(os.Getenv("SERVER_PORT"), r)
	if errorServer != nil {
		panic(errorServer)
	}
}
