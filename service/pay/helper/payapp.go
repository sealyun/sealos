package helper

import (
	"context"
	"fmt"
	"os"

	"github.com/labring/sealos/service/pay/conf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	AppID      int64    `bson:"appID"`
	Sign       string   `bson:"sign"`
	PayAppName string   `bson:"payAppName"`
	Methods    []string `bson:"methods"`
}

func InsertApp(appID int64, sign, appName string, methods []string) (*mongo.InsertManyResult, error) {
	coll := InitDB(os.Getenv(conf.DBURI), conf.Database, conf.AppColl)
	docs := []interface{}{
		App{
			AppID:      appID,
			Sign:       sign,
			PayAppName: appName,
			Methods:    methods,
		},
	}

	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		//fmt.Println("insert the data of app failed:", err)
		return nil, fmt.Errorf("insert the data of app failed: %v", err)
	}
	fmt.Println("insert the data of app success:", result)
	return result, nil
}

// CheckAppAllowOrNot checks if the appID is allowed to use the payMethod
func CheckAppAllowOrNot(appID int64, payMethod string) error {
	coll := InitDB(os.Getenv(conf.DBURI), conf.Database, conf.AppColl)
	filter := bson.D{{"appID", appID}}
	var result bson.M
	if err := coll.FindOne(context.Background(), filter).Decode(&result); err != nil {
		fmt.Println("no allowed appID could be found:", err)
		return fmt.Errorf("no allowed appID could be found: %v", err)
	}

	methods := result["methods"].(bson.A)
	for _, method := range methods {
		if method == payMethod {
			return nil
		}
	}
	return fmt.Errorf("this payment method is not allowed in this app")
}

func CheckAppNameExistOrNot(appName string) error {
	coll := InitDB(os.Getenv(conf.DBURI), conf.Database, conf.AppColl)
	filter := bson.D{{"payAppName", appName}}

	var result bson.M
	err := coll.FindOne(context.Background(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// appName does not exist, return nil
		return nil
	} else if err != nil {
		// query error
		return fmt.Errorf("query error: %v", err)
	}

	// payAppName already exist
	return fmt.Errorf("app name already exists")
}
