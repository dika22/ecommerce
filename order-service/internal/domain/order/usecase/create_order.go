package usecase

import (
	"context"
	"order-service/cmd/worker/tasks"
	"order-service/package/structs"
)


func (u *OrderUsecase) CreateOrder(ctx context.Context, req *structs.RequestCreateOrder) error{
	// call http cek stock service
	order := req.NewOrder()
	orderID,  err := u.repo.CreateOrder(ctx, order)
	if err != nil {
		return err
	}
	for _, oi := range req.Items {
		orderItem := oi.NewOrderItem(orderID)
		if err := u.repo.StoreOrderItem(ctx, &orderItem); err != nil {
			return err
		}
	}

	job := tasks.ReleaseStockPayload{
		OrderID: orderID,
	}

	_, err = u.workerClient.EnqueueContext(ctx, job.Dispatch(), nil)
	if err != nil {
		return err
	}
	return nil
}