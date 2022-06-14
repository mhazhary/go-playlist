package playlist

import (
	"fmt"
)

type Queue struct {
	title string
	url   string
}

var queue []Queue

func Add(title string, url string) {
	if title == "" || url == "" {
		fmt.Println("Please input the title and URL")
		return
	}
	queue = append(queue, Queue{
		title: title,
		url:   url,
	})
	fmt.Println("New queue has been added!")
}

func Play() {
	fmt.Println("Here is your playlist")
	if queue == nil {
		fmt.Println("There is no playlist added!")
	}
	for i, s := range queue {
		fmt.Printf("%d. Title: %s\n   Link : %s\n", i+1, s.title, s.url)
	}
	queue = nil
}
