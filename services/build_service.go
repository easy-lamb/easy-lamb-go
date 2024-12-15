package services

import (
	"fmt"
	"github.com/easy-lamb/easy-lamb-go/models"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func BuildCommand(configPath string, withParsing bool) error {
	lambdaConfig, config, err := ParseFiles(configPath)

	if err != nil {
		return err
	}

	if withParsing {
		logrus.Info("Writing Terraform file ...")
		TerraformWriter(lambdaConfig, *config)
	}

	logrus.Info("Building functions ...")

	var wg sync.WaitGroup

	errors := make(chan error, len(lambdaConfig))

	for _, lambda := range lambdaConfig {
		wg.Add(1)
		go BuildFile(lambda, *config, &wg, errors)
	}

	wg.Wait()
	close(errors)

	for err := range errors {
		if err != nil {
			fmt.Println("Erreur :", err)
		}
	}

	return nil
}

func BuildFile(lambdaConfig map[string]string, config models.Config, wg *sync.WaitGroup, errors chan error) {
	defer wg.Done()

	filePath := strings.Replace(lambdaConfig["source"], config.BuildOutput, config.LambdaDir, 1)

	logrus.Infof("Building %s", lambdaConfig["name"])

	cmd := exec.Command("go", "build", "-ldflags", "-w -s", "-o", lambdaConfig["source"]+"/bootstrap", filePath+"/main.go")
	cmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=amd64", "CGO_ENABLED=0")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		errors <- fmt.Errorf("Error during build %s : %w", filePath)
	}

	logrus.Infof("Function %s built", lambdaConfig["name"])
}
