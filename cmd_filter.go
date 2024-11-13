package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var cmdFilter = &cobra.Command{
	Use:   "filter",
	Short: "Given a json file, it generates a new folder with the desired music",
	RunE:  cmdFilterRunE,
}

func init() {
	cmdFilter.Flags().StringP("src", "s", "Music", "Source music directory")
	cmdFilter.Flags().StringP("file", "f", "", "Json file to filter")
	cmdFilter.Flags().StringP("dst", "d", "", "Destination folder")
	cmdFilter.MarkFlagRequired("src")
	cmdFilter.MarkFlagRequired("file")
	cmdFilter.MarkFlagRequired("dst")

	rootCmd.AddCommand(cmdFilter)
}

func cmdFilterRunE(cmd *cobra.Command, args []string) error {
	var filter music
	fileName, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}

	srcName, err := cmd.Flags().GetString("src")
	if err != nil {
		return err
	}

	dstName, err := cmd.Flags().GetString("dst")
	if err != nil {
		return err
	}

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteJson, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	json.Unmarshal(byteJson, &filter)

	err = os.Mkdir(dstName, 0750)
	if err != nil {
		return err
	}

	// Filtering the content from the json
	for i := 0; i < len(filter.Band); i++ {
		name := filter.Band[i].Name
		songs := filter.Band[i].Songs

		switch {
		case songs[0] == "~":
			continue
		default:
			err := Copy(srcName, dstName, name, songs)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(src string, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	err = destination.Sync()
	if err != nil {
		return err
	}
	return nil
}

func Copy(srcDir string, dstDir string, name string, matchPattern []string) error {
	for _, j := range matchPattern {
		os.Chdir(srcDir)
		matchStrArr, err := filepath.Glob(fmt.Sprintf("%v - %v*", name, j))
		if err != nil {
			return err
		}
		os.Chdir("..")

		for _, m := range matchStrArr {
			srcFile := fmt.Sprintf("%s/%s", srcDir, m)
			dstFile := fmt.Sprintf("%s/%s", dstDir, m)
			err := copyFile(srcFile, dstFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
