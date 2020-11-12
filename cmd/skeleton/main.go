package main

import (
	"github.com/urfave/cli/v2"
	"os"
	app "skeleton-code"
	"skeleton-code/components"
	"skeleton-code/components/vehicle"
	"skeleton-code/config"
	"skeleton-code/logger"
	"skeleton-code/server"
	"skeleton-code/utils"
)

func init() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	conf := config.LoadConfigFile()
	close := loggerInit(conf)
	defer close()

	container := app.NewDIContainer()
	container.Provide(
		config.LoadConfigFile,
		utils.NewCloser,
		vehicle.NewVehicleComponent,
		components.NewContext,
	)
	container.Invoke(
		server.GRPCServer,
		server.APIServer,
	)

	app := &cli.App{
		Name:    name,
		Usage:   usage,
		Version: version,
		Action: func(c *cli.Context) error {
			logger.Info("=====Skeleton Run=====")
			go container.Run()
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		logger.Error(err)
	}
}
