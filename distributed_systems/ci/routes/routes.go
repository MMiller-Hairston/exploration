package routes

import (
	"log/slog"

	"github.com/danielgtaylor/huma/v2"
)

func Register(api huma.API, logger *slog.Logger) {
	// TODO: Move to Root Mux
	huma.Get(api, "/health", (&Health{}).GetHealth)
}
