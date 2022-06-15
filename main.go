package main

import (
	"fmt"
	"os"

	"go-playlist/playlist"
)

func main() {
	var menu string
	var queue_input playlist.Input
	fmt.Println("Welcome to Go-Playlist!")
	queue_input = playlist.Input{Title: "mme"}
	fmt.Println(queue_input)
	for {
		fmt.Println("Choose the menu:")
		fmt.Println("1.  Input your playlist")
		fmt.Println("2.  Print your playlist")
		fmt.Println("00. EXIT")
		fmt.Print("Menu: ")
		fmt.Scanln(&menu)
		switch menu {
		case "1":
			fmt.Println("=START OF MENU========================")
			fmt.Print("Title: ")
			fmt.Scanln(&queue_input.Title)
			fmt.Print("Link : ")
			fmt.Scanln(&queue_input.Url)
			queue_input.Add()
			fmt.Println("=END OF MENU==========================")
		case "2":
			fmt.Println("=START OF MENU========================")
			queue_input.Play()
			fmt.Println("=END OF MENU==========================")
		case "00":
			os.Exit(0)
		default:
			fmt.Println("Please input the correct menu!")
		}
	}
}
