package tasks

import (
	"context"
	"encoding/json"
	"order-service/package/structs"

	"github.com/hibiken/asynq"
)

type ReleaseStockPayload struct {
	OrderID int64 `json:"order_id"`
}

func (rsp ReleaseStockPayload) Dispatch() *asynq.Task {
	marshal, err := json.Marshal(rsp)
	if err != nil {
		return nil
	}
	return asynq.NewTask(
		TypeReleaseStock, 
		marshal,
		asynq.Queue("lender:queue:low"),
	)
}

func (a AsyncTask) StartStockReleaseJob(ctx context.Context, task *asynq.Task) error {
	var p ReleaseStockPayload
	if err := json.Unmarshal(task.Payload(), &p); err != nil {
		return err
	}

	var order structs.Order
	if err := a.repo.GetOrderByOrderId(ctx, p.OrderID, order); err != nil {
		return err
	}

	// Jika sudah dibayar, tidak perlu release
	if order.Status != 1 {
		return nil
	}

	// Mark as expired
	if err := a.repo.UpdateOrderStatus(ctx, p.OrderID); err != nil { return err } 

	// Release stock via Stock Service
	orderItems := []*structs.OrderItem{}
	if err := a.repo.GetOrderItemsByOrderId(ctx, p.OrderID, &orderItems); err != nil {
		return err
	}
	for _, item := range orderItems {
		req := structs.RequestReleaseStock{
			OrderID: item.OrderID,
			ProductID: item.ProductID,
			WarehouseID: item.WarehouseID,
			Quantity: item.Quantity,
		}
		if err := a.http_clients.WarehouseClient.ReleaseStock(ctx, req); err != nil{
			return err
		}
	}
	return nil
}