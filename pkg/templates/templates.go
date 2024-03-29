package templates

import (
	"os"
	"path/filepath"
	log "github.com/projectdiscovery/gologger"
)

// Directory é o diretório padrão para buscar os arquivos de template
var Directory = filepath.Join(os.Getenv("HOME"), "certwatcher-templates", "templates")

// FindTemplateByID busca os templates com os IDs especificados em todas as pastas do diretório padrão e em quaisquer pastas adicionais especificadas, retornando os caminhos dos arquivos YAML correspondentes.
func Find(templateID []string, additionalDirs ...string) ([]string, error) {
	// Combine the default template directory with any additional directories to be searched
	dirs := append([]string{Directory}, additionalDirs...)

	// Create a map to store the paths of found template files
	templatePaths := make(map[string]string)

	// Search for template files in each specified directory
	for _, dir := range dirs {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Info().Msgf("Error searching for templates: %s", err)
			}
			if !info.IsDir() && filepath.Ext(path) == ".yaml" {
				// Get the filename without the extension
				filename := filepath.Base(path[:len(path)-len(filepath.Ext(path))])

				// Check if the filename matches one of the specified template IDs
				for _, id := range templateID {
					if filename == id {
						// If the file matches, store the full path to the file in the map
						templatePaths[id] = path
					}
				}
			}
			return nil
		})
		if err != nil {
			log.Info().Msgf("Error searching for templates: %s", err.Error())
		}
	}

	// Check if template files were found for each specified template ID
	missingTemplates := make([]string, 0)
	for _, id := range templateID {
		if _, ok := templatePaths[id]; !ok {
			missingTemplates = append(missingTemplates, id)
		}
	}

	if len(missingTemplates) > 0 {
		log.Fatal().Msgf("Templates with IDs %s not found", missingTemplates)
	}

	// Convert the map of file paths to a slice of file paths sorted by the specified template IDs
	sortedTemplatePaths := make([]string, len(templateID))
	for i, id := range templateID {
		sortedTemplatePaths[i] = templatePaths[id]
	}

	return sortedTemplatePaths, nil
}