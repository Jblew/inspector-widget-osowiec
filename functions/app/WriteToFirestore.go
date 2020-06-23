package app

import (
	"context"
	"fmt"
	"time"
)

type row struct {
	timestamp int64
	contents  string
}

// WriteToFirestore writes a text entry to firestore
func (app *App) WriteToFirestore(context context.Context, contents string) error {
	timestampMs := time.Now().UnixNano() / 1000000
	collName := app.Config.FirestoreCollection
	collRef := app.Firestore.Collection(collName)
	docRef := collRef.Doc(fmt.Sprintf("%d", timestampMs))

	rowToWrite := row{
		contents:  contents,
		timestamp: timestampMs,
	}

	_, err := docRef.Create(context, rowToWrite)
	if err != nil {
		return fmt.Errorf("Errr while writing to firestore (col=%s, doc=new()): %v", collName, err)
	}

	return nil
}
