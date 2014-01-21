package main

import (
	"../lastfm"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("API KEY:")
	apiKey, _ := inputReader.ReadString('\n')
	apiKey = strings.Trim(apiKey, "\r\n")
	fmt.Print("API SECRET:")
	apiSecret, _ := inputReader.ReadString('\n')
	apiSecret = strings.Trim(apiSecret, "\r\n")

	api := lastfm.New(apiKey, apiSecret)

	fmt.Print("Username:")
	userName, _ := inputReader.ReadString('\n')
	userName = strings.Trim(userName, "\r\n")
	fmt.Print("Password:")
	password, _ := inputReader.ReadString('\n')
	password = strings.Trim(password, "\r\n")

	err := api.Login(userName, password)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Artist:")
	artist, _ := inputReader.ReadString('\n')
	artist = strings.Trim(artist, "\r\n")
	fmt.Print("Track:")
	track, _ := inputReader.ReadString('\n')
	track = strings.Trim(track, "\r\n")

	p := lastfm.P{"artist": artist, "track": track}

	start := time.Now().Unix()
	_, err = api.Track.UpdateNowPlaying(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Updated Now-Playing.")

	time.Sleep(35 * time.Second)
	p["timestamp"] = start
	_, err = api.Track.Scrobble(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Scrobbled.")
}
