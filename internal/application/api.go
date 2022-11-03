package application

import (
	"dgb/meter.readings/internal/configuration"
	"dgb/meter.readings/internal/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const route = "/reading"
const SKIP = "skip"
const TAKE = "take"
const SORT = "sort"
const FILTER = "filter"

type ReadingApi struct {
	response   *Response
	config     configuration.Configuration
	repository *database.Repository
}

func (api *ReadingApi) get(w http.ResponseWriter, r *http.Request) {
	skip, _ := strconv.Atoi(r.FormValue(SKIP))
	take, _ := strconv.Atoi(r.FormValue(TAKE))
	sort := r.FormValue(SORT)
	filter := r.FormValue(FILTER)

	result := api.repository.GetAll(database.PageParams{Skip: skip, Take: take, SortDirection: sort, Filter: filter})

	if result == nil {
		api.response.NotFound(ResponseParams{W: w})
		return
	}

	api.response.Ok(ResponseParams{w, result})
}

func (api *ReadingApi) count(w http.ResponseWriter, r *http.Request) {
	filter := r.FormValue(FILTER)

	result := api.repository.Count(database.PageParams{Filter: filter})
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

	api.response.Created(ResponseParams{w, &Created{
		Url: fmt.Sprintf("%v/%v", route, id),
	}})
}

func (api *ReadingApi) delete(w http.ResponseWriter, r *http.Request) {

	deletedCount, err := api.repository.Delete(mux.Vars(r)["id"])

	if deletedCount <= 0 {
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

	subRoute.HandleFunc("/{id:[0-9a-zA\\-]+}", func(w http.ResponseWriter, r *http.Request) {
		api.response.Ok(ResponseParams{W: w})
	}).Methods(http.MethodOptions)

	subRoute.
		Path("/count").
		Queries("filter", "{filter}").
		Methods(http.MethodGet).
		HandlerFunc(api.count)

	subRoute.HandleFunc("/{id:[0-9a-zA\\-]+}", api.getById).Methods(http.MethodGet)
	subRoute.HandleFunc("/{id:[0-9a-zA\\-]+}", api.update).Methods(http.MethodPut)
	subRoute.HandleFunc("/{id:[0-9a-zA\\-]+}", api.delete).Methods(http.MethodDelete)
	subRoute.HandleFunc("/{id:[0-9a-zA\\-]+}", api.create).Methods(http.MethodPost)

	subRoute.
		Path("/").
		Queries("skip", "{skip:[0-9]+}", "take", "{take:[0-9]+}", "sort", "{sort}", "filter", "{filter}").
		Methods(http.MethodGet, http.MethodOptions).
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
