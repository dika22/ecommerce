package main

import (
	"log"
	"order-service/internal/domain/order/repository"
	"order-service/internal/domain/order/usecase"
	"order-service/package/config"
	"order-service/package/connection/database"
	"os"

	api "order-service/cmd/api"
	"order-service/cmd/migrate"
	"order-service/cmd/worker"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/urfave/cli/v2"

	http_client "order-service/package/http_client"
)

func main() {

  dbConf := config.NewDatabase()
  conf := config.NewConfig()
  conn := database.WebDB
  dbConn := database.NewDatabase(conn, dbConf)

  cacheConf := config.NewCache()
  nrApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName("order-service"),
		newrelic.ConfigLicense(conf.NewRelicLicense),
	)
	if err != nil {
		log.Print("ERROR INIT NEWRELIC", err)
	}

  httpClient := http_client.NewHTTPClients(conf)
  workerClient := worker.WorkerClient(cacheConf)
  repo := repository.NewOrderRepository(dbConn)
  usecase := usecase.NewOrderUsecase(repo, workerClient)
  cmds := []*cli.Command{}
  cmds = append(cmds, api.ServeAPI(usecase)...)
  cmds = append(cmds, migrate.NewMigrate(dbConn)...)
  cmds = append(cmds, worker.StartWorker(conf, cacheConf, repo, httpClient, workerClient, nrApp)...)

  app := &cli.App{
    Name: "order-service",
    Commands: cmds,
  }

  if err := app.Run(os.Args); err != nil {
    panic(err)
  }
}
