package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/solution-labs/tools/toolserror"
)

func Client(ctx context.Context, projectID string) (client *firestore.Client, err error) {

	client, err = firestore.NewClient(ctx, projectID)

	if err != nil {
		return nil, toolserror.Wrap("firestore:Client", err)
	}

	return client, nil
}
