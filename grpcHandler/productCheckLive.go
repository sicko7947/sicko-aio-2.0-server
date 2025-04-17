package grpcHandler

import (
	"fmt"
	"io"
	"regexp"

	"github.com/google/uuid"
	grpc_service "github.com/sicko7947/sicko-aio-backend/proto/rpc"
	"github.com/sicko7947/sicko-aio-backend/utils/psychoclient"
	"github.com/sicko7947/sickocommon"
)

func (s *streamService) ProductCheckLive(srv grpc_service.Stream_ProductCheckLiveServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// get request data
		productID := req.GetProductId()
		proxy := req.GetProxy()

		// Validation Fields
		matchProxy, _ := regexp.MatchString("http://(.*):(.*)@(.*):(.*)", proxy)
		_, unMatchUUID := uuid.Parse(productID)
		switch {
		case unMatchUUID != nil, !matchProxy:
			srv.Send(&grpc_service.StreamProductCheckLiveResponse{
				Errors: &grpc_service.Errors{
					Code:    400,
					Message: "Validation Failed",
				},
			})
			continue
		}

		// make a new response struct
		response := new(grpc_service.StreamProductCheckLiveResponse)

		endpoint := fmt.Sprintf("https://xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx%s?%s", productID, sickocommon.RandStringWithSetLength(20))
		headers := map[string]string{}

		// Handling Requests
		res, _, e := psychoclient.NewClient(&psychoclient.SessionBuilder{
			Proxy: proxy,
		}).DoNewRequest(&psychoclient.RequestBuilder{
			Endpoint: endpoint,
			Method:   "GET",
			Headers:  headers,
			Payload:  nil,
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
		case 404:
			response.Live = false
			srv.Send(response)
		case 200:

			srv.Send(response)
		default:
			response.Errors = &grpc_service.Errors{
				Code:    int64(res.StatusCode),
				Message: "Error Checking Product Live",
			}
			srv.Send(response)
		}
	}
}
