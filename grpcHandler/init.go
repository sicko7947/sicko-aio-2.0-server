package grpcHandler

import (
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/shimingyah/pool"
	"github.com/sicko7947/sicko-aio-backend/grpcHandler/middleware/cred"
	"github.com/sicko7947/sicko-aio-backend/grpcHandler/middleware/recovery"
	"github.com/sicko7947/sicko-aio-backend/grpcHandler/middleware/zap"
	grpc_service "github.com/sicko7947/sicko-aio-backend/proto/rpc"
	"github.com/sicko7947/sicko-aio-backend/utils/psychoclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// streamService struct hold all stream service method
type streamService struct{}

// GrpcServer
var GrpcServer *grpc.Server

// session does all the interal requests to kpsdk api ran on nodejs
var session psychoclient.Session

func init() {
	session, _ = psychoclient.NewSession(&psychoclient.SessionBuilder{
		UseDefaultClient: true,
	})
}

func StargrpcServer(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	// Create a new grpc server
	GrpcServer = grpc.NewServer(
		cred.TLSInterceptor(),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zap.ZapInterceptor()),
			// grpc_auth.StreamServerInterceptor(auth.AuthInterceptor),
			grpc_recovery.StreamServerInterceptor(recovery.RecoveryInterceptor()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zap.ZapInterceptor()),
			// grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor),
			grpc_recovery.UnaryServerInterceptor(recovery.RecoveryInterceptor()),
		)),

		// sete grpc connection pool configuration
		grpc.InitialWindowSize(pool.InitialWindowSize),
		grpc.InitialConnWindowSize(pool.InitialConnWindowSize),
		grpc.MaxSendMsgSize(pool.MaxSendMsgSize),
		grpc.MaxRecvMsgSize(pool.MaxRecvMsgSize),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			PermitWithoutStream: true,
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    pool.KeepAliveTime,
			Timeout: pool.KeepAliveTimeout,
		}),
	)
	grpc_service.RegisterStreamServer(GrpcServer, &streamService{})
	fmt.Println(port + " HTTP.Listing whth TLS and token...")
	err = GrpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
