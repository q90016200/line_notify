package routes

import (
	"lineNotify/api"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

var PORT string = ":" + os.Getenv("PORT")

func Register() {
	router := mux.NewRouter()

	lineNotify(router)

	if os.Getenv("PORT") == "" {
		log.Fatal("Error loading .env file")
	}

	log.Fatal(http.ListenAndServe(PORT, router))
}

func lineNotify(router *mux.Router) {
	lineNotify := router.PathPrefix("/lineNotify").Subrouter()
	lineNotify.HandleFunc("/auth", api.LineNotifyAuth).Methods("GET")
	lineNotify.HandleFunc("/callback", api.LineNotifyCallback).Methods("POST")
	lineNotify.HandleFunc("/notify", api.LineNotifySendNotify).Methods("POST")
}
