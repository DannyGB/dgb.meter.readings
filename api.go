package main

import (
	"dgb/meter.readings/database"
	"fmt"
	"log"
	"net/http"
	"os"

	"dgb/meter.readings/application"

	"github.com/gorilla/mux"
)

const route = "/reading"

type ReadingApi struct {
	response *Response
	config   application.Configuration
}

func (api *ReadingApi) get(w http.ResponseWriter, r *http.Request) {
	api.response.Ok(ResponseParams{w: w, result: fmt.Sprintf("Get Readings: %s", api.config.MONGO_CONNECTION)})
}

func (api *ReadingApi) getById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	result := database.GetReading(id, api.config)

	if result == nil {
		api.response.NotFound(ResponseParams{w: w})
		return
	}

	api.response.Ok(ResponseParams{w, result})
}

func (api *ReadingApi) update(w http.ResponseWriter, r *http.Request) {
	api.response.Ok(ResponseParams{w, fmt.Sprintf("Update Reading by Id: %s", mux.Vars(r)["id"])})
}

func (api *ReadingApi) create(w http.ResponseWriter, r *http.Request) {
	api.response.Created(ResponseParams{w, fmt.Sprintf("Create Reading by Id: %s", mux.Vars(r)["id"])})
}

func (api *ReadingApi) handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	subRoute := myRouter.PathPrefix(route).Subrouter()

	subRoute.HandleFunc("/", api.get).Methods(http.MethodGet)
	subRoute.HandleFunc("/{id}", api.getById).Methods(http.MethodGet)
	subRoute.HandleFunc("/{id}", api.update).Methods(http.MethodPut)
	subRoute.HandleFunc("/{id}", api.create).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", api.config.HTTP_PORT), subRoute))
}

func NewApi(response *Response, configurationService *application.ConfigurationService) *ReadingApi {
	env := os.Getenv("METER_READINGS_ENVIRONMENT")

	if env == "" {
		env = "dev"
	}

	config := configurationService.GetConfig(env)
	return &ReadingApi{
		response,
		config,
	}
}
