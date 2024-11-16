package main

import (
	"encoding/json"
	"os"

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
	cmdDiff.Flags().StringP("out", "o", "diff.json", "Json file with the differences between both folders")

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

	diffMusic := music{Band: make([]band, 0)}
	err = diffMusic.checkDiff(srcMusic, dstMusic)
	if err != nil {
		return err
	}

	content, err := json.MarshalIndent(diffMusic, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile("diff.json", content, 0644)
	if err != nil {
		return err
	}

	return nil
}

func musicToMss(m music) map[string][]string {
	musicMsi := make(map[string][]string, 0)
	for _, band := range m.Band {
		musicMsi[band.Name] = band.Songs
	}

	return musicMsi
}

func (m *music) checkDiff(source music, destination music) error {
	srcMsi := musicToMss(source)
	dstMsi := musicToMss(destination)

	for b := range srcMsi {
		songs := make([]string, 0)
		if _, ok := dstMsi[b]; ok {
			for _, s := range srcMsi[b] {
				if !isStringInSlice(s, dstMsi[b]) {
					songs = append(songs, s)
				}
			}
		} else {
			songs = append(songs, srcMsi[b]...)
		}

		if len(songs) > 0 {
			m.Band = append(m.Band, band{Name: b, Songs: songs})
		}
	}

	return nil
}

func isStringInSlice(pattern string, arr []string) bool {
	for _, i := range arr {
		if pattern == i {
			return true
		}
	}
	return false
}
