package worker

import (
	"shop-service/package/config"
	rabbitmq "shop-service/package/rabbit-mq"

	"github.com/urfave/cli/v2"
)


type Worker struct{
	conf config.Config
}

func WorkerClient() {
	rabbitmq.Connection()
}

func (w Worker) StartWorker(*cli.Context) error  {
	return nil
}

func StartWorker() []*cli.Command {
	return []*cli.Command{}
}