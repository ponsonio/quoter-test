package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ponsonio/quoter-test/mortgage/calculator"
	"io"
	"net/http"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}

	r := mux.NewRouter()

	r.HandleFunc("/mortgage/calculate/", a.calculate).Methods(http.MethodPost)

	a.router = r
	return a
}

func (a *api) calculate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := io.ReadAll(r.Body)
	var calc calculator.MortgageCalc
	json.Unmarshal(reqBody, &calc)
	json.NewEncoder(w).Encode(calc)
}

func (a *api) Router() http.Handler {
	return a.router
}
