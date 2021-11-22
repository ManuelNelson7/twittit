package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/manuelnelson7/twittit/middlew"
	"github.com/manuelnelson7/twittit/routers"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al Servidor */
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckBD(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.CheckBD(middlew.ValidoJWT(routers.SaveTweet))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.CheckBD(middlew.ValidoJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckBD(middlew.ValidoJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.CheckBD(middlew.ValidoJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlew.CheckBD(middlew.ValidoJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.CheckBD(middlew.ValidoJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.CheckBD(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.CheckBD(middlew.ValidoJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.CheckBD(routers.GetBanner)).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.CheckBD(middlew.ValidoJWT(routers.TopRelation))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.CheckBD(middlew.ValidoJWT(routers.LowRelation))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.CheckBD(middlew.ValidoJWT(routers.AskRelation))).Methods("GET")

	router.HandleFunc("/listaUsuarios", middlew.CheckBD(middlew.ValidoJWT(routers.UserList))).Methods("GET")
	router.HandleFunc("/leoTweetsSeguidores", middlew.CheckBD(middlew.ValidoJWT(routers.ReadTweetsFollowers))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}