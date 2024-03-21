package controller

import (
	service "exemple/service"
	"fmt"
	"net/http"
)

func ListAlbums(w http.ResponseWriter, r *http.Request) {
	ListAlbums, listAlbumsErr := service.GetAlbum("1cLLTkyM0B0v2vIG8wUhrD")
	if listAlbumsErr != nil {
		fmt.Println(listAlbumsErr.Error())
	}

	Temp.ExecuteTemplate(w, "ListAlbums", ListAlbums)
}

func AlbumTracklist(w http.ResponseWriter, r *http.Request) {
	Albumtracklist, AlbumtracklistErr := service.GetAlbumtracklist("1cLLTkyM0B0v2vIG8wUhrD")
	if AlbumtracklistErr != nil {
		fmt.Println(AlbumtracklistErr.Error())
	}

	Temp.ExecuteTemplate(w, "Albumtracklist", Albumtracklist)
}

func Artiste(w http.ResponseWriter, r *http.Request) {
	artiste, artisteErr := service.GetArtiste("5DW7NxqQ2875JtjuV7KMn4")
	if artisteErr != nil {
		fmt.Println(artisteErr.Error())
	}

	Temp.ExecuteTemplate(w, "artiste", artiste)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	detail, detailErr := service.GetDetail("5DW7NxqQ2875JtjuV7KMn4")
	if detailErr != nil {
		fmt.Println(detailErr.Error())
	}

	Temp.ExecuteTemplate(w, "detail", detail)
}

func Index(w http.ResponseWriter, r *http.Request) {
	index, indexErr := service.GetIndex("5DW7NxqQ2875JtjuV7KMn4")
	if indexErr != nil {
		fmt.Println(indexErr.Error())
	}

	Temp.ExecuteTemplate(w, "index", index)
}

func Catégorie(w http.ResponseWriter, r *http.Request) {
	catégorie, catégorieErr := service.GetCategoris("5DW7NxqQ2875JtjuV7KMn4")
	if catégorieErr != nil {
		fmt.Println(catégorieErr.Error())
	}

	Temp.ExecuteTemplate(w, "catégorie", catégorie)
}

func Collection(w http.ResponseWriter, r *http.Request) {
	collection, collectionErr := service.GetCollection("5DW7NxqQ2875JtjuV7KMn4")
	if collectionErr != nil {
		fmt.Println(collectionErr.Error())
	}

	Temp.ExecuteTemplate(w, "collection", collection)
}

func Favoris(w http.ResponseWriter, r *http.Request) {
	favoris, favorisErr := service.GetFavoris("5DW7NxqQ2875JtjuV7KMn4")
	if favorisErr != nil {
		fmt.Println(favorisErr.Error())
	}

	Temp.ExecuteTemplate(w, "favoris", favoris)
}

func Journal(w http.ResponseWriter, r *http.Request) {
	journal, journalErr := service.GetJournal("5DW7NxqQ2875JtjuV7KMn4")
	if journalErr != nil {
		fmt.Println(journalErr.Error())
	}

	Temp.ExecuteTemplate(w, "journal", journal)
}

func Recherche(w http.ResponseWriter, r *http.Request) {
	recherche, rechercheErr := service.GetRercherche("5DW7NxqQ2875JtjuV7KMn4")
	if rechercheErr != nil {
		fmt.Println(rechercheErr.Error())
	}

	Temp.ExecuteTemplate(w, "recherche", recherche)
}
