package lastfm

type geoApi struct {
	params *apiParams
}

//geo.getEvents
func (api geoApi) GetEvents(args map[string]interface{}) (result GeoGetEvents, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetEvents") }()
	err = callGet("geo.getevents", api.params, args, &result, P{
		"plain": []string{"long", "lat", "location", "distance", "page", "tag", "festivalsonly", "limit"},
	})
	return
}

//geo.getMetroArtistChart
func (api geoApi) GetMetroArtistChart(args map[string]interface{}) (result GeoGetMetroArtistChart, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetMetroArtistChart") }()
	err = callGet("geo.getmetroartistchart", api.params, args, &result, P{
		"plain": []string{"metro", "country", "start", "end", "page", "limit"},
	})
	return
}

//geo.getMetroHypeArtistChart
func (api geoApi) GetMetroHypeArtistChart(args map[string]interface{}) (result GeoGetMetroHypeArtistChart, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetMetroHypeArtistChart") }()
	err = callGet("geo.getmetrohypeartistchart", api.params, args, &result, P{
		"plain": []string{"metro", "country", "start", "end", "page", "limit"},
	})
	return
}

//geo.getMetroHypeTrackChart
func (api geoApi) GetMetroHypeTrackChart(args map[string]interface{}) (result GeoGetMetroHypeTrackChart, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetMetroHypeTrackChart") }()
	err = callGet("geo.getmetrohypetrackchart", api.params, args, &result, P{
		"plain": []string{"metro", "country", "start", "end", "page", "limit"},
	})
	return
}

//geo.getMetroTrackChart
func (api geoApi) GetMetroTrackChart(args map[string]interface{}) (result GeoGetMetroTrackChart, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetMetroTrackChart") }()
	err = callGet("geo.getmetrotrackchart", api.params, args, &result, P{
		"plain": []string{"metro", "country", "start", "end", "page", "limit"},
	})
	return
}

//geo.getMetroUniqueArtistChart
func (api geoApi) GetMetroUniqueArtistChart(args map[string]interface{}) (result GeoGetMetroUniqueArtistChart, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetMetroUniqueArtistChart") }()
	err = callGet("geo.getmetrouniqueartistchart", api.params, args, &result, P{
		"plain": []string{"metro", "country", "start", "end", "page", "limit"},
	})
	return
}

//geo.getMetroUniqueTrackChart
func (api geoApi) GetMetroUniqueChartChart(args map[string]interface{}) (result GeoGetMetroUniqueTrackChart, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetMetroUniqueTrackChart") }()
	err = callGet("geo.getmetrouniquetrackchart", api.params, args, &result, P{
		"plain": []string{"metro", "country", "start", "end", "page", "limit"},
	})
	return
}

//geo.getMetroWeeklyChartlist
func (api geoApi) GetMetroWeeklyChartlist(args map[string]interface{}) (result GeoGetMetroWeeklyChartlist, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetMetroWeeklyChartlist") }()
	err = callGet("geo.getmetroweeklychartlist", api.params, args, &result, P{
		"plain": []string{"metro"},
	})
	return
}

//geo.getMetros
func (api geoApi) GetMetros(args map[string]interface{}) (result GeoGetMetros, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetMetros") }()
	err = callGet("geo.getmetros", api.params, args, &result, P{
		"plain": []string{"country"},
	})
	return
}

//geo.getTopArtists
func (api geoApi) GetTopArtists(args map[string]interface{}) (result GeoGetTopArtists, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetTopArtists") }()
	err = callGet("geo.gettopartists", api.params, args, &result, P{
		"plain": []string{"country", "limit", "page"},
	})
	return
}

//geo.getTopTracks
func (api geoApi) GetTopTracks(args map[string]interface{}) (result GeoGetTopTracks, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetTopTracks") }()
	err = callGet("geo.gettoptracks", api.params, args, &result, P{
		"plain": []string{"country", "location", "limit", "page"},
	})
	return
}
