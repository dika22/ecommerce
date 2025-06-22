package main

import (
	"os"
	"shop-service/internal/domain/shop/repository"
	"shop-service/internal/domain/shop/usecase"
	"shop-service/package/config"
	"shop-service/package/connection/database"

	api "shop-service/cmd/api"
	"shop-service/cmd/migrate"

	"github.com/urfave/cli/v2"
)

func main() {

  dbConf := config.NewDatabase()
  // conf := config.NewConfig()
  conn := database.WebDB
  dbConn := database.NewDatabase(conn, dbConf)

  repo := repository.NewShopRepository(dbConn)
  usecase := usecase.NewShopUsecase(repo)
  cmds := []*cli.Command{}
  cmds = append(cmds, api.ServeAPI(usecase)...)
  cmds = append(cmds, migrate.NewMigrate(dbConn)...)

  app := &cli.App{
    Name: "shop-service",
    Commands: cmds,
  }

  if err := app.Run(os.Args); err != nil {
    panic(err)
  }
}
