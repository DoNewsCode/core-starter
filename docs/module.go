package docs

import (
	"embed"
	"github.com/gorilla/mux"
	"net/http"
)

//go:embed swagger docsify
var docs embed.FS

// New Create a new module
func New() Module {
	return Module{}
}

// Module is a main file
type Module struct {
}

// ProvideHTTP implements container.HTTPProvider
func (m Module) ProvideHTTP(router *mux.Router) {
	router.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.FS(docs))))
	router.PathPrefix("/docs").Handler(http.RedirectHandler("/docs/", 302))
}
