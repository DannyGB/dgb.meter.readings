package application

import "os"

type MeterEnvironment string

func NewMeterEnvironment() MeterEnvironment {
	return MeterEnvironment(os.Getenv("METER_READINGS_ENVIRONMENT"))
}
