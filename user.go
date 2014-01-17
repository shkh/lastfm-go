package lastfm

type userApi struct {
	creds *credentials
}

//user.getArtistTracks
func (api userApi) GetArtistTracks(args P) (result UserGetArtistTracks, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetArtistTracks") }()
	err = callGet("user.getArtistTracks", api.creds, args, &result, P{
		"normal": []string{"user", "artist", "startTimeStamp", "page", "endTimeStamp"},
	})
	return
}

//user.getBannedTracks
func (api userApi) GetBannedTracks(args P) (result UserGetBannedTracks, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetBanndedTracks") }()
	err = callGet("user.getbannedtracks", api.creds, args, &result, P{
		"normal": []string{"user", "limit", "page"},
	})
	return
}

//user.getEvents
func (api userApi) GetEvents(args P) (result UserGetEvents, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetEvents") }()
	err = callGet("user.getevents", api.creds, args, &result, P{
		"normal": []string{"user", "page", "festivalsonly", "limit"},
	})
	return
}

//user.getFriends
func (api userApi) GetFriends(args P) (result UserGetFriends, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetFriends") }()
	err = callGet("user.getfriends", api.creds, args, &result, P{
		"normal": []string{"user", "recenttracks", "limit", "page"},
	})
	return
}

//user.getInfo
func (api userApi) GetInfo(args P) (result UserGetInfo, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetInfo") }()
	if _, ok := args["user"]; !ok {
		args = P{"user": api.creds.username}
	}
	err = callGet("user.getinfo", api.creds, args, &result, P{
		"normal": []string{"user"},
	})
	return
}

//user.getLovedTracks
func (api userApi) GetLovedTracks(args P) (result UserGetLovedTracks, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetLovedTracks") }()
	err = callGet("user.getlovedtracks", api.creds, args, &result, P{
		"normal": []string{"user", "limit", "page"},
	})
	return
}

//user.getNeighbours
func (api userApi) GetNeighbours(args P) (result UserGetNeighbours, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetNeighbours") }()
	err = callGet("user.getneighbers", api.creds, args, &result, P{
		"normal": []string{"user", "limit"},
	})
	return
}

//user.getNewReleases
func (api userApi) GetNewReleases(args P) (result UserGetNewReleases, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetNewReleases") }()
	err = callGet("user.getnewreleases", api.creds, args, &result, P{
		"normal": []string{"user", "userrecs"},
	})
	return
}

//user.getPastEvents
func (api userApi) GetPastEvents(args P) (result UserGetPastEvents, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetPastEvents") }()
	err = callGet("user.getpastevents", api.creds, args, &result, P{
		"normal": []string{"user", "page", "limit"},
	})
	return
}

//user.getPersonalTags
func (api userApi) GetPersonalTags(args P) (result UserGetPersonalTags, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetPersonalTags") }()
	err = callGet("user.getPersonalTags", api.creds, args, &result, P{
		"normal": []string{"user", "tag", "taggingtype", "limit", "page"},
	})
	return
}

//user.getPlaylists
func (api userApi) GetPlaylists(args P) (result UserGetPlaylists, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetPlaylists") }()
	err = callGet("user.getplaylists", api.creds, args, &result, P{
		"normal": []string{"user"},
	})
	return
}

//user.getRecentStations (auth required)
func (api userApi) GetRecentStations(args P) (result UserGetRecentStations, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetRecentStations") }()
	err = callPost("user.getrecentstations", api.creds, args, &result, P{ //auth required
		"normal": []string{"user", "limit", "page"},
	})
	return
}

//user.getRecentTracks
func (api userApi) GetRecentTracks(args P) (result UserGetRecentTracks, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetRecentTracks") }()
	err = callGet("user.getrecenttracks", api.creds, args, &result, P{
		"normal": []string{"user", "limit", "page", "from", "extended", "to"},
	})
	return
}

//user.getRecommendedArtists (auth required)
func (api userApi) GetRecommendedArtists(args P) (result UserGetRecommendedArtists, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetRecommendedArtists") }()
	err = callPost("user.getrecommendedartists", api.creds, args, &result, P{
		"normal": []string{"limit", "page"},
	})
	return
}

//user.getRecommendedEvents (auth required)
func (api userApi) GetRecommendedEvents(args P) (result UserGetRecommendedEvents, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetRecommendedEvents") }()
	err = callPost("user.getrecommendedevents", api.creds, args, &result, P{
		"normal": []string{"limit", "page", "latitude", "longtitude", "festivalsonly", "country"},
	})
	return
}

//user.getShouts
func (api userApi) GetShouts(args P) (result UserGetShouts, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetShouts") }()
	err = callGet("user.getshouts", api.creds, args, &result, P{
		"normal": []string{"user", "page"},
	})
	return
}

//user.getTopAlbums
func (api userApi) GetTopAlbums(args P) (result UserGetTopAlbums, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetTopAlbums") }()
	err = callGet("user.gettopalbums", api.creds, args, &result, P{
		"normal": []string{"user", "period", "limit", "page"},
	})
	return
}

//user.getTopArtists
func (api userApi) GetTopArtists(args P) (result UserGetTopArtists, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetTopArtists") }()
	err = callGet("user.gettopartists", api.creds, args, &result, P{
		"normal": []string{"user", "period", "limit", "page"},
	})
	return
}

//user.getTopTags
func (api userApi) GetTopTags(args P) (result UserGetTopTags, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetTopTags") }()
	err = callGet("user.gettoptags", api.creds, args, &result, P{
		"normal": []string{"user", "limit"},
	})
	return
}

//user.getTopTracks
func (api userApi) GetTopTracks(args P) (result UserGetTopTracks, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetTopTracks") }()
	err = callGet("user.gettoptracks", api.creds, args, &result, P{
		"normal": []string{"user", "period", "limit", "page"},
	})
	return
}

//user.getWeeklyAlbumChart
func (api userApi) GetWeeklyAlbumChart(args P) (result UserGetWeeklyAlbumChart, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetWeeklyAlbumChart") }()
	err = callGet("user.getweeklyalbumchart", api.creds, args, &result, P{
		"normal": []string{"user", "from", "to"},
	})
	return
}

//user.getWeeklyArtistChart
func (api userApi) GetWeeklyArtistChart(args P) (result UserGetWeeklyArtistChart, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetWeeklyArtistChart") }()
	err = callGet("user.getweeklyartistchart", api.creds, args, &result, P{
		"normal": []string{"user", "from", "to"},
	})
	return
}

//user.getWeeklyChartList
func (api userApi) GetWeeklyChartList(args P) (result UserGetWeeklyChartList, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetWeeklyChartList") }()
	err = callGet("user.getweeklychartlist", api.creds, args, &result, P{
		"normal": []string{"user"},
	})
	return
}

//user.getWeeklyTrackChart
func (api userApi) GetWeeklyTrackChart(args P) (result UserGetWeeklyTrackChart, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetWeeklyTrackChart") }()
	err = callGet("user.getweeklytrackchart", api.creds, args, &result, P{
		"normal": []string{"user", "from", "to"},
	})
	return
}

//user.Shout
func (api userApi) Shout(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.User.Shout") }()
	err = callPost("user.shout", api.creds, args, nil, P{
		"normal": []string{"user", "message"},
	})
	return
}

//user.signUp (deprecated)
//user.terms (deprecated)
