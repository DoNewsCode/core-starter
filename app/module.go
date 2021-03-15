package app

import (
	"github.com/DoNewsCode/core-starter/app/commands"
	"github.com/DoNewsCode/core/contract"
	"github.com/DoNewsCode/core/di"
	"github.com/DoNewsCode/core/otgorm"
	"github.com/gorilla/mux"
	"github.com/DoNewsCode/core-starter/app/commands"
	"github.com/DoNewsCode/core-starter/app/entities"
	"github.com/DoNewsCode/core-starter/app/handlers"
	"github.com/DoNewsCode/core-starter/app/repositories"
	"github.com/spf13/cobra"
)

func New(config contract.ConfigAccessor, gin GinTransport) Module {
	return Module{
		config: config,
		gin:    gin,
	}
}

func Providers() di.Deps {
	return []interface{}{
		repositories.NewUserRepository,
		handlers.NewUserHandler,
		handlers.NewHandlers,
		NewGinTransport,
	}
}

type Module struct {
	config contract.ConfigAccessor
	gin    GinTransport
}

func (m Module) ProvideSeed() []*otgorm.Seed {
	return entities.Seeders()
}

func (m Module) ProvideMigration() []*otgorm.Migration {
	return entities.Migrations()
}

func (m Module) ProvideHTTP(router *mux.Router) {
	router.PathPrefix("/").Handler(m.gin)
}

func (m Module) ProvideCommand(command *cobra.Command) {
	command.AddCommand(
		commands.NewVersionCommand(m.config),
	)
}
