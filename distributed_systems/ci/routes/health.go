package routes

import "context"

type HealthOutput struct {
	Body []byte
}

type Health struct{}

func (h *Health) GetHealth(ctx context.Context, i *struct{}) (*HealthOutput, error) {
	resp := &HealthOutput{[]byte("Healthy")}
	return resp, nil
}
