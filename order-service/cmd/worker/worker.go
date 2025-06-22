package worker

import (
	"context"
	"fmt"
	"log"
	"order-service/cmd/worker/tasks"
	"order-service/internal/constant"
	"order-service/internal/domain/order/repository"
	"order-service/package/config"
	"time"

	http_client "order-service/package/http_client"

	"github.com/hibiken/asynq"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/urfave/cli/v2"
)


type (
	WorkerHandler func(context.Context, *asynq.Task) error
	Worker struct {
		conf *config.Config
		cacheConf *config.Cache
		tasks tasks.Tasks
		nr *newrelic.Application
		http_clients http_client.HTTPClients
	}
)

const CmdServerWorker = "start-worker"
func WorkerClient(cf *config.Cache) *asynq.Client{
	return asynq.NewClient(asynq.RedisClientOpt{Addr: fmt.Sprintf("%s:%s", cf.WorkerRedisHost, cf.WorkerRedisPort)})
}

func (w Worker) StartWorker(*cli.Context) error  {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: fmt.Sprintf("%s:%s", w.cacheConf.WorkerRedisHost, w.cacheConf.WorkerRedisPort)},
		asynq.Config{
			Concurrency:      10,
			StrictPriority:   true,
			GroupGracePeriod: time.Second * 15,
			Queues: map[string]int{
				constant.QueueHigh:   5,
				constant.QueueMedium: 4,
				constant.QueueLow:    1,
			},
		},
	)
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeReleaseStock, w.NrWorkerMiddleware(w.tasks.StartStockReleaseJob, tasks.TypeReleaseStock))
	if err := srv.Run(mux); err!= nil {
		log.Fatalf("could not run serve: %v", err)
		return err
	}
	return nil
}

func (w Worker) NrWorkerMiddleware(f WorkerHandler, taskName string) func(ctx context.Context, task *asynq.Task) error {
	return func(ctx context.Context, task *asynq.Task) error {
		txn := w.nr.StartTransaction(taskName)
		defer txn.End()
		ctx = context.WithValue(ctx, constant.NewRelicTransactionCtx, txn)
		return f(ctx, task)
	}
 }
 

func StartWorker(conf *config.Config, c *config.Cache, repo repository.OrderRepository,
	http_clients http_client.HTTPClients, workerClient *asynq.Client, nrp *newrelic.Application) []*cli.Command {
	task := tasks.NewAsynqTask(conf, repo, http_clients, workerClient)
	w := Worker{conf: conf,cacheConf: c, tasks: task, nr: nrp, http_clients: http_clients}
	return []*cli.Command{
		{
			Name: CmdServerWorker,
			Usage: "Serve Order Worker",
			Action: w.StartWorker,
		},
	}
}