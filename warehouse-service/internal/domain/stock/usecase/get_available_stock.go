package usecase

import (
	"context"
	"warehouse-service/package/structs"

	"github.com/spf13/cast"
)


func (u *UsecaseStock) GetAvailableStock(ctx context.Context, productID string) (structs.ResponseTotalStock, error){
	totalStock, err := u.repo.GetAvailableStockByProductId(ctx, cast.ToInt64(productID))
	if err != nil{
		return structs.ResponseTotalStock{}, nil
	}
	return structs.ResponseTotalStock{
		ProductID:  cast.ToInt64(productID),
		TotalStock: totalStock,
	}, nil
}