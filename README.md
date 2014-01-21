#Lastfm-go
Golang wrapper for the Last.fm API 2.0 

[http://www.lastfm.jp/api](http://www.lastfm.jp/api)

#Usage

First, create an API instance with your `API KEY` and `API SECRET`.
	
	api := lastfm.New (ApiKey, ApiSecret)

Note that some API methods require your user's permission, so make sure that your requests are authenticated before calling these methods. See "Authentication" section.

API instances contain the structs which represent API classes, and each struct has methods corresponding to their API methods. 
So you can call `user.getNeighbours` for example as following:

	result, _ := api.User.GetNeighbours (latfm.P{"user": "shkh_"}) //discarding error
	for _, u := range result.Users {
		fmt.Println (u.Name, u.Match)
	}

Methods that fetch some data return their result as a struct named `ClassMethod` (e.g. api.User.GetInfo returns its result of type UserGetInfo).
They can be found in `class_result.go`.
Please look at the file to see which fileds are exported.
		
You can use `lastfm.P` for arguments. 
It's just an alias to `map[string]interface{}`, but values must be `string`, `int` or `[]string`.
Slice of string, []string, can be used for passing multiple values for a key.

	//album.addTags (auth required)
	api.Album.AddTags(lastfm.P{ //discarding error
		"artist": "Kaene",
		"album":  "Strangeland",
		"tags":   []string{"britpop", "alternative rock", "2012"},
	})
	
	//library.addAlbum (auth required)
	api.Library.AddAlbum(lastfm.P{ //discarding error
		"artist": []string{"Ellie Goulding", "Keane"},
		"album":  []string{"Goodness Gracious", "Higher Than The Sun"},
	})



#Authentication
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

		callback = "http://spam.hum"
		authUrl, _ = api.GetAuthRequestUrl (callback)
		//Send your user to "authUrl"
		//Get the token embeded in the redirected URL, then authorize the token.
		api.LoginWithToken (token) //discarding error
	

#Supported APIs

##Album
* album.addTags
* album.getBuylinks
* album.getInfo
* album.getShouts
* album.getTags
* album.getTopTags
* album.removeTag
* album.search
* album.share

##Artist
* artist.addTags
* artist.getCorrection
* artist.getEvents
* artist.getInfo
* artist.getPastEvents
* artist.getPodcast
* artist.getShouts
* artist.getSimilar
* artist.getTags
* artist.getTopAlbums
* artist.getTopFans
* artist.getTopTags
* artist.getTopTracks
* artist.removeTag
* artist.search
* artist.share
* artist.shout

##Auth
* auth.getMobileSession
* auth.getSession
* auth.getToken

##Chart
* chart.getHypedArtists
* chart.getHypedTracks
* chart.getLovedTracks
* chart.getTopArtists
* chart.getTopTags
* chart.getTopTracks

##Event
* event.attend
* event.getAttendees
* event.getInfo
* event.getShouts
* event.share
* event.shout

##Geo
* geo.getEvents
* geo.getMetroArtistChart
* geo.getMetroHypeArtistChart
* geo.getMetroHypeTrackChart
* geo.getMetroTrackChart
* geo.getMetroUniqueArtistChart
* geo.getMetroUniqueTrackChart
* geo.getMetroWeeklyChartlist
* geo.getMetros
* geo.getTopArtists
* geo.getTopTracks

##Group
* group.getHype
* group.getMembers
* group.getWeeklyAlbumChart
* group.getWeeklyArtistChart
* group.getWeeklyChartList
* group.getWeeklyTrackChart

##Library
* library.addAlbum
* library.addArtist
* library.addTrack
* library.getAlbums
* library.getArtists
* library.getTracks
* library.removeAlbum
* library.removeArtist
* library.removeScrobble
* library.removeTrack

##Playlist
* playlist.addTrack
* playlist.create

##Radio
* radio.getPlaylist
* radio.search
* radio.tune

## Tag
* tag.getInfo
* tag.getSimilar
* tag.getTopAlbums
* tag.getTopArtists
* tag.getTopTags
* tag.getTopTracks
* tag.getWeeklyArtistChart
* tag.getWeeklyChartList
* tag.search

## Tasteometer
* tasteometer.compare
* ~~tasteometer.compareGroup~~ (deprecated)

## Track
* track.addTags
* track.ban
* track.getBuylinks
* track.getCorrection
* track.getFingerprintMetadata
* track.getInfo
* track.getShouts
* track.getSimilar
* track.getTags
* track.getTopFans
* track.getTopTags
* track.love
* track.removeTag
* track.scrobble
* track.search
* track.share
* track.unban
* track.unlove
* track.updateNowPlaying

## User
* user.getArtistTracks
* user.getBannedTracks
* user.getEvents
* user.getFriends
* user.getInfo
* user.getLovedTracks
* user.getNeighbours
* user.getNewReleases
* user.getPastEvents
* user.getPersonalTags
* user.getPlaylists
* user.getRecentStations
* user.getRecentTracks
* user.getRecommendedArtists
* user.getRecommendedEvents
* user.getShouts
* user.getTopAlbums
* user.getTopArtists
* user.getTopTags
* user.getTopTracks
* user.getWeeklyAlbumChart
* user.getWeeklyArtistChart
* user.getWeeklyChartList
* user.getWeeklyTrackChart
* user.shout
* ~~user.signUp~~ (deprecated)
* ~~user.terms~~ (deprecated)

##Venue
* venue.getEvents
* venue.getPastEvents
* venue.search



