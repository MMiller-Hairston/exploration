package routes

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
)

// TODO: Add X-Hub-Signature-256 and X-Hub-Signature for security
type WebhooksInput struct {
	EventType string `header:"x-github-event"`
	EventId   string `header:"x-github-delivery"`
}

type WebhooksOutput struct {
	Status int
}

func Webhooks(ctx context.Context, i *WebhooksInput) (*struct{}, error) {
	if i.EventType == "" {
		return nil, huma.Error400BadRequest("No event type present")
	}

	// switch i.EventType {
	// // Ping to ensure webhook works
	// case "ping":
	// 	// Pull request events
	// case "pull_request":
	// case "commit_comment":
	// default:
	// 	return nil, huma.Error500InternalServerError("Unknown event type")
	// }
	return nil, nil
}
