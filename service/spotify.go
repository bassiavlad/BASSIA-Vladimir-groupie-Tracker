package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var _ClientId string = "829cb5ac239f4a6ba89e213f865dd71c"
var _SecretClient string = "887ee2af67eb4bfa892cd1a2b8d4a9fd"
var _Token string

type ResToken struct {
	Token string `json:"access_token"`
}

// Cette méthode permet de récupérer un Token (il faudra initialiser les variables _ClientId & _SecretClient)
// Elle retourne une "erreur" lorsqu'une erreur se produit lors de la requête sinon elle retourne "nil"
func AskToken() error {
	httpClient := http.Client{Timeout: 2 * time.Second}
	// Initialisation de l'url
	urlToken := fmt.Sprintf("https://accounts.spotify.com/api/token?grant_type=client_credentials&client_id=%v&client_secret=%v", _ClientId, _SecretClient)
	// Initialisation de la requête avec la méthode, l'endpoint et le corps de la requête
	req, errReq := http.NewRequest(http.MethodPost, urlToken, nil)
	if errReq != nil {
		return fmt.Errorf("oupsss une erreur avec l'initialisation de la requete \"AskToken\"")
	}
	//
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Envoie de la requête
	res, resErr := httpClient.Do(req)
	if resErr != nil {
		return fmt.Errorf("oups une erreur avec le résultat de la requete : \"AskToken\"")
	}

	if res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		dataRes, dataResErr := io.ReadAll(res.Body)
		if dataResErr != nil {
			return fmt.Errorf("oups une erreur durant la lecture des données de la réponse : \"AskToken\"")
		}

		var data ResToken
		errDecode := json.Unmarshal(dataRes, &data)
		if errDecode != nil {
			return fmt.Errorf("oups une erreur lors du décodage des données : \"AskToken\"")
		}

		_Token = fmt.Sprintf("Bearer %s", data.Token)
		fmt.Println(_Token)
		return nil
	}

	return fmt.Errorf("oupss une erreur code réponse %v : \"AskToken\"", res.StatusCode)
}

type Album struct {
	Name   string `json:"name"`
	Label  string `json:"label"`
	Tracks struct {
		Items []struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"items"`
	} `json:"tracks"`
}

func GetAlbum(id string) (Album, error) {
	urlAPI := fmt.Sprintf("https://api.spotify.com/v1/albums/%v", id)
	AskToken()
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	// Initialisation de la requête avec la méthode, l'endpoint et le corps de la requête
	req, _ := http.NewRequest(http.MethodGet, urlAPI, nil)

	// Ajout du token dans le header avec l'attribut 'Authorization'
	req.Header.Add("Authorization", _Token)

	// Envoie de la requête
	res, resErr := client.Do(req)
	if resErr != nil {
		return Album{}, fmt.Errorf("oupss erreur avec la requete")
	}

	if res.StatusCode == http.StatusUnauthorized {
		return Album{}, fmt.Errorf("oupss erreur d'authentification : code 401")
		// Ici, il faudra gérer le token expiré...
	} else if res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		data, dataErr := io.ReadAll(res.Body)
		if dataErr != nil {
			fmt.Println("err lecture body")
			return Album{}, fmt.Errorf("oupss une erreur lors de lecture de la réponse")
		}

		var DataRes Album
		errJson := json.Unmarshal(data, &DataRes)
		if errJson != nil {
			fmt.Println(errJson)
			return Album{}, fmt.Errorf("oupss une erreur lors du decodage des données")
		}
		return DataRes, nil
	}

	return Album{}, fmt.Errorf("oupss une erreur avec la requete \" GetPlayer \" code réponse : %v", res.StatusCode)
}

type Albumtracklist struct {
	Name   string `json:"name"`
	Label  string `json:"label"`
	Tracks struct {
		Items []struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"items"`
	} `json:"tracks"`
}

func GetAlbumtracklist(id string) (Albumtracklist, error) {
	urlAPI := fmt.Sprintf("https://api.spotify.com/v1/Albumtracklist/%v", id)
	AskToken()
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	// Initialisation de la requête avec la méthode, l'endpoint et le corps de la requête
	req, _ := http.NewRequest(http.MethodGet, urlAPI, nil)

	// Ajout du token dans le header avec l'attribut 'Authorization'
	req.Header.Add("Authorization", _Token)

	// Envoie de la requête
	res, resErr := client.Do(req)
	if resErr != nil {
		return Albumtracklist{}, fmt.Errorf("oupss erreur avec la requete")
	}

	if res.StatusCode == http.StatusUnauthorized {
		return Albumtracklist{}, fmt.Errorf("oupss erreur d'authentification : code 401")
		// Ici, il faudra gérer le token expiré...
	} else if res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		data, dataErr := io.ReadAll(res.Body)
		if dataErr != nil {
			fmt.Println("err lecture body")
			return Albumtracklist{}, fmt.Errorf("oupss une erreur lors de lecture de la réponse")
		}

		var DataRes Albumtracklist
		errJson := json.Unmarshal(data, &DataRes)
		if errJson != nil {
			fmt.Println(errJson)
			return Albumtracklist{}, fmt.Errorf("oupss une erreur lors du decodage des données")
		}
		return DataRes, nil
	}

	return Albumtracklist{}, fmt.Errorf("oupss une erreur avec la requete \" GetPlayer \" code réponse : %v", res.StatusCode)
}

