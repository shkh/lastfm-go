package lastfm

type radioApi struct {
	creds *credentials
}

//radio.getPlaylist
func (api groupApi) GetPlaylist(args P) (result RadioGetPlaylist, err error) {
	defer func() { appendCaller(err, "lastfm.Radio.GetPlaylist") }()
	err = callGet("radio.getplaylist", api.creds, args, &result, P{
		"normal": []string{"discovery", "buylinks", "speed_multiplier", ""},
	})
	return
}

//radio.search
//radio.tune
