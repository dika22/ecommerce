package usecase

import (
	"context"
	"product-service/package/structs"
)

func (u *ProductUsecase) CreateProduct(ctx context.Context, req *structs.RequestCreateProduct) error {
	// disini handle sample product change with product sample
	// with example 
	// logic bagaimana cara ganti velg bawaan ke velg yang di upload user
	// misal menggunakan mobil avanza menggunakan velg standar dan change menjadi velg premium
	// logic ini berlaku jika category id misal 1 itu product velg 
	// 1 : velg 2: tire 3: acc
	return nil
}