package handlers

import (
	"net/http"
	"text/template"
	"time"
)

func LoadingPageMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Charger le template de page de chargement
		tmpl, err := template.ParseFiles("templates/loading/load.html")
		if err != nil {
			http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
			return
		}

		// Afficher la page de chargement
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Erreur d'affichage du template", http.StatusInternalServerError)
			return
		}

		// Simuler un délai pour l'effet de chargement
		time.Sleep(2 * time.Second)

		// Passer à la page suivante
		next.ServeHTTP(w, r)
	})
}
