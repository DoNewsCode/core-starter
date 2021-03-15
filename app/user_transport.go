package app

import (
	kitmw "github.com/DoNewsCode/core-kit/mw"
	kitOption "github.com/DoNewsCode/core-kit/option"
	svc "github.com/DoNewsCode/core-starter/app/gen"
	"github.com/DoNewsCode/core-starter/app/handlers"
	pb "github.com/DoNewsCode/core-starter/app/proto"
	"github.com/go-kit/kit/auth/jwt"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

type UserTransport struct {
	http.Handler
	pb.AppServer
}

func NewUserTransport(server *handlers.UserHandler) UserTransport {
	endpoints := svc.NewEndpoints(server)
	endpoints.WrapAllExcept(kitmw.Validate())
	endpoints.WrapAllExcept(kitmw.Error(kitmw.ErrorOption{
		AlwaysHTTP200: true,
		ShouldRecover: false,
	}))
	httpHandler := svc.MakeHTTPHandler(endpoints,
		httptransport.ServerBefore(
			kitOption.IPToHTTPContext(),
		),
		httptransport.ServerErrorEncoder(kitOption.ErrorEncoder),
	)
	grpcHandler := svc.MakeGRPCServer(endpoints,
		grpctransport.ServerBefore(
			kitOption.IPToGRPCContext(),
		),
		grpctransport.ServerBefore(jwt.GRPCToContext()),
	)
	return UserTransport{
		Handler:   httpHandler,
		AppServer: grpcHandler,
	}
}
