package warehouse

import (
	"context"
	"net/http"
	"order-service/package/config"
	http_client "order-service/package/connection/http-client"
	"order-service/package/structs"
)
type HTTPWarehouseClient struct {
	c   http_client.HTTPClient
	cfg *config.Config
 }
 

func (h HTTPWarehouseClient) ReleaseStock(ctx context.Context, reqBody structs.RequestReleaseStock) error {
	// h.c = h.c.WithHeader([]http_client.HTTPHeader{
	// 	{
	// 		Key:   "Authorization",
	// 	},
	// })
	dest := map[string]interface{}{}
	_, err := h.c.
		PrepareRequestJSON(ctx, reqBody, http.MethodPost, "api/v1/stocks/release").
		Do(&dest)
	if err != nil {
		return err
	}
	return err
 
}


func NewHTTPWarehouse(c *config.Config) HTTPWarehouse {
	return HTTPWarehouseClient{
		c:   http_client.NewHTTPClient(http_client.HTTPClientWarehouse, c),
		cfg: c,
	}
 }
 