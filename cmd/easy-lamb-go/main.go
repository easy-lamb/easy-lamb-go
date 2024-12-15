package main

import (
	"fmt"
	"github.com/easy-lamb/easy-lamb-go/cmd/easy-lamb-go/commands"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var logLevel string

func main() {
	// Root command
	var rootCmd = &cobra.Command{
		Use:   "easy-lamb-go",
		Short: "Easy Lambda Go",
		Long:  "Easy Lambda Go is a CLI tool to build and parse AWS Lambda Go functions to deploy them easily with Terraform",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Configurer le niveau de log en fonction du flag
			level, err := logrus.ParseLevel(logLevel)
			if err != nil {
				logrus.Fatalf("Invalid log level : %s\nAvailable levels : %s", logLevel, logrus.AllLevels)
			}
			logrus.SetLevel(level)
		},
	}

	logLevelUsage := fmt.Sprintf("Niveau de log (%s)", logrus.AllLevels)

	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "info", logLevelUsage)

	// Add subcommands
	rootCmd.AddCommand(commands.CreateParseCommand())
	rootCmd.AddCommand(commands.CreateBuildCommand())

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
}
