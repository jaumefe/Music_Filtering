package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// var opts struct {
// 	Config   string `short:"c" long:"config" description:"Provide json file to filter"`
// 	Template bool   `short:"t" long:"template" description:"Generate a template json with all the artists included (By default, it's false)"`
// }

type band struct {
	Name  string   `json:"artist"`
	Songs []string `json:"songs"`
}

type music struct {
	Band     []band `json:"Bands"`
	Playlist string `json:"playlist"`
}

func main() {

	music := music{Playlist: "Music", Band: make([]band, 0)}
	dir, err := os.ReadDir("Music")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range dir {
		singSong := strings.Split(file.Name()[:len(file.Name())-4], " - ")
		fmt.Println(singSong)
		if len(singSong) <= 1 {
			fmt.Printf("Unable to parse %v\n", singSong)
		} else {
			singer, songs := singSong[0], singSong[1]
			if (len(music.Band) == 0) || (music.Band[len(music.Band)-1].Name != singer) {
				music.Band = append(music.Band, band{Name: singer})
			}
			music.Band[len(music.Band)-1].Songs = append(music.Band[len(music.Band)-1].Songs, songs)
		}
	}

	a := app.New()
	w := a.NewWindow("Music List")
	w.Resize(fyne.Size{Width: 1080, Height: 480})

	// Función para obtener los nombres de los artistas
	getArtists := func() []string {
		var artists []string
		for _, b := range music.Band {
			artists = append(artists, b.Name)
		}
		return artists
	}

	// Función para obtener las canciones de un artista
	getSongs := func(artist string) []string {
		for _, b := range music.Band {
			if b.Name == artist {
				return b.Songs
			}
		}
		return []string{}
	}

	// Obtiene la lista de artistas
	artists := getArtists()
	artistSelect := widget.NewSelect(artists, nil)

	// Crea una lista para mostrar las canciones
	songList := widget.NewList(
		func() int {
			return 0 // Inicialmente vacía
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText("") // Inicialmente vacía
		},
	)

	// Función para actualizar la lista de canciones
	updateSongList := func(artist string) {
		songs := getSongs(artist)
		songList.Length = func() int {
			return len(songs)
		}
		songList.UpdateItem = func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(songs[i])
		}
		songList.Refresh()
	}

	// Maneja la selección de un artista
	artistSelect.OnChanged = func(artist string) {
		updateSongList(artist)
	}

	content := container.NewBorder(artistSelect, nil, nil, nil, songList)
	w.SetContent(content)

	w.ShowAndRun()
}

// func copyFile(src string, dst string) error {
// 	source, err := os.Open(src)
// 	if err != nil {
// 		return err
// 	}
// 	defer source.Close()

// 	destination, err := os.Create(dst)
// 	if err != nil {
// 		return err
// 	}
// 	defer destination.Close()

// 	_, err = io.Copy(destination, source)
// 	if err != nil {
// 		return err
// 	}

// 	err = destination.Sync()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
