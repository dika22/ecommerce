package main

import (
	"log"
	"order-service/internal/domain/order/repository"
	"order-service/internal/domain/order/usecase"
	"order-service/package/config"
	"order-service/package/connection/database"
	rabbitmq "order-service/package/rabbit-mq"
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

  mqClient, err := rabbitmq.NewRabbitMQClient(conf)
  if err != nil {
    log.Println("ERROR INIT RABBITMQ", err)
  }

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
  usecase := usecase.NewOrderUsecase(repo, httpClient, workerClient, mqClient)
  cmds := []*cli.Command{}
  cmds = append(cmds, api.ServeAPI(usecase, httpClient, cacheConf)...)
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
