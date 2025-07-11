package worker

import (
	"warehouse-service/package/config"

	"github.com/urfave/cli/v2"
)


type Worker struct{
	conf config.Config
}

func WorkerClient() {
}

func (w Worker) StartWorker(*cli.Context) error  {
	return nil
}

func StartWorker() []*cli.Command {
	return []*cli.Command{}
}