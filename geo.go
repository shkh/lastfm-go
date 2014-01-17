package lastfm

type geoApi struct {
	creds *credentials
}

//geo.getEvents
func (api geoApi) GetEvents(args P) (result GeoGetEvents, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetEvents") }()
	err = callGet("geo.getevents", api.creds, args, &result, P{
		"normal": []string{"long", "lat", "location", "distance", "page", "tag", "festivalsonly", "limit"},
	})
	return
}

//geo.getMetroArtistChart
func (api geoApi) GetMetroArtistChart(args P) (result GeoGetMetroArtistChart, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetMetroArtistChart") }()
	err = callGet("geo.getmetroartistchart", api.creds, args, &result, P{
		"normal": []string{"metro", "country", "start", "end", "page", "limit"},
	})
	return
}

//geo.getMetroHypeArtistChart
func (api geoApi) GetMetroHypeArtistChart(args P) (result GeoGetMetroHypeArtistChart, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetMetroHypeArtistChart") }()
	err = callGet("geo.getmetrohypeartistchart", api.creds, args, &result, P{
		"normal": []string{"metro", "country", "start", "end", "page", "limit"},
	})
	return
}

//geo.getMetroHypeTrackChart
func (api geoApi) GetMetroHypeTrackChart(args P) (result GeoGetMetroHypeTrackChart, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetMetroHypeTrackChart") }()
	err = callGet("geo.getmetrohypetrackchart", api.creds, args, &result, P{
		"normal": []string{"metro", "country", "start", "end", "page", "limit"},
	})
	return
}

//geo.getMetroTrackChart
func (api geoApi) GetMetroTrackChart(args P) (result GeoGetMetroTrackChart, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetMetroTrackChart") }()
	err = callGet("geo.getmetrotrackchart", api.creds, args, &result, P{
		"normal": []string{"metro", "country", "start", "end", "page", "limit"},
	})
	return
}

//geo.getMetroUniqueArtistChart
func (api geoApi) GetMetroUniqueArtistChart(args P) (result GeoGetMetroUniqueArtistChart, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetMetroUniqueArtistChart") }()
	err = callGet("geo.getmetrouniqueartistchart", api.creds, args, &result, P{
		"normal": []string{"metro", "country", "start", "end", "page", "limit"},
	})
	return
}

//geo.getMetroUniqueTrackChart
func (api geoApi) GetMetroUniqueChartChart(args P) (result GeoGetMetroUniqueTrackChart, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetMetroUniqueTrackChart") }()
	err = callGet("geo.getmetrouniquetrackchart", api.creds, args, &result, P{
		"normal": []string{"metro", "country", "start", "end", "page", "limit"},
	})
	return
}

//geo.getMetroWeeklyChartlist
func (api geoApi) GetMetroWeeklyChartlist(args P) (result GeoGetMetroWeeklyChartlist, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetMetroWeeklyChartlist") }()
	err = callGet("geo.getmetroweeklychartlist", api.creds, args, &result, P{
		"normal": []string{"metro"},
	})
	return
}

//geo.getMetros
func (api geoApi) GetMetros(args P) (result GeoGetMetros, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetMetros") }()
	err = callGet("geo.getmetros", api.creds, args, &result, P{
		"normal": []string{"country"},
	})
	return
}

//geo.getTopArtists
func (api geoApi) GetTopArtists(args P) (result GeoGetTopArtists, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetTopArtists") }()
	err = callGet("geo.gettopartists", api.creds, args, &result, P{
		"normal": []string{"country", "limit", "page"},
	})
	return
}

//geo.getTopTracks
func (api geoApi) GetTopTracks(args P) (result GeoGetTopTracks, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetTopTracks") }()
	err = callGet("geo.gettoptracks", api.creds, args, &result, P{
		"normal": []string{"country", "location", "limit", "page"},
	})
	return
}
