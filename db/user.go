package db

import (
	"forum/lib/func/forum"
	"forum/lib/func/image"
	"forum/lib/func/time"
	"forum/lib/types"
)

func GetProfileInfo(id int) types.ProfilePageData {
	GetDB()
	var profileData types.ProfilePageData

	// Récupérer les informations de base de l'utilisateur
	err := db.QueryRow("SELECT username, email, description, profile_picture, created_at FROM users WHERE id = ?", id).
		Scan(&profileData.Username, &profileData.Email, &profileData.Description, &profileData.Image, &profileData.Creation)
	if err != nil {
		// Gérer l'erreur de requête
		return types.ProfilePageData{}
	}

	// Convertir l'image de base64 à une image
	profileData.Image = image.Base64ToImage(profileData.Image)
	profileData.Creation = time.FormatTime(profileData.Creation)

	// Récupérer les posts de l'utilisateur
	rows, err := db.Query("SELECT id, title, created_at FROM posts WHERE user_id = ?", id)
	if err != nil {
		// Gérer l'erreur de requête
		return types.ProfilePageData{}
	}
	defer rows.Close()

	// Parcourir les résultats et ajouter chaque post à la liste de posts de l'utilisateur
	for rows.Next() {
		var post types.Posts
		err := rows.Scan(&post.ID, &post.Title, &post.Creation)
		if err != nil {
			// Gérer l'erreur de lecture de ligne
			return types.ProfilePageData{}
		}
		post.Creation = time.FormatTime(post.Creation)
		// Ajouter le post à la liste des posts de l'utilisateur
		profileData.Posts = append(profileData.Posts, post)
	}

	for i, post := range profileData.Posts {
		profileData.Posts[i].NbLikes = getNbLikes(post.ID)
		profileData.Posts[i].NbComs = getNbComments(post.ID)
	}

	if err := rows.Err(); err != nil {
		// Gérer l'erreur de parcours des résultats
		return types.ProfilePageData{}
	}
	// Classser les posts par date de création
	profileData.Posts = forum.OrderTopicsNewestToOldest(profileData.Posts)

	// Récupérer le nombre de posts de l'utilisateur
	var nbPosts int
	err = db.QueryRow("SELECT COUNT(*) FROM posts WHERE user_id = ?", id).Scan(&nbPosts)
	if err != nil {
		// Gérer l'erreur de requête
		return types.ProfilePageData{}
	}
	profileData.Nb_Posts = nbPosts

	// Récupérer le nombre de likes reçus par les posts de l'utilisateur
	var nbLikes = getAllLikesUserGot(id)
	profileData.Nb_Likes = nbLikes

	// Récupérer le nombre de commentaires reçus par les posts de l'utilisateur
	var nbComments = getAllCommentsUserGot(id)

	profileData.Nb_Comments = nbComments

	// Marquer l'utilisateur comme authentifié
	profileData.Authenticated = true

	// Retourner les informations du profil de l'utilisateur, y compris les posts et les nombres de posts, de likes et de commentaires
	return profileData
}

func UpdateProfile(id int, username string, email string, description string, base64IMG string) {
	GetDB()
	_, err := db.Exec("UPDATE users SET username = ?, email = ?, description = ?, profile_picture = ? WHERE id = ?", username, email, description, base64IMG, id)
	if err != nil {
		panic(err)
	}
}

func getAllLikesUserGot(id int) int {
	GetDB()
	var nbLikes int
	err := db.QueryRow("SELECT COUNT(*) FROM likes INNER JOIN posts ON likes.post_id = posts.id WHERE posts.user_id = ?", id).Scan(&nbLikes)
	if err != nil {
		// Gérer l'erreur de requête
		return 0
	}
	return nbLikes
}

func getAllCommentsUserGot(id int) int {
	GetDB()
	var nbComments int
	err := db.QueryRow("SELECT COUNT(*) FROM comments INNER JOIN posts ON comments.post_id = posts.id WHERE posts.user_id = ?", id).Scan(&nbComments)
	if err != nil {
		// Gérer l'erreur de requête
		return 0
	}
	return nbComments
}

func getNbLikes(postID int) int {
	GetDB()
	var nbLikes int
	err := db.QueryRow("SELECT COUNT(*) FROM likes WHERE post_id = ?", postID).Scan(&nbLikes)
	if err != nil {
		// Gérer l'erreur de requête
		return 0
	}
	return nbLikes
}

func getNbComments(postID int) int {
	GetDB()
	var nbComments int
	err := db.QueryRow("SELECT COUNT(*) FROM comments WHERE post_id = ?", postID).Scan(&nbComments)
	if err != nil {
		// Gérer l'erreur de requête
		return 0
	}
	return nbComments
}
