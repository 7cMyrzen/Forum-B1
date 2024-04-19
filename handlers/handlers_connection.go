package handlers

import (
	"fmt"
	"forum/db"
	"forum/lib/func/gorilla"
	"net/http"
	"text/template"
)

var store = gorilla.GetStore()

func Connect(w http.ResponseWriter, r *http.Request) {
	session := gorilla.GetSession(w, r)
	userID := session.Values["userID"]

	if userID != nil {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	} else {
		connectFilePath := "templates/connect.html"
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
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Récupérer les données du formulaire
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	password2 := r.FormValue("password2")

	// Vérifier si les mots de passe correspondent
	if password != password2 {
		http.Redirect(w, r, "/register-error", http.StatusSeeOther)
		return
	}

	// Vérifier si l'utilisateur existe déjà
	_, err := db.FindUser(username, password)
	if err == nil {
		http.Redirect(w, r, "/register-error", http.StatusSeeOther)
		return
	}

	// Ajouter l'utilisateur à la base de données
	id, err := db.AddUser(username, email, password)
	if err != nil {
		http.Redirect(w, r, "/register-error", http.StatusSeeOther)
		fmt.Println(err)
		return
	}

	// Créer une nouvelle session
	session := gorilla.CreateSession(r)

	// Stocker l'ID de l'utilisateur dans la session
	session.Values["userID"] = id

	// Enregistrer les changements dans la session
	gorilla.SaveSession(session, w, r)

	// Redirection vers une page de confirmation d'inscription
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Récupérer les données du formulaire
	email := r.FormValue("username")
	password := r.FormValue("password")

	// Chercher l'id de l'utilisateur dans la base de données
	id, err := db.FindUser(email, password)
	if err != nil {
		http.Redirect(w, r, "/login-error", http.StatusSeeOther)
		return
	}

	// Créer une nouvelle session
	session := gorilla.CreateSession(r)

	// Stocker l'ID de l'utilisateur dans la session
	session.Values["userID"] = id

	// Enregistrer les changements dans la session
	gorilla.SaveSession(session, w, r)

	// Redirection vers une page d'accueil après connexion réussie
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session := gorilla.GetSession(w, r)
	session.Values["userID"] = nil
	gorilla.SaveSession(session, w, r)
	http.Redirect(w, r, "/connect", http.StatusSeeOther)
}
