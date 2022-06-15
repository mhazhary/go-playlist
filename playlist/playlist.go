package playlist

import (
	"fmt"
)

type queue struct {
	title string
	url   string
}

var q []queue

func Add(title string, url string) {
	if title == "" || url == "" {
		fmt.Println("Please input the title and URL")
		return
	}
	q = append(q, queue{
		title: title,
		url:   url,
	})
	fmt.Println("New queue has been added!")
}

func Play() {
	fmt.Println("Here is your playlist")
	if q == nil {
		fmt.Println("There is no playlist added!")
	}
	for i, s := range q {
		fmt.Printf("%d. Title: %s\n   Link : %s\n", i+1, s.title, s.url)
	}
	q = nil
}
