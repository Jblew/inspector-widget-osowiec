package app

import (
	"context"
	"fmt"
	"time"
)

// WriteToFirestore writes a text entry to firestore
func (app *App) WriteToFirestore(context context.Context, contents string) error {
	timestampMs := time.Now().UnixNano() / 1000000
	collName := app.Config.FirestoreCollection
	collRef := app.Firestore.Collection(collName)
	docRef := collRef.NewDoc()

	doc := make(map[string]interface{})
	doc["contents"] = contents
	doc["timestamp"] = timestampMs
	_, err := docRef.Create(context, doc)
	if err != nil {
		return fmt.Errorf("Errr while writing to firestore (col=%s, doc=new()): %v", collName, err)
	}

	return nil
}
