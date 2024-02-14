# Lastfm-go
Golang wrapper for the Last.fm API 2.0

[https://www.last.fm/api/](https://www.last.fm/api/)

# Setup

Get the source codes from github :

    % go get github.com/shkh/lastfm-go/lastfm

Import the package :

    import "github.com/shkh/lastfm-go/lastfm"

# Usage

First, create an API instance with your `API KEY` and `API SECRET`.

	api := lastfm.New (ApiKey, ApiSecret)

Note that some API methods require your user's permission, so make sure that your requests are authenticated before calling these methods. See "Authentication" section.

API instances contain the structs which represent API classes, and each struct has methods corresponding to their API methods.
So you can call `artist.getTopTracks` for example as following:

	result, _ := api.Artist.GetTopTracks(lastfm.P{"artist": "Avicii"}) //discarding error
	for _, track := range result.Tracks {
		fmt.Println(track.Name)
	}

Methods that fetch some data return their result as a struct named `ClassMethod` (e.g. api.User.GetInfo returns its result of type UserGetInfo).
They can be found in `class_result.go`.
Please look at the file to see which fields are exported.

You can use `lastfm.P` for arguments.
It's just an alias to `map[string]interface{}`, but values must be `string`, `int`, `int64` (for unix timestamp) or `[]string`.
Slice of string, []string, can be used for passing multiple values for a key.

	//album.addTags (auth required)
	api.Album.AddTags(lastfm.P{ //discarding error
		"artist": "Kaene",
		"album":  "Strangeland",
		"tags":   []string{"britpop", "alternative rock", "2012"},
	})



# Authentication
There are three ways to authenticate your requests, which to choose depends on what kind of application you are making.

- for Mobile Apps

		err = api.Login (username, password)

- for Desktop Apps

		token, _ = api.GetToken () //discarding error
		authUrl = api.GetAuthTokenUrl (token)
		//Send your user to "authUrl"
		//Once the user grant permission, then authorize the token.
		api.LoginWithToken (token) //discarding error

- for Web Apps

		callback = "https://spam.hum"
		authUrl, _ = api.GetAuthRequestUrl (callback)
		//Send your user to "authUrl"
		//Get the token embeded in the redirected URL, then authorize the token.
		api.LoginWithToken (token) //discarding error


# Supported APIs

## Album
* [album.addTags](https://www.last.fm/api/show/album.addTags)
* [album.getInfo](https://www.last.fm/api/show/album.getInfo)
* [album.getTags](https://www.last.fm/api/show/album.getTags)
* [album.getTopTags](https://www.last.fm/api/show/album.getTopTags)
* [album.removeTag](https://www.last.fm/api/show/album.removeTag)
* [album.search](https://www.last.fm/api/show/album.search)

## Artist
* [artist.addTags](https://www.last.fm/api/show/artist.addTags)
* [artist.getCorrection](https://www.last.fm/api/show/artist.getCorrection)
* [artist.getInfo](https://www.last.fm/api/show/artist.getInfo)
* [artist.getSimilar](https://www.last.fm/api/show/artist.getSimilar)
* [artist.getTags](https://www.last.fm/api/show/artist.getTags)
* [artist.getTopAlbums](https://www.last.fm/api/show/artist.getTopAlbums)
* [artist.getTopTags](https://www.last.fm/api/show/artist.getTopTags)
* [artist.getTopTracks](https://www.last.fm/api/show/artist.getTopTracks)
* [artist.removeTag](https://www.last.fm/api/show/artist.removeTag)
* [artist.search](https://www.last.fm/api/show/artist.search)

## Auth
* [auth.getMobileSession](https://www.last.fm/api/show/auth.getMobileSession)
* [auth.getSession](https://www.last.fm/api/show/auth.getSession)
* [auth.getToken](https://www.last.fm/api/show/auth.getToken)

## Chart
* [chart.getTopArtists](https://www.last.fm/api/show/chart.getTopArtists)
* [chart.getTopTags](https://www.last.fm/api/show/chart.getTopTags)
* [chart.getTopTracks](https://www.last.fm/api/show/chart.getTopTracks)

## Geo
* [geo.getTopArtists](https://www.last.fm/api/show/geo.getTopArtists)
* [geo.getTopTracks](https://www.last.fm/api/show/geo.getTopTracks)

## Library
* [library.getArtists](https://www.last.fm/api/show/library.getArtists)

## Tag
* [tag.getInfo](https://www.last.fm/api/show/tag.getInfo)
* [tag.getSimilar](https://www.last.fm/api/show/tag.getSimilar)
* [tag.getTopAlbums](https://www.last.fm/api/show/tag.getTopAlbums)
* [tag.getTopArtists](https://www.last.fm/api/show/tag.getTopArtists)
* [tag.getTopTags](https://www.last.fm/api/show/tag.getTopTags)
* [tag.getTopTracks](https://www.last.fm/api/show/tag.getTopTracks)
* [tag.getWeeklyChartList](https://www.last.fm/api/show/tag.getWeeklyChartList)

## Track
* [track.addTags](https://www.last.fm/api/show/track.addTags)
* [track.getCorrection](https://www.last.fm/api/show/track.getCorrection)
* [track.getInfo](https://www.last.fm/api/show/track.getInfo)
* [track.getSimilar](https://www.last.fm/api/show/track.getSimilar)
* [track.getTags](https://www.last.fm/api/show/track.getTags)
* [track.getTopTags](https://www.last.fm/api/show/track.getTopTags)
* [track.love](https://www.last.fm/api/show/track.love)
* [track.removeTag](https://www.last.fm/api/show/track.removeTag)
* [track.scrobble](https://www.last.fm/api/show/track.scrobble)
* [track.search](https://www.last.fm/api/show/track.search)
* [track.unlove](https://www.last.fm/api/show/track.unlove)
* [track.updateNowPlaying](https://www.last.fm/api/show/track.updateNowPlaying)

## User
* [user.getArtistTracks](https://www.last.fm/api/show/user.getArtistTracks)
* [user.getFriends](https://www.last.fm/api/show/user.getFriends)
* [user.getInfo](https://www.last.fm/api/show/user.getInfo)
* [user.getLovedTracks](https://www.last.fm/api/show/user.getLovedTracks)
* [user.getPersonalTags](https://www.last.fm/api/show/user.getPersonalTags)
* [user.getRecentTracks](https://www.last.fm/api/show/user.getRecentTracks)
* [user.getTopAlbums](https://www.last.fm/api/show/user.getTopAlbums)
* [user.getTopArtists](https://www.last.fm/api/show/user.getTopArtists)
* [user.getTopTags](https://www.last.fm/api/show/user.getTopTags)
* [user.getTopTracks](https://www.last.fm/api/show/user.getTopTracks)
* [user.getWeeklyAlbumChart](https://www.last.fm/api/show/user.getWeeklyAlbumChart)
* [user.getWeeklyArtistChart](https://www.last.fm/api/show/user.getWeeklyArtistChart)
* [user.getWeeklyChartList](https://www.last.fm/api/show/user.getWeeklyChartList)
* [user.getWeeklyTrackChart(https://www.last.fm/api/show/user.getWeeklyTrackChart)

# Licence
MIT Licenced. See [LICENCE](https://github.com/shkh/lastfm-go/blob/master/LICENSE).
