package warehouse

import (
	"context"
	"order-service/package/structs"
)

type HTTPWarehouse interface {
	ReleaseStock(ctx context.Context, reqBody structs.RequestReleaseStock) error
}