package services

import (
	"github.com/easy-lamb/easy-lamb-go/models"
	"os"
)

func TerraformWriter(sources []map[string]string, config models.Config) {
	fileContent := "functions = ["

	for _, source := range sources {
		fileContent += "{"
		for k, v := range source {
			fileContent += "\"" + k + "\": \"" + v + "\","
		}
		fileContent += "},"
	}

	fileContent += "]"

	err := os.WriteFile(config.TerraformDir+"/"+config.TerraformFilename, []byte(fileContent), 0644)

	if err != nil {
		panic(err)
	}

}
