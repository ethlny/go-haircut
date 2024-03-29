package router

import (
	"encoding/json"
	"gohairdresser/router/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Res struct
type Res struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// CORS
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"))

	serverStatusRes := Res{
		Code:    200,
		Status:  "OK",
		Message: "Server is running",
		Data:    nil,
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(serverStatusRes)
	})

	routes.AdminsRoutes(r)
	routes.AppointmentsRoutes(r)
	routes.AuthentificationRoutes(r)
	routes.ClientsRoutes(r)
	routes.HairdresserRoutes(r)
	routes.SaloonRoutes(r)

	return r
}
