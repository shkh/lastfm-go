package lastfm

type tagApi struct {
    creds *credentials
}

//tag.getInfo
func (api tagApi) GetInfo (args P) (result TagGetInfo, err error) {
	defer func() { appendCaller(err, "lastfm.Tag.GetInfo") }()
	err = callGet("tag.getinfo", api.creds, args, &result, P{
		"normal": []string{"lang", "artist", "mbid"},
	})
	return
}

//tag.getSimilar
func (api tagApi) GetSimilar (args P) (result TagGetSimilar, err error) {
	defer func() { appendCaller(err, "lastfm.Tag.GetSimilar") }()
	err = callGet("tag.getsimilar", api.creds, args, &result, P{
		"normal": []string{"tag"},
	})
	return
}

//tag.getTopAlbums
func (api tagApi) GetTopAlbums (args P) (result TagGetTopAlbums, err error) {
	defer func() { appendCaller(err, "lastfm.Tag.GetTopAlbums") }()
	err = callGet("tag.gettopalbums", api.creds, args, &result, P{
		"normal": []string{"tag", "limit", "page"},
	})
	return
}

//tag.getTopArtists
func (api tagApi) GetTopArtists (args P) (result TagGetTopArtists, err error) {
	defer func() { appendCaller(err, "lastfm.Tag.GetTopArtists") }()
	err = callGet("tag.gettopartists", api.creds, args, &result, P{
		"normal": []string{"tag", "limit", "page"},
	})
	return
}

//tag.getTopTags
func (api tagApi) GetTopTags (args P) (result TagGetTopTags, err error) {
	defer func() { appendCaller(err, "lastfm.Tag.GetTopTags") }()
	err = callGet("tag.gettoptags", api.creds, args, &result, P{
	})
	return
}

//tag.getTopTracks
func (api tagApi) GetTopTracks (args P) (result TagGetTopTracks, err error) {
	defer func() { appendCaller(err, "lastfm.Tag.GetTopTracks") }()
	err = callGet("tag.gettoptracks", api.creds, args, &result, P{
		"normal": []string{"tag", "limit", "page"},
	})
	return
}

//tag.getWeeklyArtistChart
func (api tagApi) GetTopWeeklyArtistChart (args P) (result TagGetWeeklyArtistChart, err error) {
	defer func() { appendCaller(err, "lastfm.Tag.GetWeeklyArtistChart") }()
	err = callGet("tag.getweeklyartistchart", api.creds, args, &result, P{
		"normal": []string{"tag", "from", "to", "limit"},
	})
	return
}

//tag.getWeeklyChartList
func (api tagApi) GetTopWeeklyChartList (args P) (result TagGetWeeklyChartList, err error) {
	defer func() { appendCaller(err, "lastfm.Tag.GetWeeklyChartList") }()
	err = callGet("tag.getweeklychartlist", api.creds, args, &result, P{
		"normal": []string{"tag"},
	})
	return
}

//tag.search
func (api tagApi) Search (args P) (result TagSearch, err error) {
	defer func() { appendCaller(err, "lastfm.Tag.Search") }()
	err = callGet("tag.search", api.creds, args, &result, P{
		"normal": []string{"tag", "limit", "page"},
	})
	return
}
