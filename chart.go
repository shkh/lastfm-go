package lastfm

type chartApi struct {
	creds *credentials
}

//chart.getHypedArtists
func (api chartApi) GetHypedArtists(args P) (result ChartGetHypedArtists, err error) {
	defer func() { appendCaller(err, "lastfm.Chart.GetHypedArtists") }()
	err = callGet("chart.gethypedartists", api.creds, args, &result, P{
		"normal": []string{"page", "limit"},
	})
	return
}

//chart.getHypedTracks
func (api chartApi) GetHypedTracks(args P) (result ChartGetHypedTracks, err error) {
	defer func() { appendCaller(err, "lastfm.Chart.GetHypedTracks") }()
	err = callGet("chart.gethypedtracks", api.creds, args, &result, P{
		"normal": []string{"page", "limit"},
	})
	return
}

//chart.getLovedTracks
func (api chartApi) GetLovedTracks(args P) (result ChartGetLovedTracks, err error) {
	defer func() { appendCaller(err, "lastfm.Chart.GetLovedTracks") }()
	err = callGet("chart.getlovedtracks", api.creds, args, &result, P{
		"normal": []string{"page", "limit"},
	})
	return
}

//chart.getTopArtists
func (api chartApi) GetTopArtists(args P) (result ChartGetTopArtists, err error) {
	defer func() { appendCaller(err, "lastfm.Chart.GetTopArtists") }()
	err = callGet("chart.gettopartists", api.creds, args, &result, P{
		"normal": []string{"page", "limit"},
	})
	return
}

//chart.getTopTags
func (api chartApi) GetTopTags(args P) (result ChartGetTopTags, err error) {
	defer func() { appendCaller(err, "lastfm.Chart.GetTopTags") }()
	err = callGet("chart.gettoptags", api.creds, args, &result, P{
		"normal": []string{"page", "limit"},
	})
	return
}

//chart.getTopTracks
func (api chartApi) GetTopTracks(args P) (result ChartGetTopTracks, err error) {
	defer func() { appendCaller(err, "lastfm.Chart.GetTopTracks") }()
	err = callGet("chart.gettoptracks", api.creds, args, &result, P{
		"normal": []string{"page", "limit"},
	})
	return
}
