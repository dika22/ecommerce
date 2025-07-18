package usecase

import (
	"context"
	"errors"
	"fmt"
	"order-service/cmd/worker/tasks"
	"order-service/package/structs"
)


func (u *OrderUsecase) CreateOrder(ctx context.Context, req *structs.RequestCreateOrder) error{
	u.mu.Lock()
	defer u.mu.Unlock()
	destBatchStock := structs.ResponseBatchStock{}
	if err := u.http_clients.WarehouseClient.BatchStock(ctx, *req, destBatchStock); err != nil {
		return err
	}
	for _, sp := range destBatchStock.StockProducts {
		if !sp.Available {
			msgError := fmt.Sprintf("Product dengan Id %v no have stock", sp.ProductID)
			return  errors.New(msgError)
		}
	}
	order := req.NewOrder()
	orderID,  err := u.repo.CreateOrder(ctx, order)
	if err != nil { 
		return err
	}
	orderItems := make([]structs.OrderItem, 0, len(req.Items))
	for _, oi := range req.Items {
		orderItem := oi.NewOrderItem(orderID)
		orderItems = append(orderItems, orderItem)
	}
	if err := u.repo.StoreOrderItems(ctx, orderItems); err != nil { 
		return err
	}

	//reserved stock
	job := tasks.ReleaseStockPayload{
		OrderID: orderID,
	}

	_, err = u.workerClient.EnqueueContext(ctx, job.Dispatch(), nil)
	if err != nil {
		return err
	}

	//update stock
	return nil
}