package handler

import (
	"context"
	"encoding/json"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	b, err := json.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
