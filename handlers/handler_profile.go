package handlers

import (
	"fmt"
	"forum/db"
	"forum/lib/func/gorilla"
	"forum/lib/func/image"
	"forum/lib/types"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	// Récupérer la session
	session := gorilla.GetSession(w, r)

	// Vérifier si l'utilisateur est connecté en vérifiant la présence de l'ID de l'utilisateur dans la session
	userID := gorilla.GetUserID(w, r, session)

	// Faites ce que vous avez à faire avec l'ID de l'utilisateur, comme par exemple, afficher son profil, etc.
	if userID == 0 {
		http.Redirect(w, r, "/connect", http.StatusSeeOther)
		return
	} else {
		var data types.ProfilePageData
		data = db.GetProfileInfo(userID)
		data.Authenticated = true

		templateFilePath := "templates/profile.html"
		tmpl, err := template.ParseFiles(templateFilePath)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erreur lors de l'ouverture du fichier HTML : %s", err), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du modèle HTML : %s", err), http.StatusInternalServerError)
			println(err)
			return
		}
	}
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	// Récupérer la session
	session := gorilla.GetSession(w, r)

	// Vérifier si l'utilisateur est connecté en vérifiant la présence de l'ID de l'utilisateur dans la session
	userID := gorilla.GetUserID(w, r, session)

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Récupération des données du formulaire
	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	desc := r.FormValue("desc")
	imageFile, _, err := r.FormFile("image")
	if err != nil && err != http.ErrMissingFile { // Vérifier si une erreur différente de ErrMissingFile est retournée
		http.Error(w, "Unable to get image", http.StatusBadRequest)
		return
	}
	defer func() {
		if imageFile != nil {
			imageFile.Close()
		}
	}()

	var base64Image string

	if imageFile != nil { // Si une nouvelle image est téléchargée
		imageBytes, err := ioutil.ReadAll(imageFile)
		if err != nil {
			http.Error(w, "Unable to read image", http.StatusInternalServerError)
			return
		}
		base64Image = image.FilebytesToBase64(imageBytes)
	} else { // Si aucune nouvelle image n'est téléchargée, récupérer l'image actuelle de la base de données
		currentProfile := db.GetProfileInfo(userID)
		base64Image = currentProfile.Image // Suppose que GetProfileInfo renvoie une structure avec le champ Image
		// Enlever le préfixe s'il existe
		if strings.HasPrefix(base64Image, "data:image/png;base64,") {
			base64Image = strings.TrimPrefix(base64Image, "data:image/png;base64,")
		}
	}

	// Faites quelque chose avec les données récupérées, comme les enregistrer dans une base de données ou les traiter d'une autre manière
	db.UpdateProfile(userID, username, email, desc, base64Image)

	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}
