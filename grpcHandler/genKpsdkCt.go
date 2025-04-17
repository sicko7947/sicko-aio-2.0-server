package grpcHandler

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/gogf/gf/util/gconv"
	grpc_service "github.com/sicko7947/sicko-aio-backend/proto/rpc"
	"github.com/sicko7947/sicko-aio-backend/utils/psychoclient"
	"github.com/tidwall/gjson"
)

const getCtUrl = "http://localhost:8063/getct"

var headers = map[string]string{
	"accept":        "application/json",
	"content-type":  "application/json; charset=UTF-8",
	"cache-control": "no-cache",
}

func (s *streamService) GenKpsdkCt(srv grpc_service.Stream_GenKpsdkCtServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// get request data
		jobId := req.GetJobId()
		controlFlow := req.GetControlFlow()

		// make a new response struct
		response := &grpc_service.StreamKpsdkCtResponse{
			JobId: jobId,
		}

		// Preparing Data for Request
		data, _ := json.Marshal(map[string]string{
			"str": controlFlow,
		})

		reqId, _ := session.BuildRequest(&psychoclient.RequestBuilder{
			Endpoint: getCtUrl,
			Method:   "POST",
			Headers:  headers,
			Payload:  bytes.NewBuffer(data),
		})

		res, respBody, e := session.Do(reqId)
		if e != nil {
			response.Errors = &grpc_service.Errors{
				Code:    int64(e.Code),
				Message: e.Message,
			}
			srv.Send(response)
			continue
		}
		switch res.StatusCode {
		case 200:
			result := gjson.ParseBytes(respBody)
			kpsdkCt := result.Get("ct").String()
			tlarray := gconv.SliceInt64(result.Get("tlarray").Array())

			response.KpsdkCt = kpsdkCt
			response.TlArray = tlarray
		default:
			response.Errors = &grpc_service.Errors{
				Code:    int64(res.StatusCode),
				Message: "Error Getting tl array",
			}
		}
		srv.Send(response)
	}
}