type Artiste struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Dateofbirth string `json:"dateofbirth"`
	Id          string `json:"id"`
	Nationality string `json:"nationality"`
}

func GetArtiste(id string) (Artiste, error) {
	urlAPI := fmt.Sprintf("https://api.spotify.com/v1/search?q=Youssoupha&type=artist/%v", id)

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	// Initialisation de la requête avec la méthode, l'endpoint et le corps de la requête
	req, _ := http.NewRequest(http.MethodGet, urlAPI, nil)

	// Ajout du token dans le header avec l'attribut 'Authorization'
	req.Header.Add("Authorization", _Token)

	// Envoie de la requête
	res, resErr := client.Do(req)
	if resErr != nil {
		return Artiste{}, fmt.Errorf("oupss erreur avec la requete")
	}

	if res.StatusCode == http.StatusUnauthorized {
		return Artiste{}, fmt.Errorf("oupss erreur d'authentification : code 401")
		// Ici, il faudra gérer le token expiré...
	} else if res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		data, dataErr := io.ReadAll(res.Body)
		if dataErr != nil {
			fmt.Println("err lecture body")
			return Artiste{}, fmt.Errorf("oupss une erreur lors de lecture de la réponse")
		}

		var DataRes Artiste
		errJson := json.Unmarshal(data, &DataRes)
		if errJson != nil {
			fmt.Println(errJson)
			return Artiste{}, fmt.Errorf("oupss une erreur lors du decodage des données")
		}
		return DataRes, nil
	}

	return Artiste{}, fmt.Errorf("oupss une erreur avec la requete \" GetPlayer \" code réponse : %v", res.StatusCode)
}

type Catégorie struct {
	Name        string `json:"Nom"`
	Description string `json:"Description"`
	Id          string `json:"id"`
	Type        string `json:"type"`
}

func GetCategoris(id string) (Catégorie, error) {
	urlAPI := fmt.Sprintf("https://api.spotify.com/v1/search?q=Youssoupha&type=artist/%v", id)

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	// Initialisation de la requête avec la méthode, l'endpoint et le corps de la requête
	req, _ := http.NewRequest(http.MethodGet, urlAPI, nil)

	// Ajout du token dans le header avec l'attribut 'Authorization'
	req.Header.Add("Authorization", _Token)

	// Envoie de la requête
	res, resErr := client.Do(req)
	if resErr != nil {
		return Catégorie{}, fmt.Errorf("oupss erreur avec la requete")
	}

	if res.StatusCode == http.StatusUnauthorized {
		return Catégorie{}, fmt.Errorf("oupss erreur d'authentification : code 401")
		// Ici, il faudra gérer le token expiré...
	} else if res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		data, dataErr := io.ReadAll(res.Body)
		if dataErr != nil {
			fmt.Println("err lecture body")
			return Catégorie{}, fmt.Errorf("oupss une erreur lors de lecture de la réponse")
		}

		var DataRes Catégorie
		errJson := json.Unmarshal(data, &DataRes)
		if errJson != nil {
			fmt.Println(errJson)
			return Catégorie{}, fmt.Errorf("oupss une erreur lors du decodage des données")
		}
		return DataRes, nil
	}

	return Catégorie{}, fmt.Errorf("oupss une erreur avec la requete \" GetPlayer \" code réponse : %v", res.StatusCode)
}

type Collection struct {
	Name        string `json:"Nom"`
	Description string `json:"Description"`
	Id          string `json:"id"`
	Type        string `json:"label"`
}

