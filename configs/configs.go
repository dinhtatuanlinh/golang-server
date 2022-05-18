package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type config struct {
	Config struct {
		Database struct {
			Environment struct {
				POSTGRES_DB       string `yaml:"POSTGRES_DB"`
				POSTGRES_USER     string `yaml:"POSTGRES_USER"`
				POSTGRES_PASSWORD string `yaml:"POSTGRES_PASSWORD"`
				Ports             int64  `yaml:"ports"`
			} `yaml:"environment"`
		} `yaml:"database"`
	} `yaml:"config"`
}

func Conf() (*config, error) {
	buf, err := ioutil.ReadFile("./configs/config_server.yaml")
	if err != nil {
		return nil, err
	}
	var configs *config
	err = yaml.Unmarshal(buf, &configs)
	if err != nil {
		return nil, err
	}

	return configs, nil
}
