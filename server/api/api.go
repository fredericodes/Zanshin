package api

import (
	"server/repository"
	"server/util/configs"
)

const (
	ConfigsLoadErr = "configs could not be loaded"
	StartupErr     = "server could not start up"
)

type Server struct {
	DB repository.DB
}

func New(configs *configs.Configs) *Server {
	return &Server{
		DB: repository.New(configs.DbConf),
	}
}