func GetCollection(id string) (Collection, error) {
	urlAPI := fmt.Sprintf("https://api.spotify.com/v1/search?q=Youssoupha&type=artist/%v", id)

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	// Initialisation de la requête avec la méthode, l'endpoint et le corps de la requête
	req, _ := http.NewRequest(http.MethodGet, urlAPI, nil)

	// Ajout du token dans le header avec l'attribut 'Authorization'
	req.Header.Add("Authorization", _Token)

	// Envoie de la requête
	res, resErr := client.Do(req)
	if resErr != nil {
		return Collection{}, fmt.Errorf("oupss erreur avec la requete")
	}

	if res.StatusCode == http.StatusUnauthorized {
		return Collection{}, fmt.Errorf("oupss erreur d'authentification : code 401")
		// Ici, il faudra gérer le token expiré...
	} else if res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		data, dataErr := io.ReadAll(res.Body)
		if dataErr != nil {
			fmt.Println("err lecture body")
			return Collection{}, fmt.Errorf("oupss une erreur lors de lecture de la réponse")
		}

		var DataRes Collection
		errJson := json.Unmarshal(data, &DataRes)
		if errJson != nil {
			fmt.Println(errJson)
			return Collection{}, fmt.Errorf("oupss une erreur lors du decodage des données")
		}
		return DataRes, nil
	}

	return Collection{}, fmt.Errorf("oupss une erreur avec la requete \" GetPlayer \" code réponse : %v", res.StatusCode)
}

type Journal struct {
	Name        string `json:"Nom"`
	Description string `json:"Description"`
	Id          string `json:"id"`
	Type        string `json:"label"`
}

func GetJournal(id string) (Journal, error) {
	urlAPI := fmt.Sprintf("https://api.spotify.com/v1/search?q=Youssoupha&type=artist/%v", id)

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	// Initialisation de la requête avec la méthode, l'endpoint et le corps de la requête
	req, _ := http.NewRequest(http.MethodGet, urlAPI, nil)

	// Ajout du token dans le header avec l'attribut 'Authorization'
	req.Header.Add("Authorization", _Token)

	// Envoie de la requête
	res, resErr := client.Do(req)
	if resErr != nil {
		return Journal{}, fmt.Errorf("oupss erreur avec la requete")
	}

	if res.StatusCode == http.StatusUnauthorized {
		return Journal{}, fmt.Errorf("oupss erreur d'authentification : code 401")
		// Ici, il faudra gérer le token expiré...
	} else if res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		data, dataErr := io.ReadAll(res.Body)
		if dataErr != nil {
			fmt.Println("err lecture body")
			return Journal{}, fmt.Errorf("oupss une erreur lors de lecture de la réponse")
		}

		var DataRes Journal
		errJson := json.Unmarshal(data, &DataRes)
		if errJson != nil {
			fmt.Println(errJson)
			return Journal{}, fmt.Errorf("oupss une erreur lors du decodage des données")
		}
		return DataRes, nil
	}

	return Journal{}, fmt.Errorf("oupss une erreur avec la requete \" GetPlayer \" code réponse : %v", res.StatusCode)
}

type Recherche struct {
	Name        string `json:"Nom"`
	Description string `json:"Description"`
	Id          string `json:"id"`
	Type        string `json:"label"`
}

func GetRercherche(id string) (Recherche, error) {
	urlAPI := fmt.Sprintf("https://api.spotify.com/v1/search?q=Youssoupha&type=artist/%v", id)

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	// Initialisation de la requête avec la méthode, l'endpoint et le corps de la requête
	req, _ := http.NewRequest(http.MethodGet, urlAPI, nil)

	// Ajout du token dans le header avec l'attribut 'Authorization'
	req.Header.Add("Authorization", _Token)

	// Envoie de la requête
	res, resErr := client.Do(req)
	if resErr != nil {
		return Recherche{}, fmt.Errorf("oupss erreur avec la requete")
	}

	if res.StatusCode == http.StatusUnauthorized {
		return Recherche{}, fmt.Errorf("oupss erreur d'authentification : code 401")
		// Ici, il faudra gérer le token expiré...
	} else if res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		data, dataErr := io.ReadAll(res.Body)
		if dataErr != nil {
			fmt.Println("err lecture body")
			return Recherche{}, fmt.Errorf("oupss une erreur lors de lecture de la réponse")
		}

		var DataRes Recherche
		errJson := json.Unmarshal(data, &DataRes)
		if errJson != nil {
			fmt.Println(errJson)
			return Recherche{}, fmt.Errorf("oupss une erreur lors du decodage des données")
		}
		return DataRes, nil
	}

	return Recherche{}, fmt.Errorf("oupss une erreur avec la requete \" GetPlayer \" code réponse : %v", res.StatusCode)
}

type Index struct {
	Name        string `json:"Nom"`
	Description string `json:"Description"`
	Id          string `json:"id"`
	Type        string `json:"label"`
}

