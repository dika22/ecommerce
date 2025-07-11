package usecase

import (
	"context"
	"warehouse-service/package/structs"
)


func (u *UsecaseStock) BatchStock(ctx context.Context, req *structs.RequestBatchStock) (structs.ResponseBatchStock, error){
	productIDs := make([]int64, len(req.Items))
	for _, item := range req.Items {
		productIDs = append(productIDs, item.ProductID)
	}
	respStocks, err := u.repo.BatchStock(ctx, productIDs)
	if err != nil{
		return structs.ResponseBatchStock{}, nil
	}

	if respStocks == nil {
		return structs.ResponseBatchStock{
			AllAvailable: false,
			StockProducts: nil,
		}, nil
	}

	responseBatckStock := structs.ResponseBatchStock{}
	stockProducts := []structs.StockProduct{}
	for _, rs := range respStocks {
		stockProduct := structs.StockProduct{}
		if rs.Quantity > 0 {
			stockProduct.Available = true
			responseBatckStock.AllAvailable = true
		}
		stockProduct.ProductID = rs.ProductID
		stockProduct.Quantity = rs.Quantity
		stockProduct.WarehouseID = rs.WarehouseID
		stockProducts = append(stockProducts, stockProduct)
	}
	responseBatckStock.StockProducts = stockProducts
	return responseBatckStock, nil
}