package lastfm

type chartApi struct {
	params *apiParams
}

//chart.getHypedArtists
func (api chartApi) GetHypedArtists(args map[string]interface{}) (result ChartGetHypedArtists, err error) {
	defer func() { appendCaller(err, "lastfm.Chart.GetHypedArtists") }()
	err = callGet("chart.gethypedartists", api.params, args, &result, P{
		"plain": []string{"page", "limit"},
	})
	return
}

//chart.getHypedTracks
func (api chartApi) GetHypedTracks(args map[string]interface{}) (result ChartGetHypedTracks, err error) {
	defer func() { appendCaller(err, "lastfm.Chart.GetHypedTracks") }()
	err = callGet("chart.gethypedtracks", api.params, args, &result, P{
		"plain": []string{"page", "limit"},
	})
	return
}

//chart.getLovedTracks
func (api chartApi) GetLovedTracks(args map[string]interface{}) (result ChartGetLovedTracks, err error) {
	defer func() { appendCaller(err, "lastfm.Chart.GetLovedTracks") }()
	err = callGet("chart.getlovedtracks", api.params, args, &result, P{
		"plain": []string{"page", "limit"},
	})
	return
}

//chart.getTopArtists
func (api chartApi) GetTopArtists(args map[string]interface{}) (result ChartGetTopArtists, err error) {
	defer func() { appendCaller(err, "lastfm.Chart.GetTopArtists") }()
	err = callGet("chart.gettopartists", api.params, args, &result, P{
		"plain": []string{"page", "limit"},
	})
	return
}

//chart.getTopTags
func (api chartApi) GetTopTags(args map[string]interface{}) (result ChartGetTopTags, err error) {
	defer func() { appendCaller(err, "lastfm.Chart.GetTopTags") }()
	err = callGet("chart.gettoptags", api.params, args, &result, P{
		"plain": []string{"page", "limit"},
	})
	return
}

//chart.getTopTracks
func (api chartApi) GetTopTracks(args map[string]interface{}) (result ChartGetTopTracks, err error) {
	defer func() { appendCaller(err, "lastfm.Chart.GetTopTracks") }()
	err = callGet("chart.gettoptracks", api.params, args, &result, P{
		"plain": []string{"page", "limit"},
	})
	return
}
