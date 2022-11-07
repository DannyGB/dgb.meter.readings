//go:build wireinject
// +build wireinject

package main

import (
	"dgb/meter.readings/internal/application"
	"dgb/meter.readings/internal/configuration"
	"dgb/meter.readings/internal/database"
	"github.com/google/wire"
)

func CreateApi() *application.ReadingApi {

	panic(wire.Build(
		configuration.NewMeterEnvironment,
		configuration.NewConfig,
		application.NewResponse,
		database.NewRepository,
		application.NewMiddleware,
		application.NewApi,
	))
}
