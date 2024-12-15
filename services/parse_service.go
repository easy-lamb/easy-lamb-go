package services

import (
	"fmt"
	"github.com/easy-lamb/easy-lamb-go/models"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
)

func ParseCommand(configPath string) error {

	logrus.Infof("Parsing with config: %s", configPath)

	lambdaConfig, config, err := ParseFiles(configPath)

	if err != nil {
		logrus.Fatalf("Error during parsing: %s", err)
	}

	logrus.Infof("Writing terraform file ...")

	TerraformWriter(lambdaConfig, *config)

	return nil

}

func ParseFiles(configPath string) ([]map[string]string, *models.Config, error) {

	logrus.Info("Loading configuration ...")

	config, err := GetConfig(configPath)

	if err != nil {
		return nil, nil, err
	}

	logrus.Debug("Configuration loaded")
	logrus.Debugf("Configuration : %s", config)

	logrus.Info("Listing files ...")

	files, err := ListFiles(config.LambdaDir)

	if err != nil {
		return nil, nil, err
	}

	logrus.Debugf("Files listed (%d)", len(files))

	var lambdaConfigs = make([]map[string]string, 0)

	var mutex sync.Mutex
	var wg sync.WaitGroup

	errors := make(chan error, len(files))

	logrus.Info("Parsing files ...")

	for _, file := range files {
		wg.Add(1)
		go parseFile(file, config, &lambdaConfigs, &mutex, &wg, errors)
	}

	wg.Wait()
	close(errors)

	for err := range errors {
		if err != nil {
			logrus.Error("Error during parsing :", err)
		}
	}

	logrus.Infof("Generating terraform file ...")

	TerraformWriter(lambdaConfigs, config)

	return lambdaConfigs, &config, nil
}

func parseFile(file string, config models.Config, lambdaConfigs *[]map[string]string, mutex *sync.Mutex, wg *sync.WaitGroup, errors chan error) {
	defer wg.Done()

	logrus.Debugf("Parsing file %s", file)

	result, err := CollectComments(file)
	if err != nil {
		errors <- fmt.Errorf("Error during parsing %s : %w", file, err)
		return
	}

	if len(result) == 0 {
		logrus.Debugf("No comments found in %s", file)
		return
	}

	// Merge result with defaultParams
	for k, v := range config.DefaultParams {
		if _, ok := result[k]; !ok {
			result[k] = v
		}
	}

	// Get filepath without filename
	filePath := strings.Replace(file, config.LambdaDir+"/", "", 1)
	filePath = strings.Replace(filePath, "/main.go", "", 1)

	result["source"] = config.BuildOutput + "/" + filePath

	if _, ok := result["name"]; !ok {
		logrus.Debugf("No name found for %s, using filepath", file)
		result["name"] = strings.Replace(strings.ToLower(filePath), "/", "-", -1)
	}

	// Append to lambdaConfigs (protected by a mutex)
	mutex.Lock()
	*lambdaConfigs = append(*lambdaConfigs, result)
	mutex.Unlock()
}
