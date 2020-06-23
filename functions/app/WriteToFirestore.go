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
	collName := app.Config.FirestoreCollection
	collRef := app.Firestore.Collection(collName)
	docRef := collRef.NewDoc()

	rowToWrite := row{
		contents:  contents,
		timestamp: time.Now().UnixNano() / 1000000,
	}

	_, err := docRef.Create(context, rowToWrite)
	if err != nil {
		return fmt.Errorf("Errr while writing to firestore (col=%s, doc=new()): %v", collName, err)
	}

	return nil
}
