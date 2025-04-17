package redis

import (
	"time"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/util/gconv"
	"github.com/sicko7947/sicko-aio-backend/models"
)

func GetSkuData(keyName string) (result *models.NikeSku, err error) {
	conn := redisClient.Conn()
	defer conn.Close()

	res, err := conn.Do("HGETALL", keyName)
	if err != nil {
		return nil, err
	}

	temp := new(models.HashMapSkuData)
	gconv.Struct(res, &temp)

	result = &models.NikeSku{
		StyleColor:         temp.StyleColor,
		MerchGroup:         temp.MerchGroup,
		ProductName:        temp.ProductName,
		ProductDescription: temp.ProductDescription,
		ProductId:          temp.ProductId,
		LaunchId:           temp.LaunchId,
		Price:              temp.Price,
		CurrentPrice:       temp.CurrentPrice,
		PublishType:        temp.PublishType,
		QuantityLimit:      gconv.Int(temp.QuantityLimit),
		Status:             temp.Status,
		Exclusive:          temp.Exclusive,
		Discountability:    temp.Discountability,
	}
	result.CountryExclusion = gconv.SliceStr(temp.CountryExclusion)
	result.CommerceStartTime, _ = time.Parse("2006-01-02T15:04:05.000Z", temp.CommerceStartTime)
	if err = gjson.DecodeTo(temp.SizeSkuMap, &result.SizeSkuMap); err != nil {
		return nil, err
	}

	return result, nil
}
