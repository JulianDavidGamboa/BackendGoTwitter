package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexión a la BD*/
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://zerocode:zerocode@dbtwitter.zuowu.mongodb.net/test")

/*ConectarBD es la funcion que me permite conectar a la BD*/
func ConectarBD() *mongo.Client {
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
	log.Println("Conexión exitosa con la BD")
	return client

}

/*ChequeoConnection es el ping a la BD*/
func ChequeoConnection() int {

	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
