package templates

import (
	"os"
	"path/filepath"
	log "github.com/sirupsen/logrus"
)

// DefaultTemplateDir é o diretório padrão para buscar os arquivos de template
var DefaultTemplateDir = filepath.Join(os.Getenv("HOME"), ".certwatcher-templates", "templates")

// FindTemplateByID busca os templates com os IDs especificados em todas as pastas do diretório padrão e retorna o caminho dos arquivos YAML correspondentes
func Find(templateID []string, additionalDirs ...string) ([]string, error) {
	// Combine the default template directory with any additional directories to search
	dirs := append([]string{DefaultTemplateDir}, additionalDirs...)

	// Search for the template files in each directory
	templatePaths := make([]string, 0)
	for _, dir := range dirs {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && filepath.Ext(path) == ".yaml" {
				// Get the filename without the extension
				filename := filepath.Base(path[:len(path)-len(filepath.Ext(path))])

				// Check if the filename matches any of the specified template IDs
				for _, templateID := range templateID {
					if filename == templateID {
						// If the file matches, add its path to the slice
						templatePaths = append(templatePaths, path)
					}
				}
			}
			return nil
		})
		if err != nil {
			log.Warnf("error searching for templates: %s", err.Error())
		}
	}

	// If no templates were found, return an error
	if len(templatePaths) == 0 {
		log.Warnf("templates with IDs '%v' not found", templateID)
	}

	return templatePaths, nil
}