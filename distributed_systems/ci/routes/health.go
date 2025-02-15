package routes

import "context"

type HealthOutput struct {
	Body []byte
}

func (t *Tree) Health(ctx context.Context, i *struct{}) (*HealthOutput, error) {
	resp := &HealthOutput{[]byte("Healthy")}
	return resp, nil
}
