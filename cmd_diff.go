package main

import "github.com/spf13/cobra"

var cmdDiff = &cobra.Command{
	Use:   "diff",
	Short: "Generates a json file with the difference between two directories",
	RunE:  cmdDiffRunE,
}

func init() {
	cmdDiff.Flags().StringP("dst", "d", "", "Destination folder")
	cmdDiff.MarkFlagRequired("dst")

	rootCmd.AddCommand(cmdDiff)
}

func cmdDiffRunE(cmd *cobra.Command, args []string) error {
	return nil
}
