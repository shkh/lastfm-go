package lastfm

type playlistApi struct {
    creds *credentials
}

//playlist.addTrack
func (api playlistApi) AddTrack (args P) (err error) {
    defer func () { appendCaller (err, "lastfm.Playlist.AddTrack") } ()
	err = callPost("playlist.addtrack", api.creds, args, nil, P{
		"normal": []string{"playlistID", "track", "artist"},
	})
    return
}

//playlist.create
func (api playlistApi) Create (args P) (err error) {
    defer func () { appendCaller (err, "lastfm.Playlist.Create") } ()
	err = callPost("playlist.create", api.creds, args, nil, P{
		"normal": []string{"title", "description"},
	})
    return
}
