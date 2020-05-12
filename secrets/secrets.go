package secrets

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
)

// https://cloud.google.com/secret-manager/docs/quickstart

func CreateClient(ctx context.Context) (*secretmanager.Client, error) {
	return secretmanager.NewClient(ctx)
}
