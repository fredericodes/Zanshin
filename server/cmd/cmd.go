package main

import (
	"fmt"
	"net/http"

	"server/api"
	"server/util/configs"
)

func main() {
	// Get configs of db, server and any other services
	config, err := configs.LoadConfigs()
	if err != nil {
		panic(api.ConfigsLoadErr)
	}

	// pass configs to server connection
	srv := api.New(config)
	appServer := api.InitializeServerRoutes(srv)
	service := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.ServerConf.Port),
		Handler: appServer,
	}

	if err := service.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(api.StartupErr)
	}
}
