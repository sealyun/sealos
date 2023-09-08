package helper

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Init(c *gin.Context, client *mongo.Client) (*Request, error) {
	payRequest := &Request{}
	err := c.ShouldBindJSON(payRequest)
	if err != nil {
		return nil, fmt.Errorf("bind json error : %v", err)
	}
	// Identity authentication
	if err = Authenticate(payRequest, client); err != nil {
		return nil, fmt.Errorf("authenticate error : %v", err)
	}
	return payRequest, nil
}

func Authenticate(r *Request, client *mongo.Client) error {
	coll := InitDBAndColl(client, Database, AppColl)
	appID := r.AppID
	sign := r.Sign

	// create filter
	filter := bson.D{
		{"appID", appID},
		{"sign", sign},
	}

	var result bson.M
	if err := coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		// processing query error
		if err == mongo.ErrNoDocuments {
			// no matching document found, error returned
			return fmt.Errorf("no matching app found: %v", err)
		} else {
			// for other errors, print the error message and handle it
			return fmt.Errorf("unknown error: %v", err)
		}
	}
	return nil
}

func InitMongoClient(URI string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	return client
}

func InitDBAndColl(client *mongo.Client, datebase string, collection string) *mongo.Collection {
	coll := client.Database(datebase).Collection(collection)
	return coll
}
