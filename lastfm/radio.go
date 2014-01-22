package lastfm

type radioApi struct {
	params *apiParams
}

//radio.getPlaylist (auth required)
func (api radioApi) GetPlaylist(args map[string]interface{}) (result RadioGetPlaylist, err error) {
	defer func() { appendCaller(err, "lastfm.Radio.GetPlaylist") }()
	err = callPost("radio.getplaylist", api.params, args, &result, P{
		"plain": []string{"discovery", "buylinks", "speed_multiplier", "bitrate", "rtp"},
	})
	return
}

//radio.search
func (api radioApi) Search(args map[string]interface{}) (result RadioSearch, err error) {
	defer func() { appendCaller(err, "lastfm.Radio.Search") }()
	err = callGet("radio.search", api.params, args, &result, P{
		"plain": []string{"name"},
	})
	return
}

//radio.tune
func (api radioApi) Tune(args map[string]interface{}) (result RadioTune, err error) {
	defer func() { appendCaller(err, "lastfm.Radio.Tune") }()
	err = callPost("radio.tune", api.params, args, &result, P{
		"plain": []string{"lang", "station"},
	})
	return
}
