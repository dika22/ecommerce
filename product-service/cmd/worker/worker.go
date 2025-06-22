package worker

import (
	"product-service/package/config"
	rabbitmq "product-service/package/rabbit-mq"

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