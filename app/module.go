//go:generate trs proto/app.proto --doc=../docs/ --lib=../proto-lib --svcout=.
package app

import (
	"github.com/DoNewsCode/core-starter/app/commands"
	"github.com/DoNewsCode/core/contract"
	"github.com/DoNewsCode/core/di"
	"github.com/DoNewsCode/core/otgorm"
	"github.com/gorilla/mux"
	"github.com/nfangxu/core-skeleton/app/commands"
	"github.com/nfangxu/core-skeleton/app/entities"
	"github.com/nfangxu/core-skeleton/app/handlers"
	pb "github.com/nfangxu/core-skeleton/app/proto"
	"github.com/nfangxu/core-skeleton/app/repositories"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func New(config contract.ConfigAccessor, user UserTransport) Module {
	return Module{
		config: config,
		user:   user,
	}
}

func Providers() di.Deps {
	return []interface{}{
		repositories.NewUserRepository,
		handlers.NewUserHandler,
		NewUserTransport,
	}
}

type Module struct {
	config contract.ConfigAccessor
	user   UserTransport
}

func (m Module) ProvideGRPC(server *grpc.Server) {
	pb.RegisterAppServer(server, m.user)
}

func (m Module) ProvideSeed() []*otgorm.Seed {
	return entities.Seeders()
}

func (m Module) ProvideMigration() []*otgorm.Migration {
	return entities.Migrations()
}

func (m Module) ProvideHTTP(router *mux.Router) {
	router.PathPrefix("/").Handler(m.user)
}

func (m Module) ProvideCommand(command *cobra.Command) {
	command.AddCommand(
		commands.NewVersionCommand(m.config),
	)
}
