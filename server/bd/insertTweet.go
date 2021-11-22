package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/manuelnelson7/twittit/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*InsertoTweet graba el Tweet en la BD */
func InsertTweet(t models.SaveTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittit")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid":  t.UserID,
		"message": t.Mensaje,
		"date":   t.Fecha,
	}
	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}