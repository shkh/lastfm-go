package lastfm

type artistApi struct {
	creds *credentials
}

//artist.addTags
func (api artistApi) AddTags(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Artist.AddTags") }()
	err = callPost("artist.addtags", api.creds, args, nil, P{
		"normal": []string{"artist", "tags"},
	})
	return
}

//artist.getCorrection
func (api artistApi) GetCorrection(args P) (result ArtistGetCorrection, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetCorrection") }()
	err = callGet("artist.getcorrection", api.creds, args, &result, P{
		"normal": []string{"artist"},
	})
	return
}

//artist.getEvents
func (api artistApi) GetEvents(args P) (result ArtistGetEvents, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetEvents") }()
	err = callGet("artist.getevents", api.creds, args, &result, P{
		"normal": []string{"artist", "mbid", "autocorrect", "limit", "page", "festivalsonly"},
	})
	return

}

//artist.getInfo
func (api artistApi) GetInfo(args P) (result ArtistGetInfo, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetInfo") }()
	if _, ok := args["username"]; !ok {
		args["username"] = api.creds.username
	}
	err = callGet("artist.getinfo", api.creds, args, &result, P{
		"normal": []string{"artist", "mbid", "lang", "autocorrect", "username"},
	})
	return
}

//artist.getPastEvents
func (api artistApi) GetPastEvents(args P) (result ArtistGetPastEvents, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetPastEvents") }()
	err = callGet("artist.getpastevents", api.creds, args, &result, P{
		"normal": []string{"artist", "mbid", "page", "autocorrect", "limit"},
	})
	return
}

//artist.getPodcast
func (api artistApi) GetPodcast(args P) (result ArtistGetPodcast, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetPodcast") }()
	err = callGet("artist.getpodcast", api.creds, args, &result, P{
		"normal": []string{"artist", "mbid", "autocorrect"},
	})
	return
}

//artist.getShouts
func (api artistApi) GetShouts(args P) (result ArtistGetShouts, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetShouts") }()
	err = callGet("artist.getshouts", api.creds, args, &result, P{
		"normal": []string{"artist", "mbid", "limit", "autocorrect", "page"},
	})
	return
}

//artist.getSimilar
func (api artistApi) GetSimilar(args P) (result ArtistGetSimilar, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetSimilar") }()
	err = callGet("artist.getsimilar", api.creds, args, &result, P{
		"normal": []string{"limit", "artist", "autocorrect", "mbid"},
	})
	return
}

//artist.getTags
func (api artistApi) GetTags(args P) (result ArtistGetTags, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetTags") }()
	if _, ok := args["user"]; !ok {
		args["user"] = api.creds.username
	}
	err = callGet("artist.gettags", api.creds, args, &result, P{
		"normal": []string{"artist", "mbid", "user", "autocorrect"},
	})
	return
}

//artist.getTopAlbums
func (api artistApi) GetTopAlbums(args P) (result ArtistGetTopAlbums, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetTopAlbums") }()
	err = callGet("artist.gettopalbums", api.creds, args, &result, P{
		"normal": []string{"artist", "mbid", "autocorrect", "page", "limit"},
	})
	return
}

//artist.getTopFans
func (api artistApi) GetTopFans(args P) (result ArtistGetTopFans, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetTopFans") }()
	err = callGet("artist.gettopfans", api.creds, args, &result, P{
		"normal": []string{"artist", "mbid", "autocorrect"},
	})
	return
}

//artist.getTopTags
func (api artistApi) GetTopTags(args P) (result ArtistGetTopTags, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetTopTags") }()
	err = callGet("artist.gettoptags", api.creds, args, &result, P{
		"normal": []string{"artist", "mbid", "autocorrect"},
	})
	return
}

//artist.getTopTracks
func (api artistApi) GetTopTracks(args P) (result ArtistGetTopTracks, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetTopTracks") }()
	err = callGet("artist.gettoptracks", api.creds, args, &result, P{
		"normal": []string{"artist", "mbid", "autocorrect", "page", "limit"},
	})
	return
}

//artist.removeTag
func (api artistApi) RemoveTag(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Artist.RemoveTag") }()
	err = callPost("artist.removetag", api.creds, args, nil, P{
		"normal": []string{"artist", "tag"},
	})
	return
}

//artist.search
func (api artistApi) Search(args P) (result ArtistSearch, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.Search") }()
	err = callGet("artist.search", api.creds, args, &result, P{
		"normal": []string{"limit", "page", "artist"},
	})
	return
}

//artist.share
func (api artistApi) Share(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Artist.Share") }()
	err = callPost("artist.share", api.creds, args, nil, P{
		"normal": []string{"artist", "recipient", "message", "public"},
	})
	return
}

//artist.shout
func (api artistApi) Shout(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Artist.Shout") }()
	err = callPost("artist.shout", api.creds, args, nil, P{
		"normal": []string{"artist", "message"},
	})
	return
}
