package app

import (
	"github.com/DoNewsCode/core/kitmw"
	"github.com/go-kit/kit/auth/jwt"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
	svc "github.com/nfangxu/core-skeleton/app/gen"
	"github.com/nfangxu/core-skeleton/app/handlers"
	pb "github.com/nfangxu/core-skeleton/app/proto"
	"net/http"
)

type UserTransport struct {
	http.Handler
	pb.AppServer
}

func NewUserTransport(server *handlers.UserHandler) UserTransport {
	endpoints := svc.NewEndpoints(server)
	endpoints.WrapAllExcept(kitmw.MakeValidationMiddleware())
	endpoints.WrapAllExcept(kitmw.MakeErrorConversionMiddleware(kitmw.ErrorOption{
		AlwaysHTTP200: true,
		ShouldRecover: false,
	}))
	httpHandler := svc.MakeHTTPHandler(endpoints,
		httptransport.ServerBefore(
			kitmw.IpToHTTPContext(),
		),
		httptransport.ServerErrorEncoder(kitmw.ErrorEncoder),
	)
	grpcHandler := svc.MakeGRPCServer(endpoints,
		grpctransport.ServerBefore(
			kitmw.IpToGRPCContext(),
		),
		grpctransport.ServerBefore(jwt.GRPCToContext()),
	)
	return UserTransport{
		Handler:   httpHandler,
		AppServer: grpcHandler,
	}
}
