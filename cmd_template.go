package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var cmdTemplate = &cobra.Command{
	Use:   "template",
	Short: "Provide a json file with all the music gathered by artist",
	RunE:  cmdTemplateRunE,
}

func init() {
	cmdTemplate.Flags().StringP("src", "s", "Music", "Source music directory")

	rootCmd.AddCommand(cmdTemplate)
}

func cmdTemplateRunE(cmd *cobra.Command, args []string) error {
	music := music{Band: make([]band, 0)}

	dirName, err := cmd.Flags().GetString("src")
	if err != nil {
		return err
	}

	err = music.getSongsListFolder(dirName)
	if err != nil {
		return err
	}

	content, err := json.MarshalIndent(music, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("%s.json", dirName), content, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (m *music) getSongsListFolder(dirName string) error {
	dir, err := os.ReadDir(dirName)
	if err != nil {
		return err
	}

	for _, file := range dir {
		singSong := strings.Split(file.Name(), " - ")
		if len(singSong) <= 1 {
			fmt.Printf("Unable to parse %v\n", singSong)
		} else {
			singer, songs := singSong[0], singSong[1]
			if (len(m.Band) == 0) || (m.Band[len(m.Band)-1].Name != singer) {
				m.Band = append(m.Band, band{Name: singer})
			}
			m.Band[len(m.Band)-1].Songs = append(m.Band[len(m.Band)-1].Songs, songs)
		}
	}
	return nil
}
