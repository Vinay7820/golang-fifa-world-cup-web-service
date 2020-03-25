package handlers

import (
	"golang-fifa-world-cup-web-service/data"
	"net/http"
)

// RootHandler returns an empty body status code
func RootHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNoContent, 204)
}

// ListWinners returns winners from the list
func ListWinners(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	year := req.URL.Query().Get(year)
	if year == "" {
		winners, err := data.ListAllJSON()
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError, 500)
			return
		}
		res.Write(winner)
	} else {
		filteredWinners, err := data.ListAllByYear(year)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest, 400)
			return
		}
		res.Write(filteredWinners)
	}
}

// AddNewWinner adds new winner to the list
func AddNewWinner(res http.ResponseWriter, req *http.Request) {

}

// WinnersHandler is the dispatcher for all /winners URL
func WinnersHandler(res http.ResponseWriter, req *http.Request) {

}
