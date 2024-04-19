package server

import "fmt"

func Show(g, b, r, d, ip string) {
	// Affichez les adresses du site
	fmt.Println(g, "Le site est disponible aux adresses suivantes :")
	fmt.Println()
	fmt.Println(b, "Accueil :")
	fmt.Println(b, "     http://localhost:8080")
	fmt.Println(b, "     http://"+ip+":8080")
	fmt.Println()
	fmt.Println(r, "Appuyez sur Ctrl+C pour arrêter le serveur", d)
}
