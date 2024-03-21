package routes

import (
	//"exemple/controller"
	// "exemple/service"
	// temp "exemple/templates"
	controller "exemple/controller"
	"fmt"
	"net/http"
)

func SetRoutes() {

	http.HandleFunc("/artiste", controller.Artiste)
	http.HandleFunc("/ListAlbums", controller.ListAlbums)
	http.HandleFunc("/tracklist", controller.AlbumTracklist)
	http.HandleFunc("/index", controller.Index)
	http.HandleFunc("/catégorie", controller.Catégorie)
	http.HandleFunc("/collection", controller.Collection)
	http.HandleFunc("/detail", controller.Detail)
	http.HandleFunc("/favoris", controller.Favoris)
	http.HandleFunc("/journal", controller.Journal)
	http.HandleFunc("/recherche", controller.Recherche)

	fileserver := http.FileServer(http.Dir("rootDoc" + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("Le serveur est opérationel : http://localhost:8080/index")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
