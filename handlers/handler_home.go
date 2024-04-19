package handlers

import (
	"fmt"
	"forum/lib/func/gorilla"
	"forum/lib/types"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// Récupérer la session
	session := gorilla.GetSession(w, r)
	// Vérifier si l'utilisateur est connecté
	authenticated := gorilla.Authenticated(w, r, session)

	// Données à passer au modèle HTML
	data := types.HomePageData{
		Authenticated: authenticated,
	}

	homeFilePath := "templates/home.html"
	homeFile, err := template.ParseFiles(homeFilePath)

	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'ouverture du fichier HTML : %s", err), http.StatusInternalServerError)
		return
	}

	err = homeFile.Execute(w, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du modèle HTML : %s", err), http.StatusInternalServerError)
		return
	}
}
