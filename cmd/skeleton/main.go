package main

import (
	"github.com/urfave/cli/v2"
	"os"
	app "skeleton-code"
	"skeleton-code/components"
	"skeleton-code/components/member"
	"skeleton-code/components/vehicle"
	"skeleton-code/components/vehicle/vehiclemodel"
	"skeleton-code/config"
	"skeleton-code/database"
	"skeleton-code/logger"
	"skeleton-code/server"
	"skeleton-code/server/handlers"
	"skeleton-code/utils"
)

func init() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	conf := config.LoadConfigFile()
	close := loggerInit(conf)
	defer close()

	app := skeletonApp()
	logger.Error(app.Run(os.Args))
}

func skeletonApp() *cli.App {
	container := app.NewDIContainer()
	container.Provide(
		config.LoadConfigFile,
		utils.NewCloser,
		database.NewDatabase,

		//vehiclemodel components
		vehiclemodel.NewVehicleModelRepository,
		vehiclemodel.NewVehicleService,

		//vehicle components
		vehicle.NewVehicleRepository,
		vehicle.NewVehicleService,

		//member components
		member.NewMemberService,

		// components context
		components.NewComponentContext,

		//server
		handlers.NewVehicleHandler,
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
			logger.Info("===== Skeleton Run =====")
			container.Run()
			return nil
		},
	}
	return app
}
