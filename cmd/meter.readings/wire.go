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
		wire.Struct(new(application.Response), "*"),
		database.NewRepository,
		application.NewApi,
	))
}
