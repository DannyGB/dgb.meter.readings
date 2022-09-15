package application

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type ConfigurationService struct{}

type Configuration struct {
	MONGO_CONNECTION string
	HTTP_PORT        string
}

func (configuration *ConfigurationService) GetConfig(env string) Configuration {

	if env == "" {
		panic("Environment missing for configuration")
	}

	config := Configuration{}
	fileName := fmt.Sprintf("./%s_config.json", env)

	gonfig.GetConf(fileName, &config)

	return config
}
