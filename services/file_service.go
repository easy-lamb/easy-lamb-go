package services

import (
	"os"
	"path/filepath"
)

func ListFiles(directory string) ([]string, error) {
	// Liste pour stocker les fichiers trouvés
	var goFiles []string

	// Parcourir tous les fichiers et répertoires récursivement
	err := filepath.WalkDir(directory, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err // Gérer les erreurs lors du parcours
		}

		// Vérifier si le fichier a une extension ".go"
		if !d.IsDir() && filepath.Ext(path) == ".go" {
			goFiles = append(goFiles, path)
		}
		return nil
	})

	return goFiles, err
}
