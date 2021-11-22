package bd

import (
	"context"
	"time"

	"github.com/manuelnelson7/twittit/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModificoRegistro permite modificar el perfil del usuario */
func ModifyRegister(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittit")
	col := db.Collection("users")

	registro := make(map[string]interface{})
	if len(u.Name) > 0 {
		registro["nombre"] = u.Name
	}
	if len(u.LastName) > 0 {
		registro["apellidos"] = u.LastName
	}
	registro["fechaNacimiento"] = u.Birth
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Header) > 0 {
		registro["banner"] = u.Header
	}
	if len(u.Bio) > 0 {
		registro["biografia"] = u.Bio
	}
	if len(u.Location) > 0 {
		registro["ubicacion"] = u.Location
	}
	if len(u.Website) > 0 {
		registro["sitioWeb"] = u.Website
	}

	updtString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}