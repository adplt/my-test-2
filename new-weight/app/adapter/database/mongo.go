package database

import (
	"context"
	"github.com/gomodul/envy"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoDB struct {
	*mongo.Client
}

func SetupMongoDB() *MongoDB {
	var (
		opt *options.ClientOptions
		host = envy.Get("SCIF_HOST")
		port = envy.Get("SCIF_PORT")
		user = envy.Get("SCIF_USERNAME")
		pass = envy.Get("SCIF_PASSWORD")
	)

	if user != "" && pass != "" {
		opt = options.Client().ApplyURI("mongodb://" + host + ":" + port).
			SetAuth(options.Credential{
				Username: user,
				Password: pass,
			})
	} else {
		opt = options.Client().ApplyURI("mongodb://" + host + ":" + port)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	client, err := mongo.NewClient(opt)
	if err != nil {
		panic(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return &MongoDB{client}
}