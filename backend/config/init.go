package config

import "github.com/BurntSushi/toml"

var s *service

type service struct {
	Conf Conf
}

func init() {
	s = new(service)

	configPath := "./config.toml"
	c := new(Conf)
	_, err := toml.DecodeFile(configPath, &c)
	if err != nil {
		panic(err)
	}
	s.Conf = *c
}

func Get() *service {
	return s
}