package warehouse

import (
	"context"
	"order-service/package/structs"
)

type HTTPWarehouse interface {
	ReleaseStock(ctx context.Context, reqBody structs.RequestReleaseStock) error
	ReserveStock(ctx context.Context, reqBody structs.RequestReserveStock) error
	BatchStock(ctx context.Context, reqBody structs.RequestCreateOrder, dest interface{}) error
}