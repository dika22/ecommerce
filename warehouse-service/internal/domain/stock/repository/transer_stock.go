package repository

import (
	"context"
	"errors"
	"warehouse-service/package/structs"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r StokRepository) TransferStockProduct(ctx context.Context, req *structs.RequestTransferStockProduct) error  {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		var fromStock, toStock structs.Stock
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("product_id = ? AND warehouse_id = ?", req.ProductID, req.FromWarehouseID).
			First(&fromStock).Error; err != nil {
			return err
		}

		if fromStock.Quantity < req.Quantity {
			return errors.New("Insufficient stock in source warehouse")
		}

		fromStock.Quantity -= req.Quantity
		if err := tx.Save(&fromStock).Error; err != nil {
			return err
		}

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("product_id = ? AND warehouse_id = ?", req.ProductID, req.ToWarehouseID).
			FirstOrCreate(&toStock, structs.Stock{ProductID: req.ProductID, WarehouseID: req.ToWarehouseID}).Error; err != nil {
			return err
		}

		toStock.Quantity += req.Quantity
		return tx.Save(&toStock).Error
	})

	if err != nil {
		return err
	}
	return nil
}