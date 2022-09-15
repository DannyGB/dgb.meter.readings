//go:build wireinject
// +build wireinject

package main

import (
	"dgb/meter.readings/application"
	"github.com/google/wire"
)

func CreateApi() *ReadingApi {

	panic(wire.Build(
		wire.Struct(new(Response), "*"),
		wire.Struct(new(application.ConfigurationService), "*"),
		NewApi,
	))
}
