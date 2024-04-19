package db

import (
	"database/sql"
	"fmt"
	"forum/lib/func/color"
	"forum/lib/func/json"
	"forum/lib/func/terminal"
	"forum/lib/types"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() {
	terminal.ClearTerminal()
	fmt.Println("Initialisation de la base de données...")
	if dbExists() {
		fmt.Println("La base de données existe déjà.")
	} else {
		fmt.Println("La base de données n'existe pas.")
		askDatabaseInfo()
	}
	config := json.ReadJsonConfig()
	err := connectDB(config)
	if err != nil {
		fmt.Println("Erreur lors de la connexion à la base de données:", err)
		return
	}
}

//////////////////////////////////////////////////////////////////////
// Vérifier si la base de données existe déjà ////////////////////////
//////////////////////////////////////////////////////////////////////

func dbExists() bool {
	// Vérifier si il y a un fichier config.json dans le dossier forum/db
	path := "db/config.json"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

//////////////////////////////////////////////////////////////////////
// Demander à l'utilisateur de créer la base de données //////////////
//////////////////////////////////////////////////////////////////////

func askDatabaseInfo() {
	terminal.ClearTerminal()
	g, b, _, d := color.GetTColors()
	fmt.Println(g, "Veuillez entrer les informations pour la base de données.")
	fmt.Println(g, "Host (exemple : ", b, "localhost - 10.6.0.161):", d)
	var host string
	fmt.Scanln(&host)
	fmt.Println(g, "Port (exemple : ", b, "3306):", d)
	var port string
	fmt.Scanln(&port)
	fmt.Println(g, "Nom de la base de données (exemple : ", b, "forum):", d)
	var dbname string
	fmt.Scanln(&dbname)
	fmt.Println(g, "Nom d'utilisateur (exemple : ", b, "root):", d)
	var user string
	fmt.Scanln(&user)
	fmt.Println(g, "Mot de passe (exemple : ", b, "password):", d)
	var password string
	fmt.Scanln(&password)
	terminal.ClearTerminal()

	//transformer port en int
	portInt, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("Erreur lors de la conversion du port en int:", err)
		return
	}

	config := types.MySQLConfig{
		User:     user,
		Password: password,
		Database: dbname,
		Host:     host,
		Port:     portInt,
	}

	// Créer le fichier config.json
	json.WriteJsonConfig(config, "db/config.json")
}

//////////////////////////////////////////////////////////////////////
// Créer la base de donnée ///////////////////////////////////////////
//////////////////////////////////////////////////////////////////////

var db *sql.DB

func connectDB(config types.MySQLConfig) error {
	createDatabasedbname(config.Database, config)
	return nil
}

func createDatabasedbname(dbname string, config types.MySQLConfig) error {
	// Connection parameters
	username := config.User
	password := config.Password
	hostname := fmt.Sprintf("%s:%d", config.Host, config.Port)

	// Chaîne de connexion à la base de données
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/", username, password, hostname)

	// Connexion à la base de données MySQL
	fmt.Println("Connexion à la base de données MySQL...")
	fmt.Println()
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}

	// Vérifier si la connexion à la base de données est réussie
	err = db.Ping()
	if err != nil {
		return err
	}
	fmt.Println("Connexion à la base de données réussie!")
	fmt.Println()

	// Créer la base de données
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbname)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Base de données créée")
	fmt.Println()

	// Créer les tables
	err = createTables(dbname)
	if err != nil {
		return err
	}
	time.Sleep(2 * time.Second)
	terminal.ClearTerminal()

	return nil
}

func createTables(database string) error {

	// Query to create the tables
	query := fmt.Sprintf("USE %s", database)

	// Execute the query
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Tabe utilisée :", database)
	fmt.Println()

	// Query to create the table 'users'
	// Query to create the table "users"
	query = `CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(50) NOT NULL UNIQUE,
		password TEXT NOT NULL,
		description TEXT,
		profile_picture BLOB(1000000),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Table 'users' créée")

	// Query to create the table "posts"
	query = `CREATE TABLE IF NOT EXISTS posts (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(50) NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		user_id INT,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Table 'posts' créée")

	// Query to create the table "comments"
	query = `CREATE TABLE IF NOT EXISTS comments (
		id INT AUTO_INCREMENT PRIMARY KEY,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		user_id INT,
		post_id INT,
		FOREIGN KEY (user_id) REFERENCES users(id),	
		FOREIGN KEY (post_id) REFERENCES posts(id)
	)`
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Table 'comments' créée")

	// Query to create the table "likes"
	query = `CREATE TABLE IF NOT EXISTS likes (
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT,
		post_id INT,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (post_id) REFERENCES posts(id)
	)`
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Table 'likes' créée")

	return nil
}
