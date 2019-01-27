package cmd

import (
	"fmt"
	"os"

	docker "docker.io/go-docker"
	"github.com/leopardslab/Dunner/internal/logger"
	"github.com/spf13/cobra"
)

var log = logger.Log

var rootCmd = &cobra.Command{
	Use:   "dunner",
	Short: "Dunner is a Docker based task-runner",
	Long:  `You can define a set of commands and on what Docker images these commands should run as steps. A task has many steps. Then you can run these tasks with 'dunner do nameoftask'`,
	Run: func(cmd *cobra.Command, args []string) {

		_, err := docker.NewEnvClient()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Dunner running!")
	},
}

// Execute method executes the 'Run' method of rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
