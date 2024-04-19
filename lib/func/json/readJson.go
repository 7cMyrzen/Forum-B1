package json

import (
	"encoding/json"
	"forum/lib/types"
	"io/ioutil"
	"log"
)

func ReadJsonConfig() types.MySQLConfig {
	path := "db/config.json"

	// Lire le contenu du fichier
	configJSON, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier JSON: %v", err)
	}

	// Déclarer une variable pour stocker la configuration
	var config types.MySQLConfig

	// Décodez le JSON dans la structure de la configuration
	err = json.Unmarshal(configJSON, &config)
	if err != nil {
		log.Fatalf("Erreur lors du décodage du fichier JSON: %v", err)
	}

	// Renvoyer la configuration
	return config
}
