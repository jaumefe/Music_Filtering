package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdDiff = &cobra.Command{
	Use:   "diff",
	Short: "Generates a json file with the difference between two directories",
	RunE:  cmdDiffRunE,
}

func init() {
	cmdDiff.Flags().StringP("src", "s", "Music", "Source music directory")
	cmdDiff.MarkFlagRequired("src")
	cmdDiff.Flags().StringP("dst", "d", "", "Destination folder")
	cmdDiff.MarkFlagRequired("dst")

	rootCmd.AddCommand(cmdDiff)
}

func cmdDiffRunE(cmd *cobra.Command, args []string) error {
	srcDirName, err := cmd.Flags().GetString("src")
	if err != nil {
		return err
	}

	dstDirName, err := cmd.Flags().GetString("dst")
	if err != nil {
		return err
	}

	srcMusic := music{Band: make([]band, 0)}
	err = srcMusic.getSongsListFolder(srcDirName)
	if err != nil {
		return err
	}

	dstMusic := music{Band: make([]band, 0)}
	err = dstMusic.getSongsListFolder(dstDirName)
	if err != nil {
		return err
	}

	fmt.Println(srcMusic.Band[0])
	fmt.Println(dstMusic.Band[0])

	return nil
}
