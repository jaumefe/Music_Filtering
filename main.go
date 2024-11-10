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

	// Crea el accordion para mostrar los artistas y sus canciones
	// accordion := widget.NewAccordion()

	// Añadir artistas y sus canciones al accordion
	// for _, artist := range music.Band {
	// 	accordion.Append(widget.NewAccordionItem(artist.Name, widget.NewLabel("Test")))
	// }

	a := app.New()
	w := a.NewWindow("Music List")
	w.Resize(fyne.Size{Width: 1080, Height: 480})

	displayLimit := 100
	if len(music.Band) < displayLimit {
		displayLimit = len(music.Band)
	}

	accordion := widget.NewAccordion()

	// for i := 0; i < 409; i++ {
	// 	item := widget.NewAccordionItem(strconv.Itoa(i), widget.NewLabel("Prova"))
	// 	accordion.Append(item)
	// }

	// w.SetContent(container.NewVBox(
	// 	widget.NewLabel("Accordion Example"),
	// 	accordion,
	// ))

	// Set the initial number of items to display and increment size
	initialDisplayCount := 100
	increment := 100

	// Function to load items into the accordion
	loadItems := func(start, end int) {
		for i := start; i < end && i < len(music.Band); i++ {
			artist := music.Band[i]
			item := widget.NewAccordionItem(artist.Name, widget.NewLabel("Test"))
			accordion.Append(item)
		}
	}

	// Initially load items
	loadItems(0, initialDisplayCount)

	// Create a button to load more items
	loadMoreButton := widget.NewButton("Load More", func() {
		currentCount := len(accordion.Items)
		nextCount := currentCount + increment
		loadItems(currentCount, nextCount)

	})

	// Wrap the accordion in a scroll container
	scrollContainer := container.NewScroll(accordion)
	scrollContainer.SetMinSize(fyne.NewSize(1080, 400)) // Set the minimum size for the scroll container

	// Set the content of the window to include the scroll container and the load more button
	w.SetContent(container.NewVBox(
		widget.NewLabel("Accordion Example"),
		scrollContainer,
		loadMoreButton,
	))

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
