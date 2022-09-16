// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
  "dgb/meter.readings/application"
  "dgb/meter.readings/database"
)

// Injectors from container.go:

func CreateApi() *ReadingApi {
  response := &Response{}
  meterEnvironment := application.NewMeterEnvironment()
  configuration := application.NewConfig(meterEnvironment)
  repository := database.NewRepository(configuration)
  readingApi := NewApi(response, configuration, repository)
  return readingApi
}
