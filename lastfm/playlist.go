package lastfm

type playlistApi struct {
	params *apiParams
}

//playlist.addTrack
func (api playlistApi) AddTrack(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Playlist.AddTrack") }()
	err = callPost("playlist.addtrack", api.params, args, nil, P{
		"plain": []string{"playlistID", "track", "artist"},
	})
	return
}

//playlist.create
func (api playlistApi) Create(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Playlist.Create") }()
	err = callPost("playlist.create", api.params, args, nil, P{
		"plain": []string{"title", "description"},
	})
	return
}
