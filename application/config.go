package application

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	MONGO_CONNECTION string
	MONGO_DB         string
	MONGO_COLLECTION string
	HTTP_PORT        string
}

func (configuration *Configuration) getConfig(env string) {
	gonfig.GetConf(fmt.Sprintf("./%s_config.json", env), configuration)
}

func NewConfig(env MeterEnvironment) Configuration {

	if env == "" {
		panic("No environment supplied to load config")
	}

	configuration := &Configuration{}
	configuration.getConfig(string(env))

	return *configuration
}
