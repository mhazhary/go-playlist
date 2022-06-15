package playlist

import (
	"fmt"
)

type Input interface {
	Add(queuer Queuer)
	Play()
}

type Queuer struct {
	Title string
	Url   string
}

var q []Queuer

func Add(queuer Queuer) {
	if queuer.Title == "" || queuer.Url == "" {
		fmt.Println("Please input the title and URL")
		return
	}
	q = append(q, Queuer{
		Title: queuer.Title,
		Url:   queuer.Url,
	})
	fmt.Println("New q has been added!")
}

func Play() {
	fmt.Println("Here is your playlist")
	if q == nil {
		fmt.Println("There is no playlist added!")
	}
	for i, s := range q {
		fmt.Printf("%d. Title: %s\n   Link : %s\n", i+1, s.Title, s.Url)
	}
	q = nil
}
