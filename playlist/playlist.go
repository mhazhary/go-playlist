package playlist

import (
	"fmt"
)

type playlists []struct {
	title string
	url   string
}

var p playlists

func Add(new_title string, new_url string) {
	if new_title == "" || new_url == "" {
		fmt.Println("Please input the title and URL")
		return
	}
	p = append(p, struct {
		title string
		url   string
	}{
		title: new_title,
		url:   new_url,
	})
	fmt.Println("New queue has been added!")
}

func Play() {
	fmt.Println("Here is your playlist")
	if p == nil {
		fmt.Println("There is no playlist added!")
	}
	for i, song := range p {
		fmt.Printf("%d. Title: %s\n   Link : %s\n", i+1, song.title, song.url)
	}
	p = nil
}
