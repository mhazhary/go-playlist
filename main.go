package main

import (
	"fmt"
	"os"

	"go-playlist/playlist"
)

func main() {
	var menu, title, url string
	fmt.Println("Welcome to Go-Playlist!")
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
			fmt.Scanln(&title)
			fmt.Print("Link : ")
			fmt.Scanln(&url)
			playlist.Add(title, url)
			fmt.Println("=END OF MENU==========================")
		case "2":
			fmt.Println("=START OF MENU========================")
			playlist.Play()
			fmt.Println("=END OF MENU==========================")
		case "00":
			os.Exit(0)
		default:
			fmt.Println("Please input the correct menu!")
		}
	}
}
