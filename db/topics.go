package db

import (
	"fmt"
	"forum/lib/func/forum"
	"forum/lib/func/image"
	"forum/lib/func/time"
	"forum/lib/types"
)

func CreateNewTopic(userID int, title, content string) error {
	created_at := time.GetCurrentTime()
	_, err := db.Exec("INSERT INTO POSTs(title, content, created_at, user_id) VALUES(?, ?, ?, ?)", title, content, created_at, userID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTopic(userID int, topicID string) error {
	DeleteAllCommentsForTopic(topicID)
	DeleteAllLikesForTopic(topicID)
	_, err := db.Exec("DELETE FROM POSTs WHERE id = ? AND user_id = ?", topicID, userID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllTopics() []types.Posts {
	// Exécutez une requête SQL avec une jointure pour récupérer les données des tables 'posts' et 'users'
	rows, err := db.Query("SELECT p.id, p.title, p.content, p.created_at, p.user_id, u.username, u.profile_picture FROM posts p JOIN users u ON p.user_id = u.id")
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête:", err)
		return nil
	}
	defer rows.Close()

	var posts []types.Posts
	// Parcourir les résultats et les afficher
	for rows.Next() {
		var id int
		var title string
		var content string
		var createdAt string
		var userID int
		var username string
		var profilePicture string

		err := rows.Scan(&id, &title, &content, &createdAt, &userID, &username, &profilePicture)
		if err != nil {
			fmt.Println("Erreur lors de la lecture de la ligne:", err)
			return nil
		}
		profilePicture = image.Base64ToImage(profilePicture)
		createdAt = time.FormatTime(createdAt)
		posts = append(posts, types.Posts{ID: id, Title: title, Content: content, Creation: createdAt, Author: username, AuthorPic: profilePicture})
	}

	for i, post := range posts {
		nbLikes := GetNbLikesForTopic(fmt.Sprintf("%d", post.ID))
		posts[i].NbLikes = nbLikes
		nbComs := GetNbCommentsForTopic(fmt.Sprintf("%d", post.ID))
		posts[i].NbComs = nbComs
	}

	// Vérifier s'il y a eu une erreur lors de l'itération sur les lignes
	if err = rows.Err(); err != nil {
		fmt.Println("Erreur lors de l'itération sur les résultats:", err)
		return nil
	}

	posts = forum.OrderTopicsNewestToOldest(posts)
	return posts
}

func CreateNewComment(content, createdAt string, userID int, topicID string) error {
	_, err := db.Exec("INSERT INTO comments(content, created_at, user_id, post_id) VALUES(?, ?, ?, ?)", content, createdAt, userID, topicID)
	if err != nil {
		fmt.Println("Erreur lors de l'insertion du commentaire:", err)
		return err
	}
	return nil
}

func GetInfoForTopic(topicID string) types.TopicPageData {
	Post := types.Posts{}
	Comments := []types.Comments{}

	// Récupérer les informations du sujet
	row := db.QueryRow("SELECT p.id, p.title, p.content, p.created_at, p.user_id, u.username, u.profile_picture FROM posts p JOIN users u ON p.user_id = u.id WHERE p.id = ?", topicID)
	var id int
	var title string
	var content string
	var createdAt string
	var userID int
	var username string
	var profilePicture string
	err := row.Scan(&id, &title, &content, &createdAt, &userID, &username, &profilePicture)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des informations du sujet:", err)
		return types.TopicPageData{}
	}
	profilePicture = image.Base64ToImage(profilePicture)
	createdAt = time.FormatTime(createdAt)
	var NbLikes = GetNbLikesForTopic(topicID)
	var NbComs = GetNbCommentsForTopic(topicID)

	Post = types.Posts{ID: id, Title: title, Content: content, Creation: createdAt, Author: username, AuthorPic: profilePicture, NbLikes: NbLikes, NbComs: NbComs}

	// Récupérer les commentaires du sujet avec username et profile_picture
	rows, err := db.Query("SELECT c.id, c.content, c.created_at, c.user_id, u.username, u.profile_picture FROM comments c JOIN users u ON c.user_id = u.id WHERE c.post_id = ?", topicID)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des commentaires:", err)
		return types.TopicPageData{}
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var content string
		var createdAt string
		var userID int
		var username string
		var profilePicture string
		err := rows.Scan(&id, &content, &createdAt, &userID, &username, &profilePicture)
		if err != nil {
			fmt.Println("Erreur lors de la lecture de la ligne:", err)
			return types.TopicPageData{}
		}
		profilePicture = image.Base64ToImage(profilePicture)
		createdAt = time.FormatTime(createdAt)
		Comments = append(Comments, types.Comments{ID: id, Content: content, Creation: createdAt, Author: username, AuthorPic: profilePicture})
	}

	// Vérifier s'il y a eu une erreur lors de l'itération sur les lignes
	if err = rows.Err(); err != nil {
		fmt.Println("Erreur lors de l'itération sur les résultats:", err)
		return types.TopicPageData{}
	}

	return types.TopicPageData{Post: Post, Comments: Comments}
}

func DeleteAllLikesForTopic(topicID string) error {
	_, err := db.Exec("DELETE FROM likes WHERE post_id = ?", topicID)
	if err != nil {
		fmt.Println("Erreur lors de la suppression des likes:", err)
		return err
	}
	return nil
}

func DeleteAllCommentsForTopic(topicID string) error {
	_, err := db.Exec("DELETE FROM comments WHERE post_id = ?", topicID)
	if err != nil {
		fmt.Println("Erreur lors de la suppression des commentaires:", err)
		return err
	}
	return nil
}

func GetNbLikesForTopic(topicID string) int {
	var nbLikes int
	row := db.QueryRow("SELECT COUNT(*) FROM likes WHERE post_id = ?", topicID)
	err := row.Scan(&nbLikes)
	if err != nil {
		fmt.Println("Erreur lors de la récupération du nombre de likes:", err)
		return 0
	}
	return nbLikes
}

func GetNbCommentsForTopic(topicID string) int {
	var nbComments int
	row := db.QueryRow("SELECT COUNT(*) FROM comments WHERE post_id = ?", topicID)
	err := row.Scan(&nbComments)
	if err != nil {
		fmt.Println("Erreur lors de la récupération du nombre de commentaires:", err)
		return 0
	}
	return nbComments
}
