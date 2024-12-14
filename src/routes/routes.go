package routes

import (
	"go-http-server/src/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", handlers.GetBaseRoute)
	http.HandleFunc("/people", handlers.GetPeopleHandler)
}
