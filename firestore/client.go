package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
)

func Client(ctx context.Context, projectID string) (client *firestore.Client, err error) {

	client, err = firestore.NewClient(ctx, projectID)

	if err != nil {
		return nil, err
	}

	return client, nil
}
