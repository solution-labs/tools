package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/solution-labs/tools/toolserror"
)

// GetDocumentByID - Return document Snapshot
func GetDocumentByID(ctx context.Context, client *firestore.Client, Collection string, DocumentID string) (dsnap *firestore.DocumentSnapshot, err error) {

	dsnap, err = client.Collection(Collection).Doc(DocumentID).Get(ctx)
	if err != nil {
		return nil, toolserror.Wrap("firestore:GetDocumentByID", err)
	}

	return dsnap, nil
}

// SetDataByDocumentID - Save data to a collection
func SetDataByDocumentID(ctx context.Context, client *firestore.Client, Collection string, DocumentID string, data interface{}) (wr *firestore.WriteResult, err error) {
	wr, err = client.Collection(Collection).Doc(DocumentID).Set(ctx, data)
	if err != nil {
		return nil, toolserror.Wrap("firestore:SetDataByDocumentID", err)
	}

	return wr, nil
}
