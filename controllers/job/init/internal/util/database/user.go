package database

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"

	"github.com/labring/sealos/controllers/job/init/internal/util/controller"
)

func PresetAdminUser(ctx context.Context) error {
	//init mongodb database
	client, err := InitMongoDB(ctx)
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	cs, _ := connstring.ParseAndValidate(mongoUri)
	collection := client.Database(cs.Database).Collection(mongoUserCollection)

	user, err := newAdminUser()
	// check if the user already exists
	exist, err := user.IsExists(ctx, collection)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("admin user already exists")
	}

	// insert root user
	if _, err := collection.InsertOne(ctx, user); err != nil {
		return err
	}
	return nil
}

func newAdminUser() (*User, error) {
	hashedPassword, err := hashPassword(DefaultAdminPassword)
	if err != nil {
		return nil, err
	}
	return newUser(uuid.New().String(), DefaultAdminUserName, DefaultAdminUserName, hashedPassword, controller.DefaultAdminUserName), nil
}

func newUser(uid, name, passwordUser, hashedPassword, k8sUser string) *User {
	return &User{
		UID:          uid,
		Name:         name,
		PasswordUser: passwordUser,
		Password:     hashedPassword,
		// to iso string
		CreatedTime: time.Now().Format(time.RFC3339),
		K8sUsers: []K8sUser{
			{
				Name: k8sUser,
			},
		},
	}
}
