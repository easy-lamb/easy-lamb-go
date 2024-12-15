package commands

import (
	"github.com/easy-lamb/easy-lamb-go/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func CreateBuildCommand() *cobra.Command {
	var configPath string
	var directory string
	var withParsing bool

	var buildCmd = &cobra.Command{
		Use:   "build",
		Short: "Build the functions and create a Terraform variable file",
		Run: func(cmd *cobra.Command, args []string) {
			logrus.Infof("Build with config: %s in directory: %s", configPath, directory)
			os.Chdir(directory)
			err := services.BuildCommand(configPath, withParsing)

			if err != nil {
				logrus.Fatalf("Error during parsing: %s", err)
			}
		},
	}

	buildCmd.Flags().StringVarP(&configPath, "config", "c", "easy-lamb.json", "Easy Lamb Go config file")
	buildCmd.Flags().StringVarP(&directory, "dir", "d", ".", "Project directory")
	buildCmd.Flags().BoolVarP(&withParsing, "with-parsing", "n", true, "Generate the Terraform file during build")

	return buildCmd
}
