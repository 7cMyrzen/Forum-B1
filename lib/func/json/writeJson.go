package json

import (
	"encoding/json"
	"forum/lib/types"
	"io/ioutil"
	"log"
)

func WriteJsonConfig(config types.MySQLConfig, path string) {
	// Convertir la configuration en JSON
	configJSON, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		log.Fatalf("Erreur lors de la conversion de la configuration en JSON: %v", err)
	}

	// Écrire le JSON dans le fichier
	err = ioutil.WriteFile(path, configJSON, 0644)
	if err != nil {
		log.Fatalf("Erreur lors de l'écriture du fichier JSON: %v", err)
	}
}
