package main

import (
	"context"
	"log"
	"time"

	"logger.svc/data"
)

type RPCServier struct{}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServier) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println("error writing to mongo: ", err.Error())
		return err 
	}

	*resp = "processed payload via RPC: " + payload.Name

	return nil
}
