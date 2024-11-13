package main

import (
	"os"

	"github.com/spf13/cobra"
)

type band struct {
	Name  string   `json:"artist"`
	Songs []string `json:"songs"`
}

type music struct {
	Band []band `json:"Bands"`
}

var rootCmd = &cobra.Command{
	Use:          "music",
	Short:        "Music filtering tool",
	SilenceUsage: false,
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
