package main

import (
	"os"
	"product-service/internal/domain/product/repository"
	"product-service/internal/domain/product/usecase"
	"product-service/package/config"
	"product-service/package/connection/database"

	api "product-service/cmd/api"
	"product-service/cmd/migrate"

	"github.com/urfave/cli/v2"
)

func main() {

  dbConf := config.NewDatabase()
  // conf := config.NewConfig()
  conn := database.WebDB
  dbConn := database.NewDatabase(conn, dbConf)
  repo := repository.NewProductRepository(dbConn)
  usecase := usecase.NewProductUsecase(repo)
  cmds := []*cli.Command{}
  cmds = append(cmds, api.ServeAPI(usecase)...)
  cmds = append(cmds, migrate.NewMigrate(dbConn)...)

  app := &cli.App{
    Name: "product-service",
    Commands: cmds,
  }

  if err := app.Run(os.Args); err != nil {
    panic(err)
  }
}
