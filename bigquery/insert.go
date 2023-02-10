package bigquery

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
	"github.com/solution-labs/tools/toolserror"
)

func Insert(ctx context.Context, client *bigquery.Client, dataSet string, table string, bi interface{}) error {

	if len(dataSet) == 0 {
		return toolserror.Wrap("bigquery:Insert", fmt.Errorf("blank dataSet variable"))
	}

	if len(table) == 0 {
		return toolserror.Wrap("bigquery:Insert", fmt.Errorf("blank table variable"))
	}

	inserter := client.Dataset(dataSet).Table(dataSet).Inserter()

	var errors error

	if err := inserter.Put(ctx, bi); err != nil {
		if multiError, ok := err.(bigquery.PutMultiError); ok {
			for _, err1 := range multiError {
				for _, err2 := range err1.Errors {
					errors = fmt.Errorf("%w", err2)
				}
			}
		} else {
			errors = fmt.Errorf("%w", fmt.Errorf("unspecified error - retry save of data"))
		}
	}

	if errors != nil {
		return toolserror.Wrap("bigquery:Insert.Put", errors)
	}

	return nil

}
