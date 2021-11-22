package bd

import (
	"context"
	"time"

	"github.com/manuelnelson7/twittit/models"
)

/*InsertoRelacion graba la relación en la BD */
func InsertRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittit")
	col := db.Collection("relation")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}