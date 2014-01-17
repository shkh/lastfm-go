package lastfm

type groupApi struct {
	creds *credentials
}

//group.getHype
func (api groupApi) GetHype(args P) (result GroupGetHype, err error) {
	defer func() { appendCaller(err, "lastfm.Group.GetHype") }()
	err = callGet("group.gethype", api.creds, args, &result, P{
		"normal": []string{"group"},
	})
	return
}

//group.getMembers
func (api groupApi) GetMembers(args P) (result GroupGetMembers, err error) {
	defer func() { appendCaller(err, "lastfm.Group.GetMembers") }()
	err = callGet("group.getmembers", api.creds, args, &result, P{
		"normal": []string{"group", "page", "limit"},
	})
	return
}

//group.getWeeklyAlbumChart
func (api groupApi) GetWeeklyAlbumChart(args P) (result GroupGetWeeklyAlbumChart, err error) {
	defer func() { appendCaller(err, "lastfm.Group.GetWeeklyAlbumChart") }()
	err = callGet("group.weeklyalbumchart", api.creds, args, &result, P{
		"normal": []string{"group", "from", "to"},
	})
	return
}

//group.getWeeklyArtistChart
func (api groupApi) GetWeeklyArtistChart(args P) (result GroupGetWeeklyArtistChart, err error) {
	defer func() { appendCaller(err, "lastfm.Group.GetWeeklyArtistChart") }()
	err = callGet("group.weeklyartistchart", api.creds, args, &result, P{
		"normal": []string{"group", "from", "to"},
	})
	return
}

//group.getWeeklyChartList
func (api groupApi) GetWeeklyChartlist(args P) (result GroupGetWeeklyChartlist, err error) {
	defer func() { appendCaller(err, "lastfm.Group.GetWeeklyChartlist") }()
	err = callGet("group.weeklychartlist", api.creds, args, &result, P{
		"normal": []string{"group"},
	})
	return
}

//group.getWeeklyTrackChart
func (api groupApi) GetWeeklyTrackChart(args P) (result GroupGetWeeklyTrackChart, err error) {
	defer func() { appendCaller(err, "lastfm.Group.GetWeeklyTrackChart") }()
	err = callGet("group.weeklytrackchart", api.creds, args, &result, P{
		"normal": []string{"group", "from", "to"},
	})
	return
}
