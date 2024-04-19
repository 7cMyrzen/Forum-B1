package ip

import (
	"fmt"
	"net"
	"strings"
)

func GetIP() (string, error) {
	// Obtenez toutes les interfaces réseau de la machine
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("Erreur lors de la récupération des interfaces réseau: %v", err)
	}

	// Parcourez chaque interface pour obtenir les adresses IP
	for _, iface := range interfaces {
		// Vérifiez si l'interface est une carte Wi-Fi en vérifiant le nom
		if strings.Contains(iface.Name, "Wi-Fi") {
			addrs, err := iface.Addrs()
			if err != nil {
				return "", fmt.Errorf("Erreur lors de la récupération des adresses IP pour l'interface %s: %v", iface.Name, err)
			}

			// Parcourez chaque adresse IP associée à l'interface
			for _, addr := range addrs {
				// Vérifiez si l'adresse est de type IPv4
				ipNet, ok := addr.(*net.IPNet)
				if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					return ipNet.IP.String(), nil
				}
			}
		}
	}

	return "", fmt.Errorf("Aucune adresse IPv4 trouvée pour la carte Wi-Fi")
}

func MyIP() string {
	ip, err := GetIP()
	if err != nil {
		return err.Error()
	}
	return ip
}
