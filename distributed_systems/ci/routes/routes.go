package routes

import (
	"log/slog"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type Tree struct {
	logger *slog.Logger
}

func Register(api huma.API, logger *slog.Logger) {
	t := &Tree{logger: logger}

	huma.Register(api, huma.Operation{
		OperationID:   "health",
		Method:        http.MethodGet,
		Path:          "/health",
		Summary:       "Get health status",
		Description:   "Het health status of the current service and it's immediate dependencies.",
		Tags:          []string{"Health"},
		DefaultStatus: http.StatusOK,
	}, t.Health)
}
