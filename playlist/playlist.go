package playlist

import (
	"fmt"
)

type queue []struct {
	Title string
	Url   string
}

type Input struct {
	Title string
	Url   string
}

var q queue

func (i Input) Add() {
	if i.Title == "" || i.Url == "" {
		fmt.Println("Please input the title and URL")
		return
	}
	q = append(q, i)
	fmt.Println("New q has been added!")
}

func (i Input) Play() {
	fmt.Println("Here is your playlist")
	if q == nil {
		fmt.Println("There is no playlist added!")
	}
	for i, s := range q {
		fmt.Printf("%d. Title: %s\n   Link : %s\n", i+1, s.Title, s.Url)
	}
	q = nil
}
