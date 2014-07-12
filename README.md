#Lastfm-go
Golang wrapper for the Last.fm API 2.0 

[http://www.lastfm.jp/api](http://www.lastfm.jp/api)

#Setup

Get the source codes from github :
  
    % go get github.com/shkh/lastfm-go/lastfm

Import the package :

    import "github.com/shkh/lastfm-go/lastfm"

#Usage

First, create an API instance with your `API KEY` and `API SECRET`.
	
	api := lastfm.New (ApiKey, ApiSecret)

Note that some API methods require your user's permission, so make sure that your requests are authenticated before calling these methods. See "Authentication" section.

API instances contain the structs which represent API classes, and each struct has methods corresponding to their API methods. 
So you can call `user.getNeighbours` for example as following:

	result, _ := api.User.GetNeighbours (lastfm.P{"user": "shkh_"}) //discarding error
	for _, u := range result.Users {
		fmt.Println (u.Name, u.Match)
	}

Methods that fetch some data return their result as a struct named `ClassMethod` (e.g. api.User.GetInfo returns its result of type UserGetInfo).
They can be found in `class_result.go`.
Please look at the file to see which fileds are exported.
		
You can use `lastfm.P` for arguments. 
It's just an alias to `map[string]interface{}`, but values must be `string`, `int`, `int64` (for unix timestamp) or `[]string`.
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
* [album.addTags](http://www.lastfm.jp/api/show/album.addTags)
* [album.getBuylinks](http://www.lastfm.jp/api/show/album.getBuylinks)
* [album.getInfo](http://www.lastfm.jp/api/show/album.getInfo)
* [album.getShouts](http://www.lastfm.jp/api/show/album.getShouts)
* [album.getTags](http://www.lastfm.jp/api/show/album.getTags)
* [album.getTopTags](http://www.lastfm.jp/api/show/album.getTopTags)
* [album.removeTag](http://www.lastfm.jp/api/show/album.removeTag)
* [album.search](http://www.lastfm.jp/api/show/album.search)
* [album.share](http://www.lastfm.jp/api/show/album.share)

##Artist
* [artist.addTags](http://www.lastfm.jp/api/show/artist.addTags)
* [artist.getCorrection](http://www.lastfm.jp/api/show/artist.getCorrection)
* [artist.getEvents](http://www.lastfm.jp/api/show/artist.getEvents)
* [artist.getInfo](http://www.lastfm.jp/api/show/artist.getInfo)
* [artist.getPastEvents](http://www.lastfm.jp/api/show/artist.getPastEvents)
* [artist.getPodcast](http://www.lastfm.jp/api/show/artist.gePodcast)
* [artist.getShouts](http://www.lastfm.jp/api/show/artist.getShouts)
* [artist.getSimilar](http://www.lastfm.jp/api/show/artist.getSimilar)
* [artist.getTags](http://www.lastfm.jp/api/show/artist.getTags)
* [artist.getTopAlbums](http://www.lastfm.jp/api/show/artist.getTopAlbums)
* [artist.getTopFans](http://www.lastfm.jp/api/show/artist.getTopFans)
* [artist.getTopTags](http://www.lastfm.jp/api/show/artist.getTopTags)
* [artist.getTopTracks](http://www.lastfm.jp/api/show/artist.getTopTracks)
* [artist.removeTag](http://www.lastfm.jp/api/show/artist.removeTag)
* [artist.search](http://www.lastfm.jp/api/show/artist.search)
* [artist.share](http://www.lastfm.jp/api/show/artist.share)
* [artist.shout](http://www.lastfm.jp/api/show/artist.shout)

##Auth
* auth.getMobileSession
* auth.getSession
* auth.getToken

##Chart
* [chart.getHypedArtists](http://www.lastfm.jp/api/show/chart.getHypedArtists)
* [chart.getHypedTracks](http://www.lastfm.jp/api/show/chart.getHypedTracks)
* [chart.getLovedTracks](http://www.lastfm.jp/api/show/chart.getLovedTracks)
* [chart.getTopArtists](http://www.lastfm.jp/api/show/chart.getTopArtists)
* [chart.getTopTags](http://www.lastfm.jp/api/show/chart.getTopTags)
* [chart.getTopTracks](http://www.lastfm.jp/api/show/chart.getTopTracks)

##Event
* [event.attend](http://www.lastfm.jp/api/show/event.attend)
* [event.getAttendees](http://www.lastfm.jp/api/show/event.getAttendees)
* [event.getInfo](http://www.lastfm.jp/api/show/event.getInfo)
* [event.getShouts](http://www.lastfm.jp/api/show/event.getShouts)
* [event.share](http://www.lastfm.jp/api/show/event.share)
* [event.shout](http://www.lastfm.jp/api/show/event.shout)

##Geo
* [geo.getEvents](http://www.lastfm.jp/api/show/geo.getEvents)
* [geo.getMetroArtistChart](http://www.lastfm.jp/api/show/geo.getMetroArtistChart)
* [geo.getMetroHypeArtistChart](http://www.lastfm.jp/api/show/geo.getMetroHypeArtistChart)
* [geo.getMetroHypeTrackChart](http://www.lastfm.jp/api/show/geo.getMetroHypeTrackChart)
* [geo.getMetroTrackChart](http://www.lastfm.jp/api/show/geo.getMetroTrackChart)
* [geo.getMetroUniqueArtistChart](http://www.lastfm.jp/api/show/geo.getMetroUniqueArtistChart)
* [geo.getMetroUniqueTrackChart](http://www.lastfm.jp/api/show/geo.getMetroUniqueTrackChart)
* [geo.getMetroWeeklyChartlist](http://www.lastfm.jp/api/show/geo.getMetroWeeklyChartlist)
* [geo.getMetros](http://www.lastfm.jp/api/show/geo.getMetros)
* [geo.getTopArtists](http://www.lastfm.jp/api/show/geo.getTopArtists)
* [geo.getTopTracks](http://www.lastfm.jp/api/show/geo.getTopTracks)

##Group
* [group.getHype](http://www.lastfm.jp/api/show/group.getHype)
* [group.getMembers](http://www.lastfm.jp/api/show/group.getMembers)
* [group.getWeeklyAlbumChart](http://www.lastfm.jp/api/show/group.getWeeklyAlbumChart)
* [group.getWeeklyArtistChart](http://www.lastfm.jp/api/show/group.getWeeklyArtistChart)
* [group.getWeeklyChartList](http://www.lastfm.jp/api/show/group.getWeeklyChartList)
* [group.getWeeklyTrackChart](http://www.lastfm.jp/api/show/group.getWeeklyTrackChart)

##Library
* [library.addAlbum](http://www.lastfm.jp/api/show/library.addAlbum)
* [library.addArtist](http://www.lastfm.jp/api/show/library.addArtist)
* [library.addTrack](http://www.lastfm.jp/api/show/library.addTrack)
* [library.getAlbums](http://www.lastfm.jp/api/show/library.getAlbums)
* [library.getArtists](http://www.lastfm.jp/api/show/library.getArtists)
* [library.getTracks](http://www.lastfm.jp/api/show/library.getTracks)
* [library.removeAlbum](http://www.lastfm.jp/api/show/library.removeAlbum)
* [library.removeArtist](http://www.lastfm.jp/api/show/library.removeArtist)
* [library.removeScrobble](http://www.lastfm.jp/api/show/library.removeScrobble)
* [library.removeTrack](http://www.lastfm.jp/api/show/library.removeTrack)

##Playlist
* [playlist.addTrack](http://www.lastfm.jp/api/show/playlist.addTrack)
* [playlist.create](http://www.lastfm.jp/api/show/playlist.create)

##Radio
* [radio.getPlaylist](http://www.lastfm.jp/api/show/radio.getPlaylist)
* [radio.search](http://www.lastfm.jp/api/show/radio.search)
* [radio.tune](http://www.lastfm.jp/api/show/radio.tune)

## Tag
* [tag.getInfo](http://www.lastfm.jp/api/show/tag.getInfo)
* [tag.getSimilar](http://www.lastfm.jp/api/show/tag.getSimilar)
* [tag.getTopAlbums](http://www.lastfm.jp/api/show/tag.getTopAlbums)
* [tag.getTopArtists](http://www.lastfm.jp/api/show/tag.getTopArtists)
* [tag.getTopTags](http://www.lastfm.jp/api/show/tag.getTopTags)
* [tag.getTopTracks](http://www.lastfm.jp/api/show/tag.getTopTracks)
* [tag.getWeeklyArtistChart](http://www.lastfm.jp/api/show/tag.getWeeklyArtistChart)
* [tag.getWeeklyChartList](http://www.lastfm.jp/api/show/tag.getWeeklyChartList)
* [tag.search](http://www.lastfm.jp/api/show/tag.search)

## Tasteometer
* [tasteometer.compare](http://www.lastfm.jp/api/show/tasteometer.compare)
* ~~tasteometer.compareGroup~~ (deprecated)

## Track
* [track.addTags](http://www.lastfm.jp/api/show/track.addTags)
* [track.ban](http://www.lastfm.jp/api/show/track.ban)
* [track.getBuylinks](http://www.lastfm.jp/api/show/track.getBuylinks)
* [track.getCorrection](http://www.lastfm.jp/api/show/track.getCorrection)
* [track.getFingerprintMetadata](http://www.lastfm.jp/api/show/track.getFingerprintMetadata)
* [track.getInfo](http://www.lastfm.jp/api/show/track.getInfo)
* [track.getShouts](http://www.lastfm.jp/api/show/track.getShouts)
* [track.getSimilar](http://www.lastfm.jp/api/show/track.getSimilar)
* [track.getTags](http://www.lastfm.jp/api/show/track.getTags)
* [track.getTopFans](http://www.lastfm.jp/api/show/track.getTopFans)
* [track.getTopTags](http://www.lastfm.jp/api/show/track.getTopTags)
* [track.love](http://www.lastfm.jp/api/show/track.love)
* [track.removeTag](http://www.lastfm.jp/api/show/track.removeTag)
* [track.scrobble](http://www.lastfm.jp/api/show/track.scrobble)
* [track.search](http://www.lastfm.jp/api/show/track.search)
* [track.share](http://www.lastfm.jp/api/show/track.share)
* [track.unban](http://www.lastfm.jp/api/show/track.unban)
* [track.unlove](http://www.lastfm.jp/api/show/track.unlove)
* [track.updateNowPlaying](http://www.lastfm.jp/api/show/track.updateNowPlaying)

## User
* [user.getArtistTracks](http://www.lastfm.jp/api/show/user.getArtistTracks)
* [user.getBannedTracks](http://www.lastfm.jp/api/show/user.getBannedTracks)
* [user.getEvents](http://www.lastfm.jp/api/show/user.getEvents)
* [user.getFriends](http://www.lastfm.jp/api/show/user.getFriends)
* [user.getInfo](http://www.lastfm.jp/api/show/user.getInfo)
* [user.getLovedTracks](http://www.lastfm.jp/api/show/user.getLovedTracks)
* [user.getNeighbours](http://www.lastfm.jp/api/show/user.getNeighbours)
* [user.getNewReleases](http://www.lastfm.jp/api/show/user.getNewReleases)
* [user.getPastEvents](http://www.lastfm.jp/api/show/user.getPastEvents)
* [user.getPersonalTags](http://www.lastfm.jp/api/show/user.getPersonalTags)
* [user.getPlaylists](http://www.lastfm.jp/api/show/user.getPlaylists)
* [user.getRecentStations](http://www.lastfm.jp/api/show/user.getRecentStations)
* [user.getRecentTracks](http://www.lastfm.jp/api/show/user.getRecentTracks)
* [user.getRecommendedArtists](http://www.lastfm.jp/api/show/user.getRecommendedArtists)
* [user.getRecommendedEvents](http://www.lastfm.jp/api/show/user.getRecommendedEvents)
* [user.getShouts](http://www.lastfm.jp/api/show/user.getShouts)
* [user.getTopAlbums](http://www.lastfm.jp/api/show/user.getTopAlbums)
* [user.getTopArtists](http://www.lastfm.jp/api/show/user.getTopArtists)
* [user.getTopTags](http://www.lastfm.jp/api/show/user.getTopTags)
* [user.getTopTracks](http://www.lastfm.jp/api/show/user.getTopTracks)
* [user.getWeeklyAlbumChart](http://www.lastfm.jp/api/show/user.getWeeklyAlbumChart)
* [user.getWeeklyArtistChart](http://www.lastfm.jp/api/show/user.getWeeklyArtistChart)
* [user.getWeeklyChartList](http://www.lastfm.jp/api/show/user.getWeeklyChartList)
* [user.getWeeklyTrackChart](http://www.lastfm.jp/api/show/user.getWeeklyTrackChart)
* [user.shout](http://www.lastfm.jp/api/show/user.shout)
* ~~user.signUp~~ (deprecated)
* ~~user.terms~~ (deprecated)

##Venue
* [venue.getEvents](http://www.lastfm.jp/api/show/venue.getEvents)
* [venue.getPastEvents](http://www.lastfm.jp/api/show/venue.getPastEvents)
* [venue.search](http://www.lastfm.jp/api/show/venue.search)

#Licence
MIT Licenced. See [LICENCE](https://github.com/shkh/lastfm-go/blob/master/LICENSE).
