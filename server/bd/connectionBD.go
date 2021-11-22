package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN is the connection object to the DB*/
var MongoCN = ConnectBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://memories:memories098@cluster0.qbm2k.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/*ConnectDB allows me to connect to the BD*/

func ConnectBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Success DB")
	return client
}

/*CheckConnection is ping*/
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
