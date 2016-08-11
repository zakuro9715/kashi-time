package main

import (
	"fmt"
	"log"
	"os"
)

func usage() {
	fmt.Println("Usage kashi-time [urls...]")
}

func main() {
	urls := os.Args[1:]
	if len(urls) == 0 {
		usage()
		os.Exit(1)
	}
	for _, url := range urls {
		song, err := fetchSong(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n--\n%v\n", song.Title, song.Lyrics)
	}
}
