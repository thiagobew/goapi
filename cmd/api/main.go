package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/thiagobew/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API server on port 8000")

	err := http.ListenAndServe(":8000", r)

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Failed to start server")
	}
}