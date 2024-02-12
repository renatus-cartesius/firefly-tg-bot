package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Firefly  Firefly
	Telegram Telegram
}

type Telegram struct {
	Token string `yaml:"token"`
}

type Firefly struct {
	Token  string `yaml:"token"`
	ApiUrl string `yaml:"apiUrl"`
}

func ReadConfig(configpath string) (*Config, error) {
	buf, err := os.ReadFile(configpath)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		return nil, fmt.Errorf("config file %q parse error:\n\t%w", configpath, err)
	}
	return c, nil
}
