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
	rootCmd.PersistentFlags().StringP("src", "s", "Music", "Source music directory")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	// if template {
	// 	// Reading Music directory
	// 	music := music{Playlist: "Music", Band: make([]band, 0)}
	// 	dir, err := os.ReadDir("Music")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	for _, file := range dir {
	// 		singSong := strings.Split(file.Name(), " - ")
	// 		if len(singSong) <= 1 {
	// 			fmt.Printf("Unable to parse %v\n", singSong)
	// 		} else {
	// 			singer, songs := singSong[0], singSong[1]
	// 			if (len(music.Band) == 0) || (music.Band[len(music.Band)-1].Name != singer) {
	// 				music.Band = append(music.Band, band{Name: singer})
	// 			}
	// 			music.Band[len(music.Band)-1].Songs = append(music.Band[len(music.Band)-1].Songs, songs)
	// 		}
	// 	}

	// 	content, err := json.MarshalIndent(music, "", "\t")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	err = os.WriteFile("All_Music.json", content, 0644)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// } else if config != "" {
	// 	// Opening and parsing json file for filtering songs
	// 	var filter music
	// 	jsonFile, err := os.Open(config)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	defer jsonFile.Close()

	// 	byteValue, _ := io.ReadAll(jsonFile)
	// 	json.Unmarshal(byteValue, &filter)

	// 	err = os.Mkdir(filter.Playlist, 0750)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	// Filtering the content from the json
	// 	for i := 0; i < len(filter.Band); i++ {
	// 		name := filter.Band[i].Name
	// 		songs := filter.Band[i].Songs
	// 		switch {
	// 		case songs[0] == "*":
	// 			os.Chdir("Music/")
	// 			match, err := filepath.Glob(fmt.Sprintf("%v - *", name))
	// 			os.Chdir("../")
	// 			if err != nil {
	// 				fmt.Println(err)
	// 			}
	// 			for _, j := range match {
	// 				src := fmt.Sprintf("Music/%v", j)
	// 				dst := fmt.Sprintf("%v/%v", filter.Playlist, j)
	// 				err := copyFile(src, dst)
	// 				if err != nil {
	// 					fmt.Println(err)
	// 				}
	// 			}
	// 		case songs[0] == "~":
	// 			continue
	// 		default:
	// 			for _, j := range songs {
	// 				os.Chdir("Music/")
	// 				match, err := filepath.Glob(fmt.Sprintf("%v - %v*", name, j))
	// 				os.Chdir("../")
	// 				if err != nil {
	// 					fmt.Println(err)
	// 				}
	// 				src := fmt.Sprintf("Music/%v", match[0])
	// 				dst := fmt.Sprintf("%v/%v", filter.Playlist, match[0])
	// 				err = copyFile(src, dst)
	// 				if err != nil {
	// 					fmt.Println(err)
	// 				}
	// 			}
	// 		}
	// 	}
	// } else {
	// 	fmt.Println("No filter in json format provided")
	// }

}
