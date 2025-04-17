package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sicko7947/sicko-aio-backend/grpcHandler"
	"github.com/sicko7947/sicko-aio-backend/utils"
)

// CheckExpireCookieInRedis : Check expired cookie loop 5 second each time
func CheckExpireCookieInRedis() {
	for {
		go utils.CheckExpireCookieInRedis()
		<-time.Tick(5 * time.Second)
	}
}

func main() {
	go CheckExpireCookieInRedis()
	port := fmt.Sprintf(":%s", os.Args[1])
	grpcHandler.StargrpcServer(port)
}
