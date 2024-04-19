package main

import (
	"forum/db"
	"forum/handlers"
	"forum/lib/func/color"
	"forum/lib/func/ip"
	"forum/lib/func/server"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDatabase()

	g, b, re, d := color.GetTColors()
	ip := ip.MyIP()

	// Initialiser le routeur Gorilla Mux
	r := mux.NewRouter()

	staticDir := "/static/"
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("static"))))

	// Gestionnaire pour la route d'accueil
	r.HandleFunc("/", handlers.Home)

	// Gestionnaire pour la route de connexion
	r.HandleFunc("/connect", handlers.Connect)
	r.HandleFunc("/register", handlers.Register)
	r.HandleFunc("/login", handlers.Login)
	r.HandleFunc("/logout", handlers.Logout)

	// Gestionnaire pour les pages de profil
	r.HandleFunc("/profile", handlers.Profile)
	r.HandleFunc("/update-profile", handlers.UpdateProfile)

	// Gestionnaire pour les pages de forum
	r.HandleFunc("/forum", handlers.Forum)
	r.HandleFunc("/new-topic", handlers.NewTopic)
	r.HandleFunc("/delete-topic", handlers.DeleteTopic)
	r.HandleFunc("/topic/{id}", handlers.TopicPage)
	r.HandleFunc("/new-com", handlers.NewComment)

	// Définir la gestion de la route par défaut pour les erreurs 404
	r.NotFoundHandler = http.HandlerFunc(handlers.Default)

	// Définir la route pour les autres erreurs
	r.HandleFunc("/register-error", handlers.RegError)
	r.HandleFunc("/login-error", handlers.LogError)

	// Démarrer le serveur
	server.Show(g, b, re, d, ip)
	http.ListenAndServe(":8080", r)
}
