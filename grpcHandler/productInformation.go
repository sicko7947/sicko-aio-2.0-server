package grpcHandler

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"strings"

	http "github.com/zMrKrabz/fhttp"

	"github.com/sicko7947/sicko-aio-backend/constants"
	grpc_service "github.com/sicko7947/sicko-aio-backend/proto/rpc"
	"github.com/sicko7947/sicko-aio-backend/redis"
	"github.com/sicko7947/sicko-aio-backend/utils/psychoclient"
	"github.com/tidwall/gjson"
)

type productInformationResponse struct {
	httpResponse *http.Response
	respBody     []byte
}

func (s *streamService) ProductInformation(srv grpc_service.Stream_ProductInformationServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// get request data
		styleColors := req.GetIds()
		language := req.GetLanguage()
		country := req.GetCountry()
		merchGroup := req.GetMerchGroup()
		proxies := req.GetProxy()

		// make a new response struct
		response := new(grpc_service.StreamProductInformationResponse)
		response.Objects = make(map[string]*grpc_service.StreamProductInformationResponseSkuData)
		channel := make(chan *productInformationResponse, len(styleColors))
		defer close(channel)

		switch {
		case len(styleColors) == 0, len(language) == 0, len(country) == 0, len(proxies) == 0:
			response.Errors = &grpc_service.Errors{
				Code:    400,
				Message: "Field Error",
			}
			srv.Send(response)
			continue
		}

		var skulist []string
		for _, styleColor := range styleColors {
			object, _ := redis.GetSkuData(fmt.Sprintf("%s.%s", styleColor, merchGroup))
			if object == nil || len(object.ProductId) == 0 {
				skulist = append(skulist, styleColor)
			} else {
				sizeSkus := make(map[string]*grpc_service.SizeSkuMap)
				for k, v := range object.SizeSkuMap {
					sizeSkus[k] = &grpc_service.SizeSkuMap{
						Gtin:  v.Gtin,
						SkuId: v.SkuId,
					}
				}

				obj := &grpc_service.StreamProductInformationResponseSkuData{
					ProductId:          object.ProductId,
					ProductName:        object.ProductName,
					ProductDescription: object.ProductDescription,
					Price:              object.Price,
					PublishType:        object.PublishType,
					LaunchId:           object.LaunchId,
					QuantityLimit:      int64(object.QuantityLimit),
					SizeSkus:           sizeSkus,
				}

				response.Objects[styleColor] = obj
			}
		}

		if len(skulist) == 0 {
			srv.Send(response)
			continue
		} else {
			headers := map[string]string{
				"accept":          "application/json",
				"content-type":    "application/json; charset=UTF-8",
				"accept-language": "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7,es;q=0.6",
				"cache-control":   "no-cache",
				"user-agent":      constants.GetRandomUserAgent(),
			}

			totalSkus := strings.Join(skulist[:], ",")
			count := int(math.Ceil(float64(len(totalSkus)) / 3000))
			for i := 0; i < count; i++ {
				index := rand.Intn(len(proxies))
				endpoint := fmt.Sprintf(`https://xxxxxxxxxxxxxxxx`)

				go func() {
					// Handling Requests
					res, respBody, err := psychoclient.NewClient(&psychoclient.SessionBuilder{
						Proxy: proxies[index],
					}).DoNewRequest(&psychoclient.RequestBuilder{
						Endpoint: endpoint,
						Method:   "GET",
						Headers:  headers,
						Payload:  nil,
					})

					if err != nil {
						return
					}
					channel <- &productInformationResponse{
						httpResponse: res,
						respBody:     respBody,
					}
				}()
			}

			requestCount := 0
		loop:
			for {
				select {
				case res := <-channel:
					requestCount++

					switch res.httpResponse.StatusCode {
					case 200:
						result := gjson.Parse(string(res.respBody))
						objects := result.Get("objects")
						objects.ForEach(func(key, value gjson.Result) bool {
							productInfo := gjson.Get(value.String(), "productInfo.0")

							status := productInfo.Get("merchProduct.status").String()                         // Get Product status
							publishType := productInfo.Get("merchProduct.publishType").String()               // Get Product Publish Type
							styleColor := productInfo.Get("merchProduct.styleColor").String()                 // Get style Color
							productName := productInfo.Get("productContent.fullTitle").String()               // Get product Name
							productDescription := productInfo.Get("productContent.colorDescription").String() // Get Product Description
							productID := productInfo.Get("merchProduct.id").String()                          // Get Product id
							quantityLimit := productInfo.Get("merchProduct.quantityLimit").Int()              // Get Cart Limit
							price := productInfo.Get("merchPrice.currentPrice").String()                      // Get Product Price
							currency := productInfo.Get("merchPrice.currenty").String()                       // Get Product Currenty
							launchID := productInfo.Get("launchView.id").String()                             // Get Launch id
							sizes := productInfo.Get("skus")                                                  // Get different Size Skus

							// Append gtins and size skus into object
							sizeSkus := make(map[string]*grpc_service.SizeSkuMap)
							sizes.ForEach(func(key, value gjson.Result) bool {
								nikeSize := gjson.Get(value.String(), "nikeSize").String() // Get Nike Size (US)
								skuId := gjson.Get(value.String(), "id").String()          // Get size sku
								gtin := gjson.Get(value.String(), "gtin").String()         // Get gtin

								sizeSkus[nikeSize] = &grpc_service.SizeSkuMap{
									Gtin:  gtin,
									SkuId: skuId,
								}
								return true
							})

							obj := &grpc_service.StreamProductInformationResponseSkuData{
								Status:             status,
								ProductId:          productID,
								ProductName:        productName,
								PublishType:        publishType,
								ProductDescription: productDescription,
								Price:              price + currency,
								QuantityLimit:      quantityLimit,
								LaunchId:           launchID,
								SizeSkus:           sizeSkus,
							}

							response.Objects[styleColor] = obj
							return true
						})
					default:
						break
					}
				}
				// Count request count
				if requestCount == count {
					break loop
				}
			}

			srv.Send(response)
		}
	}
}

// Differentiate slice into parts
func splitSlice(arr []string, num int64) [][]string {
	max := int64(len(arr))
	if max < num {
		return nil
	}
	var segmens = make([][]string, 0)
	quantity := max / num
	end := int64(0)
	for i := int64(1); i <= num; i++ {
		qu := i * quantity
		if i != num {
			segmens = append(segmens, arr[i-1+end:qu])
		} else {
			segmens = append(segmens, arr[i-1+end:])
		}
		end = qu - i
	}
	return segmens
}
