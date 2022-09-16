package application

import (
	"dgb/meter.readings/internal/configuration"
	"dgb/meter.readings/internal/database"
	models "dgb/meter.readings/internal/models"
	"log"

	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"

	"strconv"
)

const route = "/reading"

type ReadingApi struct {
	response   *Response
	config     configuration.Configuration
	repository *database.Repository
}

func (api *ReadingApi) get(w http.ResponseWriter, r *http.Request) {
	skip := r.FormValue("skip")
	take := r.FormValue("take")

	s, _ := strconv.Atoi(skip)
	t, _ := strconv.Atoi(take)

	result := api.repository.GetAll(database.PageParams{Skip: s, Take: t})

	if result == nil {
		api.response.NotFound(ResponseParams{W: w})
		return
	}

	api.response.Ok(ResponseParams{w, result})
}

func (api *ReadingApi) getById(w http.ResponseWriter, r *http.Request) {
	result := api.repository.GetSingle(mux.Vars(r)["id"])

	if result == nil {
		api.response.NotFound(ResponseParams{W: w})
		return
	}

	api.response.Ok(ResponseParams{w, result})
}

func (api *ReadingApi) update(w http.ResponseWriter, r *http.Request) {

	var reading primitive.M

	if err := json.NewDecoder(r.Body).Decode(&reading); err != nil {
		api.response.BadRequest(ResponseParams{W: w})
		return
	}

	err := api.repository.Update(mux.Vars(r)["id"], reading)

	if err != nil {
		api.response.ServerError(ResponseParams{W: w})
	}

	api.response.Ok(ResponseParams{W: w})
}

func (api *ReadingApi) create(w http.ResponseWriter, r *http.Request) {
	var reading primitive.M

	if err := json.NewDecoder(r.Body).Decode(&reading); err != nil {
		api.response.BadRequest(ResponseParams{W: w})
		return
	}

	id, err := api.repository.Insert(reading)

	if err != nil {
		api.response.ServerError(ResponseParams{W: w})
	}

	api.response.Created(ResponseParams{w, &models.Created{
		Url: fmt.Sprintf("%v/%v", route, id),
	}})
}

func (api *ReadingApi) delete(w http.ResponseWriter, r *http.Request) {

	deletedCount, err := api.repository.Delete(mux.Vars(r)["id"])

	if deletedCount < 0 {
		api.response.NotFound(ResponseParams{W: w})
	}

	if err != nil {
		api.response.ServerError(ResponseParams{W: w})
	}

	api.response.Ok(ResponseParams{W: w})
}

func (api *ReadingApi) HandleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	subRoute := myRouter.PathPrefix(route).Subrouter()

	subRoute.HandleFunc("/{id}", api.getById).Methods(http.MethodGet)
	subRoute.HandleFunc("/{id}", api.update).Methods(http.MethodPut)
	subRoute.HandleFunc("/{id}", api.delete).Methods(http.MethodDelete)
	subRoute.HandleFunc("/{id}", api.create).Methods(http.MethodPost)

	subRoute.
		Path("/").
		Queries("skip", "{skip:[0-9]+}", "take", "{take:[0-9]+}").
		Methods(http.MethodGet).
		HandlerFunc(api.get)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", api.config.HTTP_PORT), subRoute))
}

func NewApi(response *Response, configuration configuration.Configuration, repository *database.Repository) *ReadingApi {
	return &ReadingApi{
		response,
		configuration,
		repository,
	}
}