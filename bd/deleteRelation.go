package bd

import (
	"context"
	"time"

	"github.com/manuelnelson7/twittit/models"
)

/*BorroRelacion borra la relacion en la BD */
func DeleteRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittit")
	col := db.Collection("relation")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}