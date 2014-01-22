package main

import (
	"../lastfm"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getTrimmedString(r *bufio.Reader, msg string) (res string) {
	fmt.Print(msg, ":")
	res, _ = r.ReadString('\n')
	res = strings.Trim(res, "\r\n")
	return
}

func main() {
	r := bufio.NewReader(os.Stdin)
	apiKey := getTrimmedString(r, "API KEY:")
	apiSecret := getTrimmedString(r, "API SECRET")

	api := lastfm.New(apiKey, apiSecret)

	username := getTrimmedString(r, "Username")
	password := getTrimmedString(r, "Password")

	err := api.Login(username, password)
	if err != nil {
		fmt.Println(err)
		return
	}

	target := getTrimmedString(r, "Target")
	message := getTrimmedString(r, "Message")

	err = api.User.Shout(lastfm.P{"user": target, "message": message})
	if err != nil {
		fmt.Println(err)
	}
}
