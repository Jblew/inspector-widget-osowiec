package app

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
)

// Init initializes the app
func Init(config Config) (*App, error) {
	ctx := context.Background()
	conf := &firebase.Config{}

	firebase, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, fmt.Errorf("Cannot initialize firebase.NewApp: %v", err)
	}

	firestore, err := firebase.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("Cannot initialize firebase.Firestore: %v", err)
	}

	app := &App{
		Firebase:  firebase,
		Firestore: firestore,
		Config:    config,
	}
	return app, nil
}