func GetIndex(id string) (Index, error) {
	urlAPI := fmt.Sprintf("https://api.spotify.com/v1/search?q=Youssoupha&type=artist/%v", id)

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	// Initialisation de la requête avec la méthode, l'endpoint et le corps de la requête
	req, _ := http.NewRequest(http.MethodGet, urlAPI, nil)

	// Ajout du token dans le header avec l'attribut 'Authorization'
	req.Header.Add("Authorization", _Token)

	// Envoie de la requête
	res, resErr := client.Do(req)
	if resErr != nil {
		return Index{}, fmt.Errorf("oupss erreur avec la requete")
	}

	if res.StatusCode == http.StatusUnauthorized {
		return Index{}, fmt.Errorf("oupss erreur d'authentification : code 401")
		// Ici, il faudra gérer le token expiré...
	} else if res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		data, dataErr := io.ReadAll(res.Body)
		if dataErr != nil {
			fmt.Println("err lecture body")
			return Index{}, fmt.Errorf("oupss une erreur lors de lecture de la réponse")
		}

		var DataRes Index
		errJson := json.Unmarshal(data, &DataRes)
		if errJson != nil {
			fmt.Println(errJson)
			return Index{}, fmt.Errorf("oupss une erreur lors du decodage des données")
		}
		return DataRes, nil
	}

	return Index{}, fmt.Errorf("oupss une erreur avec la requete \" GetPlayer \" code réponse : %v", res.StatusCode)
}

type detail struct {
	Name        string `json:"Nom"`
	Description string `json:"Description"`
	Id          string `json:"id"`
	Type        string `json:"label"`
}

func GetDetail(id string) (detail, error) {
	urlAPI := fmt.Sprintf("https://api.spotify.com/v1/search?q=Youssoupha&type=artist/%v", id)

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	// Initialisation de la requête avec la méthode, l'endpoint et le corps de la requête
	req, _ := http.NewRequest(http.MethodGet, urlAPI, nil)

	// Ajout du token dans le header avec l'attribut 'Authorization'
	req.Header.Add("Authorization", _Token)

	// Envoie de la requête
	res, resErr := client.Do(req)
	if resErr != nil {
		return detail{}, fmt.Errorf("oupss erreur avec la requete")
	}

	if res.StatusCode == http.StatusUnauthorized {
		return detail{}, fmt.Errorf("oupss erreur d'authentification : code 401")
		// Ici, il faudra gérer le token expiré...
	} else if res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		data, dataErr := io.ReadAll(res.Body)
		if dataErr != nil {
			fmt.Println("err lecture body")
			return detail{}, fmt.Errorf("oupss une erreur lors de lecture de la réponse")
		}

		var DataRes detail
		errJson := json.Unmarshal(data, &DataRes)
		if errJson != nil {
			fmt.Println(errJson)
			return detail{}, fmt.Errorf("oupss une erreur lors du decodage des données")
		}
		return DataRes, nil
	}

	return detail{}, fmt.Errorf("oupss une erreur avec la requete \" GetPlayer \" code réponse : %v", res.StatusCode)
}

type Favoris struct {
	Name        string `json:"Nom"`
	Description string `json:"Description"`
	Id          string `json:"id"`
	Type        string `json:"label"`
}

func GetFavoris(id string) (Favoris, error) {
	urlAPI := fmt.Sprintf("https://api.spotify.com/v1/search?q=Youssoupha&type=artist/%v", id)

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	// Initialisation de la requête avec la méthode, l'endpoint et le corps de la requête
	req, _ := http.NewRequest(http.MethodGet, urlAPI, nil)

	// Ajout du token dans le header avec l'attribut 'Authorization'
	req.Header.Add("Authorization", _Token)

	// Envoie de la requête
	res, resErr := client.Do(req)
	if resErr != nil {
		return Favoris{}, fmt.Errorf("oupss erreur avec la requete")
	}

	if res.StatusCode == http.StatusUnauthorized {
		return Favoris{}, fmt.Errorf("oupss erreur d'authentification : code 401")
		// Ici, il faudra gérer le token expiré...
	} else if res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		data, dataErr := io.ReadAll(res.Body)
		if dataErr != nil {
			fmt.Println("err lecture body")
			return Favoris{}, fmt.Errorf("oupss une erreur lors de lecture de la réponse")
		}

		var DataRes Favoris
		errJson := json.Unmarshal(data, &DataRes)
		if errJson != nil {
			fmt.Println(errJson)
			return Favoris{}, fmt.Errorf("oupss une erreur lors du decodage des données")
		}
		return DataRes, nil
	}

	return Favoris{}, fmt.Errorf("oupss une erreur avec la requete \" GetPlayer \" code réponse : %v", res.StatusCode)
}
