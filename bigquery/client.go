package bigquery

import (
	"cloud.google.com/go/bigquery"
	"context"
	"github.com/solution-labs/tools/toolserror"
)

func Client(ctx context.Context, ProjectID string) (client *bigquery.Client, err error) {

	if len(ProjectID) == 0 {
		return nil, toolserror.Wrap("bigquery:Client [Blank ProjectID]", err)
	}

	client, err = bigquery.NewClient(ctx, ProjectID)
	if err != nil {
		return nil, toolserror.Wrap("bigquery:Client: %w", err)
	}

	return client, nil

}
