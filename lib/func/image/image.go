package image

import (
	"encoding/base64"
	"io/ioutil"
)

// ImageToBase64 convertit une image en base64
func ImageToBase64(imagepath string) (string, error) {
	// Lire le fichier d'image
	imgBytes, err := ioutil.ReadFile(imagepath)
	if err != nil {
		return "", err
	}

	// Encoder l'image en base64
	imgBase64 := base64.StdEncoding.EncodeToString(imgBytes)

	return imgBase64, nil
}

// Base64ToImage convertit une image base64 en image pour le html
func Base64ToImage(base64String string) string {
	return "data:image/png;base64," + base64String
}

func FilebytesToBase64(fileBytes []byte) string {
	return base64.StdEncoding.EncodeToString(fileBytes)
}
