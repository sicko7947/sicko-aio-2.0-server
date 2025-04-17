package grpcHandler

import (
	"bytes"
	"encoding/json"
	"io"
	"regexp"

	"github.com/google/uuid"
	grpc_service "github.com/sicko7947/sicko-aio-backend/proto/rpc"
	"github.com/sicko7947/sicko-aio-backend/utils/psychoclient"
	"github.com/tidwall/gjson"
)

func (s *streamService) ProductMonitor(srv grpc_service.Stream_ProductMonitorServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		srv.Send(&grpc_service.StreamProductMonitorGraphqlResponse{
			Errors: &grpc_service.Errors{
				Code:    400,
				Message: "Method Clipped, Use V1 for monitoring",
			},
		})
		continue

		productID := req.GetProductId()
		proxy := req.GetProxy()

		// Validation Fields
		matchProxy, _ := regexp.MatchString("http://(.*):(.*)@(.*):(.*)", proxy)
		_, unMatchUUID := uuid.Parse(productID)
		switch {
		case unMatchUUID != nil, !matchProxy:
			srv.Send(&grpc_service.StreamProductMonitorGraphqlResponse{
				Errors: &grpc_service.Errors{
					Code:    400,
					Message: "Validation Failed",
				},
			})
			continue
		}

		// make a new response struct
		response := new(grpc_service.StreamProductMonitorGraphqlResponse)
		response.ProductId = req.GetProductId()
		response.SizeSkuMap = make(map[string]*grpc_service.SizeSkuMap)

		// Preparing Data for Request
		endpoint := "https://xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
		headers := map[string]string{}
		data, _ := json.Marshal(map[string]interface{}{})

		// Handling Requests
		res, respBody, e := psychoclient.NewClient(&psychoclient.SessionBuilder{
			Proxy: proxy,
		}).DoNewRequest(&psychoclient.RequestBuilder{
			Endpoint: endpoint,
			Method:   "POST",
			Headers:  headers,
			Payload:  bytes.NewBuffer(data),
		})

		if e != nil {
			response.Errors = &grpc_service.Errors{
				Code:    500,
				Message: "Error Requesting",
			}
			srv.Send(response)
			continue
		}

		switch res.StatusCode {
		case 200:
			result := gjson.Get(string(respBody), "data.product")
			content := result.Get("content")

			if content.Exists() {
				response.ProductId = result.Get("productId").String()
				response.QuantityLimit = result.Get("quantityLimit").Int()
				response.ProductDescription = content.Get("colorDescription").String()
				response.ProductName = content.Get("fullTitle").String()
			}

			availableSizeObj := result.Get("skus.#(available==true)#").Array()
			for _, v := range availableSizeObj {
				response.SizeSkuMap[v.Get("nikeSize").String()] = &grpc_service.SizeSkuMap{
					Gtin:  v.Get("gtin").String(),
					SkuId: v.Get("id").String(),
				}
			}
			srv.Send(response)
		default:
			response.Errors = &grpc_service.Errors{
				Code:    int64(res.StatusCode),
				Message: "Error Monitoring",
			}
			srv.Send(response)
		}
	}
}
