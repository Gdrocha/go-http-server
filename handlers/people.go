package handlers

import (
	"encoding/json"
	"go-http-server/models"
	"net/http"
)

func GetPeopleHandler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	people := []models.Person{
		{Name: "Gabriel", Nickname: "Torangex"},
		{Name: "Jonathan", Nickname: "Jaspion"},
		{Name: "Marco", Nickname: "Lebiróbi"},
	}

	json.NewEncoder(responseWriter).Encode(people)
}
