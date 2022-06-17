package app

import (
	"github.com/DoNewsCode/core-starter/app/commands"
	"github.com/DoNewsCode/core-starter/app/handlers"
	"github.com/DoNewsCode/core-starter/proto/users"
	"github.com/DoNewsCode/core/contract"
	"github.com/DoNewsCode/core/di"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
)

// New is the constructor of the app module.
func New(config contract.ConfigAccessor, usersSrv users.UsersServer) Module {
	return Module{config: config, usersSrv: usersSrv}
}

// Providers registers the dependencies provided by the app module (so that other modules can consume them).
func Providers() di.Deps {
	return di.Deps{
		handlers.NewUsers,
		di.Bind(new(*handlers.Users), new(users.UsersServer)),
	}
}

// Module is the struct representing the app Module.
type Module struct {
	config   contract.ConfigAccessor
	usersSrv users.UsersServer
}

// ProvideGRPC implements core.GRPCProvider
func (m Module) ProvideGRPC(server *grpc.Server) {
	users.RegisterUsersServer(server, m.usersSrv)
}

// ProvideHTTP implements core.HTTPProvider
func (m Module) ProvideHTTP(router *mux.Router) {
	users.RegisterUsersHTTPServer(
		http.NewServer(
			http.DisableEndpoint(),                                   // Added by nfangxu
			http.MuxRouter(router.PathPrefix("/users/").Subrouter()), // Added by nfangxu
		),
		m.usersSrv,
	)
}

// ProvideCommand implements core.CommandProvider
func (m Module) ProvideCommand(command *cobra.Command) {
	command.AddCommand(
		commands.NewExampleCommand(m.config),
	)
}
