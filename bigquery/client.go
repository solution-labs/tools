package bigquery

import (
	"cloud.google.com/go/bigquery"
	"context"
	"errors"
	"fmt"
)

var ErrMissingProjectInformation = errors.New("missing project information")
var ErrMissingDatasetInformation = errors.New("missing dataset information")
var ErrMissingTableInformation = errors.New("missing table information")

func Client(ctx context.Context, ProjectID string) (client *bigquery.Client, err error) {

	if len(ProjectID) == 0 {
		return nil, ErrMissingProjectInformation
	}

	client, err = bigquery.NewClient(ctx, ProjectID)
	if err != nil {
		return nil, fmt.Errorf("bigquery:Client: %w", err)
	}

	return client, nil

}
