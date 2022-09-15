//go:build wireinject
// +build wireinject

package main

import (
	"dgb/meter.readings/application"
	"dgb/meter.readings/database"
	"github.com/google/wire"
)

func CreateApi() *ReadingApi {

	panic(wire.Build(
		application.NewMeterEnvironment,
		application.NewConfig,
		wire.Struct(new(Response), "*"),
		database.NewRepository,
		NewApi,
	))
}
