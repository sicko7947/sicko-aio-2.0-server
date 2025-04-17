package cred

import (
	"crypto/tls"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// TLSInterceptor : TLS Certiciate Interceptor
func TLSInterceptor() grpc.ServerOption {
	// set credentials
	credentials.NewServerTLSFromCert(&tls.Certificate{})
	creds, err := credentials.NewServerTLSFromFile(`/root/auth/api.sickoaio.com.crt`, `/root/auth/api.sickoaio.com.key`)
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
	return grpc.Creds(creds)
}
