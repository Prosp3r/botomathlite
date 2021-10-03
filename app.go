package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(username, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, dbname)
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func RespondWithJSON(w http.ResponseWriter, responseCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, responseCode int, message string) {
	RespondWithJSON(w, responseCode, map[string]string{"error": message})
}

func (a *App) getHomePage(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("Welcome to monolith Bot-o-math Solution")
	RespondWithJSON(w, http.StatusOK, message)
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.getHomePage).Methods("GET")
}
