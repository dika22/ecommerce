package httpclient

import (
	"order-service/package/config"
	"order-service/package/http_client/warehouse"
)


type HTTPClients struct {
	WarehouseClient   warehouse.HTTPWarehouse
} 

func NewHTTPClients(c *config.Config) HTTPClients {
	return HTTPClients{
		WarehouseClient:   warehouse.NewHTTPWarehouse(c),
	}
 }
 