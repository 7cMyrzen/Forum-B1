package handlers

import (
	"encoding/json"
	"fmt"
	"forum/db"
	"forum/lib/func/gorilla"
	"forum/lib/func/time"
	"forum/lib/types"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func Forum(w http.ResponseWriter, r *http.Request) {
	session := gorilla.GetSession(w, r)
	authenticated := gorilla.Authenticated(w, r, session)
	forums := types.ForumPageData{
		Authenticated: authenticated,
	}

	posts := db.GetAllTopics()
	forums.Posts = posts

	forumFilePath := "templates/forum.html"
	forumFile, err := template.ParseFiles(forumFilePath)

	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'ouverture du fichier HTML : %s", err), http.StatusInternalServerError)
		return
	}

	err = forumFile.Execute(w, forums)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du modèle HTML : %s", err), http.StatusInternalServerError)
		return
	}
}

func NewTopic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Récupérer la session
	session := gorilla.GetSession(w, r)
	// Vérifier si l'utilisateur est connecté
	userID := gorilla.GetUserID(w, r, session)

	// Récupérer les données du formulaire
	title := r.FormValue("title")
	content := r.FormValue("content")

	// Créer un nouveau sujet
	err := db.CreateNewTopic(userID, title, content)
	if err != nil {
		http.Error(w, "Erreur lors de la création du nouveau sujet", http.StatusInternalServerError)
		return
	}

	// Redirection vers la page du forum
	http.Redirect(w, r, "/forum", http.StatusSeeOther)
}

func DeleteTopic(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Récupérer la session
	session := gorilla.GetSession(w, r)
	// Vérifier si l'utilisateur est connecté
	userID := gorilla.GetUserID(w, r, session)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		fmt.Println("Method not allowed")
		return
	}

	// Lire le corps de la requête
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		fmt.Println("Failed to read request body", err)
		return
	}

	// Décodez le corps de la requête en une structure DeletePostRequest
	var deleteRequest types.DeletePostRequest
	if err := json.Unmarshal(body, &deleteRequest); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		fmt.Println("Failed to parse request body", err)
		return
	}

	db.DeleteTopic(userID, deleteRequest.TopicID)
}

func TopicPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Récupérer l'ID du sujet
	vars := mux.Vars(r)
	topicID := vars["id"]

	topicData := db.GetInfoForTopic(topicID)

	topicFilePath := "templates/topic.html"
	topicFile, err := template.ParseFiles(topicFilePath)

	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'ouverture du fichier HTML : %s", err), http.StatusInternalServerError)
		return
	}

	err = topicFile.Execute(w, topicData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du modèle HTML : %s", err), http.StatusInternalServerError)
		return
	}

}

func NewComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Récupérer la session
	session := gorilla.GetSession(w, r)
	// Vérifier si l'utilisateur est connecté
	userID := gorilla.GetUserID(w, r, session)

	// Récupérer les données du formulaire
	content := r.FormValue("comment")
	topicID := r.FormValue("id")

	createdAt := time.GetCurrentTime()

	// Créer un nouveau commentaire
	err := db.CreateNewComment(content, createdAt, userID, topicID)
	if err != nil {
		http.Error(w, "Erreur lors de la création du nouveau commentaire", http.StatusInternalServerError)
		return
	}

	// Redirection vers la page du sujet
	http.Redirect(w, r, "/topic/"+topicID, http.StatusSeeOther)
}
