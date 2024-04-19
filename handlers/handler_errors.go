package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

// Default est le gestionnaire pour la route /404

func Default(w http.ResponseWriter, r *http.Request) {
	connectFilePath := "templates/error/404.html"
	connectFile, err := template.ParseFiles(connectFilePath)

	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'ouverture du fichier HTML : %s", err), http.StatusInternalServerError)
		return
	}

	err = connectFile.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du modèle HTML : %s", err), http.StatusInternalServerError)
		return
	}
}

func RegError(w http.ResponseWriter, r *http.Request) {
	connectFilePath := "templates/error/register.html"
	connectFile, err := template.ParseFiles(connectFilePath)

	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'ouverture du fichier HTML : %s", err), http.StatusInternalServerError)
		return
	}

	err = connectFile.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du modèle HTML : %s", err), http.StatusInternalServerError)
		return
	}
}

func LogError(w http.ResponseWriter, r *http.Request) {
	connectFilePath := "templates/error/login.html"
	connectFile, err := template.ParseFiles(connectFilePath)

	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'ouverture du fichier HTML : %s", err), http.StatusInternalServerError)
		return
	}

	err = connectFile.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du modèle HTML : %s", err), http.StatusInternalServerError)
		return
	}
}
