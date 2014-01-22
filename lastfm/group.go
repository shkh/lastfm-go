package lastfm

type groupApi struct {
	params *apiParams
}

//group.getHype
func (api groupApi) GetHype(args map[string]interface{}) (result GroupGetHype, err error) {
	defer func() { appendCaller(err, "lastfm.Group.GetHype") }()
	err = callGet("group.gethype", api.params, args, &result, P{
		"plain": []string{"group"},
	})
	return
}

//group.getMembers
func (api groupApi) GetMembers(args map[string]interface{}) (result GroupGetMembers, err error) {
	defer func() { appendCaller(err, "lastfm.Group.GetMembers") }()
	err = callGet("group.getmembers", api.params, args, &result, P{
		"plain": []string{"group", "page", "limit"},
	})
	return
}

//group.getWeeklyAlbumChart
func (api groupApi) GetWeeklyAlbumChart(args map[string]interface{}) (result GroupGetWeeklyAlbumChart, err error) {
	defer func() { appendCaller(err, "lastfm.Group.GetWeeklyAlbumChart") }()
	err = callGet("group.weeklyalbumchart", api.params, args, &result, P{
		"plain": []string{"group", "from", "to"},
	})
	return
}

//group.getWeeklyArtistChart
func (api groupApi) GetWeeklyArtistChart(args map[string]interface{}) (result GroupGetWeeklyArtistChart, err error) {
	defer func() { appendCaller(err, "lastfm.Group.GetWeeklyArtistChart") }()
	err = callGet("group.weeklyartistchart", api.params, args, &result, P{
		"plain": []string{"group", "from", "to"},
	})
	return
}

//group.getWeeklyChartList
func (api groupApi) GetWeeklyChartlist(args map[string]interface{}) (result GroupGetWeeklyChartlist, err error) {
	defer func() { appendCaller(err, "lastfm.Group.GetWeeklyChartlist") }()
	err = callGet("group.weeklychartlist", api.params, args, &result, P{
		"plain": []string{"group"},
	})
	return
}

//group.getWeeklyTrackChart
func (api groupApi) GetWeeklyTrackChart(args map[string]interface{}) (result GroupGetWeeklyTrackChart, err error) {
	defer func() { appendCaller(err, "lastfm.Group.GetWeeklyTrackChart") }()
	err = callGet("group.weeklytrackchart", api.params, args, &result, P{
		"plain": []string{"group", "from", "to"},
	})
	return
}
