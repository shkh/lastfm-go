package main

/*
DO NOT USE THIS CODE IN YOUR PRODUCTION.
*/

import (
	"../lastfm"
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const TopPage = `
<html>
<body>
<a href="%v">Login to Last.fm</a>
</body>
</html>
`

const ShowPage = `
<html>
<body>
%v
</body>
</html>
`

var api *lastfm.Api

func init() {
	fmt.Print("API KEY:")
	inputReader := bufio.NewReader(os.Stdin)
	apiKey, _ := inputReader.ReadString('\n')
	apiKey = strings.Trim(apiKey, "\r\n")
	fmt.Print("API SECRET:")
	apiSecret, _ := inputReader.ReadString('\n')
	apiSecret = strings.Trim(apiSecret, "\r\n")

	api = lastfm.New(apiKey, apiSecret)
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	callback := "http://localhost:8080/authorized"
	io.WriteString(writer, fmt.Sprintf(TopPage, api.GetAuthRequestUrl(callback)))
}

func authorizedHandler(writer http.ResponseWriter, request *http.Request) {
	query, _ := url.ParseQuery(request.URL.RawQuery)
	token := query.Get("token")
	api.LoginWithToken(token)
	http.Redirect(writer, request, "http://localhost:8080/show", 302)
}

func showInfoHandler(writer http.ResponseWriter, request *http.Request) {
	result, err := api.User.GetInfo(nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	var t string
	t += "<p>" + result.Id + "</p>"
	t += "<p>" + result.Name + "</p>"
	t += "<p>" + result.RealName + "</p>"
	t += "<p>" + result.Url + "</p>"

	io.WriteString(writer, fmt.Sprintf(ShowPage, t))
}

func main() {
	http.HandleFunc("/authorized", authorizedHandler)
	http.HandleFunc("/show", showInfoHandler)
	http.HandleFunc("/", rootHandler)
	log.Println("listening at 127.0.0.1:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
