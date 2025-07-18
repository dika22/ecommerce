package main

import (
	"fmt"
	"log"
	"os"
	"warehouse-service/internal/domain/stock/repository"
	"warehouse-service/internal/domain/stock/usecase"
	"warehouse-service/package/config"
	"warehouse-service/package/connection/database"
	rabbitmq "warehouse-service/package/rabbit-mq"

	repoWh "warehouse-service/internal/domain/warehouse/repository"
	uw "warehouse-service/internal/domain/warehouse/usecase"

	api "warehouse-service/cmd/api"
	"warehouse-service/cmd/migrate"

	"github.com/urfave/cli/v2"
)

func main() {

  dbConf := config.NewDatabase()
  conf := config.NewConfig()
  conn := database.WebDB
  dbConn := database.NewDatabase(conn, dbConf)

  mqClient, err := rabbitmq.NewRabbitMQClient(conf)
  if err != nil {
    log.Println("ERROR INIT RABBITMQ", err)
  }

  fmt.Println("DEBUG", mqClient, err)


  repo := repository.NewRepositoryStock(dbConn)
  usecase := usecase.NewUsecaseStock(repo, mqClient)

  repoWarehouse := repoWh.NewRepositoryWarehouse(dbConn)
  usecaseWarehouse := uw.NewUsecaseWarehouse(repoWarehouse, mqClient)

  cmds := []*cli.Command{}
  cmds = append(cmds, api.ServeAPI(usecase, usecaseWarehouse)...)
  cmds = append(cmds, migrate.NewMigrate(dbConn)...)

  app := &cli.App{
    Name: "warehouse-service",
    Commands: cmds,
  }

  if err := app.Run(os.Args); err != nil {
    panic(err)
  }
}
