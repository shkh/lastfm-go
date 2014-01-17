package lastfm

type libraryApi struct {
	creds *credentials
}

//library.addAlbum
func (api libraryApi) AddAlbum(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Library.AddAlbum") }()
	err = callPost("library.addalbum", api.creds, args, nil, P{
		"indexing": []string{"artist", "album"},
	})
	return
}

//library.addArtist
func (api libraryApi) AddArtist(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Library.AddArtist") }()
	err = callPost("library.addartist", api.creds, args, nil, P{
		"indexing": []string{"artist"},
	})
	return
}

//library.addTrack
func (api libraryApi) AddTrack(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Library.AddTrack") }()
	err = callPost("library.addtrack", api.creds, args, nil, P{
		"normal": []string{"artist", "track"},
	})
	return
}

//library.getAlbums
func (api libraryApi) GetAlbums(args P) (result LibraryGetAlbums, err error) {
	defer func() { appendCaller(err, "lastfm.Library.GetAlbums") }()
	err = callGet("library.getalbums", api.creds, args, &result, P{
		"normal": []string{"user", "artist", "limit", "page"},
	})
	return
}

//library.getArtists
func (api libraryApi) GetArtists(args P) (result LibraryGetArtists, err error) {
	defer func() { appendCaller(err, "lastfm.Library.GetArtists") }()
	err = callGet("library.getartists", api.creds, args, &result, P{
		"normal": []string{"user", "limit", "page"},
	})
	return
}

//library.getTracks
func (api libraryApi) GetTracks(args P) (result LibraryGetTracks, err error) {
	defer func() { appendCaller(err, "lastfm.Library.GetTracks") }()
	err = callGet("library.gettracks", api.creds, args, &result, P{
		"normal": []string{"user", "artist", "album", "limit", "page"},
	})
	return
}

//library.removeAlbum
func (api libraryApi) RemoveAllbum(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Library.RemoveAlbum") }()
	err = callPost("library.removealbum", api.creds, args, nil, P{
		"normal": []string{"artist", "album"},
	})
	return
}

//library.removeArtist
func (api libraryApi) RemoveArtist(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Library.RemoveArtist") }()
	err = callPost("library.removeartist", api.creds, args, nil, P{
		"normal": []string{"artist"},
	})
	return
}

//library.removeScrobble
func (api libraryApi) RemoveScrobble(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Library.RemoveScrobble") }()
	err = callPost("library.removescrobble", api.creds, args, nil, P{
		"normal": []string{"artist", "track", "timestamp"},
	})
	return
}

//library.removeTrack
func (api libraryApi) RemoveTrack(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Library.RemoveTrack") }()
	err = callPost("library.removetrack", api.creds, args, nil, P{
		"normal": []string{"artist", "track"},
	})
	return
}
