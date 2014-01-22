package lastfm

type userApi struct {
	params *apiParams
}

//user.getArtistTracks
func (api userApi) GetArtistTracks(args map[string]interface{}) (result UserGetArtistTracks, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetArtistTracks") }()
	err = callGet("user.getartisttracks", api.params, args, &result, P{
		"plain": []string{"user", "artist", "startTimeStamp", "page", "endTimeStamp"},
	})
	return
}

//user.getBannedTracks
func (api userApi) GetBannedTracks(args map[string]interface{}) (result UserGetBannedTracks, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetBanndedTracks") }()
	err = callGet("user.getbannedtracks", api.params, args, &result, P{
		"plain": []string{"user", "limit", "page"},
	})
	return
}

//user.getEvents
func (api userApi) GetEvents(args map[string]interface{}) (result UserGetEvents, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetEvents") }()
	err = callGet("user.getevents", api.params, args, &result, P{
		"plain": []string{"user", "page", "festivalsonly", "limit"},
	})
	return
}

//user.getFriends
func (api userApi) GetFriends(args map[string]interface{}) (result UserGetFriends, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetFriends") }()
	err = callGet("user.getfriends", api.params, args, &result, P{
		"plain": []string{"user", "recenttracks", "limit", "page"},
	})
	return
}

//user.getInfo
func (api userApi) GetInfo(args map[string]interface{}) (result UserGetInfo, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetInfo") }()
	if _, ok := args["user"]; !ok && api.params.sk != "" {
		err = callPost("user.getinfo", api.params, args, &result, P{})
	} else {
		err = callGet("user.getinfo", api.params, args, &result, P{
			"plain": []string{"user"},
		})
	}
	return
}

//user.getLovedTracks
func (api userApi) GetLovedTracks(args map[string]interface{}) (result UserGetLovedTracks, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetLovedTracks") }()
	err = callGet("user.getlovedtracks", api.params, args, &result, P{
		"plain": []string{"user", "limit", "page"},
	})
	return
}

//user.getNeighbours
func (api userApi) GetNeighbours(args map[string]interface{}) (result UserGetNeighbours, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetNeighbours") }()
	err = callGet("user.getneighbours", api.params, args, &result, P{
		"plain": []string{"user", "limit"},
	})
	return
}

//user.getNewReleases
func (api userApi) GetNewReleases(args map[string]interface{}) (result UserGetNewReleases, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetNewReleases") }()
	err = callGet("user.getnewreleases", api.params, args, &result, P{
		"plain": []string{"user", "userrecs"},
	})
	return
}

//user.getPastEvents
func (api userApi) GetPastEvents(args map[string]interface{}) (result UserGetPastEvents, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetPastEvents") }()
	err = callGet("user.getpastevents", api.params, args, &result, P{
		"plain": []string{"user", "page", "limit"},
	})
	return
}

//user.getPersonalTags
func (api userApi) GetPersonalTags(args map[string]interface{}) (result UserGetPersonalTags, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetPersonalTags") }()
	err = callGet("user.getPersonalTags", api.params, args, &result, P{
		"plain": []string{"user", "tag", "taggingtype", "limit", "page"},
	})
	return
}

//user.getPlaylists
func (api userApi) GetPlaylists(args map[string]interface{}) (result UserGetPlaylists, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetPlaylists") }()
	err = callGet("user.getplaylists", api.params, args, &result, P{
		"plain": []string{"user"},
	})
	return
}

//user.getRecentStations (auth required)
func (api userApi) GetRecentStations(args map[string]interface{}) (result UserGetRecentStations, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetRecentStations") }()
	err = callPost("user.getrecentstations", api.params, args, &result, P{ //auth required
		"plain": []string{"user", "limit", "page"},
	})
	return
}

