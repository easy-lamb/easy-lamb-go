package commands

import (
	"easy-lamb-go/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func CreateParseCommand() *cobra.Command {
	var configPath string
	var directory string

	var parseCmd = &cobra.Command{
		Use:   "parse",
		Short: "Parse the functions and create a Terraform variable file",
		Run: func(cmd *cobra.Command, args []string) {
			logrus.Infof("Parsing with config: %s in directory: %s", configPath, directory)
			os.Chdir(directory)
			err := services.ParseCommand(configPath)

			if err != nil {
				logrus.Fatalf("Error during parsing: %s", err)
			}
		},
	}

	parseCmd.Flags().StringVarP(&configPath, "config", "c", "easy-lamb.json", "Easy Lamb Go config file")
	parseCmd.Flags().StringVarP(&directory, "dir", "d", ".", "Project directory")

	return parseCmd
}
