package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/manuelnelson7/twittit/models"
)

/*GeneroJWT genera el encriptado con JWT */
func GeneroJWT(t models.User) (string, error) {

	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Name,
		"apellidos":        t.LastName,
		"fecha_nacimiento": t.Birth,
		"biografia":        t.Bio,
		"ubicacion":        t.Location,
		"sitioweb":         t.Website,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}