//user.getRecentTracks
func (api userApi) GetRecentTracks(args map[string]interface{}) (result UserGetRecentTracks, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetRecentTracks") }()
	err = callGet("user.getrecenttracks", api.params, args, &result, P{
		"plain": []string{"user", "limit", "page", "from", "extended", "to"},
	})
	return
}

//user.getRecommendedArtists (auth required)
func (api userApi) GetRecommendedArtists(args map[string]interface{}) (result UserGetRecommendedArtists, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetRecommendedArtists") }()
	err = callPost("user.getrecommendedartists", api.params, args, &result, P{
		"plain": []string{"limit", "page"},
	})
	return
}

//user.getRecommendedEvents (auth required)
func (api userApi) GetRecommendedEvents(args map[string]interface{}) (result UserGetRecommendedEvents, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetRecommendedEvents") }()
	err = callPost("user.getrecommendedevents", api.params, args, &result, P{
		"plain": []string{"limit", "page", "latitude", "longtitude", "festivalsonly", "country"},
	})
	return
}

//user.getShouts
func (api userApi) GetShouts(args map[string]interface{}) (result UserGetShouts, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetShouts") }()
	err = callGet("user.getshouts", api.params, args, &result, P{
		"plain": []string{"user", "page"},
	})
	return
}

//user.getTopAlbums
func (api userApi) GetTopAlbums(args map[string]interface{}) (result UserGetTopAlbums, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetTopAlbums") }()
	err = callGet("user.gettopalbums", api.params, args, &result, P{
		"plain": []string{"user", "period", "limit", "page"},
	})
	return
}

//user.getTopArtists
func (api userApi) GetTopArtists(args map[string]interface{}) (result UserGetTopArtists, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetTopArtists") }()
	err = callGet("user.gettopartists", api.params, args, &result, P{
		"plain": []string{"user", "period", "limit", "page"},
	})
	return
}

//user.getTopTags
func (api userApi) GetTopTags(args map[string]interface{}) (result UserGetTopTags, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetTopTags") }()
	err = callGet("user.gettoptags", api.params, args, &result, P{
		"plain": []string{"user", "limit"},
	})
	return
}

//user.getTopTracks
func (api userApi) GetTopTracks(args map[string]interface{}) (result UserGetTopTracks, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetTopTracks") }()
	err = callGet("user.gettoptracks", api.params, args, &result, P{
		"plain": []string{"user", "period", "limit", "page"},
	})
	return
}

//user.getWeeklyAlbumChart
func (api userApi) GetWeeklyAlbumChart(args map[string]interface{}) (result UserGetWeeklyAlbumChart, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetWeeklyAlbumChart") }()
	err = callGet("user.getweeklyalbumchart", api.params, args, &result, P{
		"plain": []string{"user", "from", "to"},
	})
	return
}

//user.getWeeklyArtistChart
func (api userApi) GetWeeklyArtistChart(args map[string]interface{}) (result UserGetWeeklyArtistChart, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetWeeklyArtistChart") }()
	err = callGet("user.getweeklyartistchart", api.params, args, &result, P{
		"plain": []string{"user", "from", "to"},
	})
	return
}

//user.getWeeklyChartList
func (api userApi) GetWeeklyChartList(args map[string]interface{}) (result UserGetWeeklyChartList, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetWeeklyChartList") }()
	err = callGet("user.getweeklychartlist", api.params, args, &result, P{
		"plain": []string{"user"},
	})
	return
}

//user.getWeeklyTrackChart
func (api userApi) GetWeeklyTrackChart(args map[string]interface{}) (result UserGetWeeklyTrackChart, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetWeeklyTrackChart") }()
	err = callGet("user.getweeklytrackchart", api.params, args, &result, P{
		"plain": []string{"user", "from", "to"},
	})
	return
}

//user.Shout
func (api userApi) Shout(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.User.Shout") }()
	err = callPost("user.shout", api.params, args, nil, P{
		"plain": []string{"user", "message"},
	})
	return
}

//user.signUp (deprecated)
//user.terms (deprecated)
