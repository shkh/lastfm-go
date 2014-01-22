package lastfm

type libraryApi struct {
	params *apiParams
}

//library.addAlbum
func (api libraryApi) AddAlbum(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Library.AddAlbum") }()
	err = callPost("library.addalbum", api.params, args, nil, P{
		"indexing": []string{"artist", "album"},
	})
	return
}

//library.addArtist
func (api libraryApi) AddArtist(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Library.AddArtist") }()
	err = callPost("library.addartist", api.params, args, nil, P{
		"indexing": []string{"artist"},
	})
	return
}

//library.addTrack
func (api libraryApi) AddTrack(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Library.AddTrack") }()
	err = callPost("library.addtrack", api.params, args, nil, P{
		"plain": []string{"artist", "track"},
	})
	return
}

//library.getAlbums
func (api libraryApi) GetAlbums(args map[string]interface{}) (result LibraryGetAlbums, err error) {
	defer func() { appendCaller(err, "lastfm.Library.GetAlbums") }()
	err = callGet("library.getalbums", api.params, args, &result, P{
		"plain": []string{"user", "artist", "limit", "page"},
	})
	return
}

//library.getArtists
func (api libraryApi) GetArtists(args map[string]interface{}) (result LibraryGetArtists, err error) {
	defer func() { appendCaller(err, "lastfm.Library.GetArtists") }()
	err = callGet("library.getartists", api.params, args, &result, P{
		"plain": []string{"user", "limit", "page"},
	})
	return
}

//library.getTracks
func (api libraryApi) GetTracks(args map[string]interface{}) (result LibraryGetTracks, err error) {
	defer func() { appendCaller(err, "lastfm.Library.GetTracks") }()
	err = callGet("library.gettracks", api.params, args, &result, P{
		"plain": []string{"user", "artist", "album", "limit", "page"},
	})
	return
}

//library.removeAlbum
func (api libraryApi) RemoveAllbum(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Library.RemoveAlbum") }()
	err = callPost("library.removealbum", api.params, args, nil, P{
		"plain": []string{"artist", "album"},
	})
	return
}

//library.removeArtist
func (api libraryApi) RemoveArtist(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Library.RemoveArtist") }()
	err = callPost("library.removeartist", api.params, args, nil, P{
		"plain": []string{"artist"},
	})
	return
}

//library.removeScrobble
func (api libraryApi) RemoveScrobble(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Library.RemoveScrobble") }()
	err = callPost("library.removescrobble", api.params, args, nil, P{
		"plain": []string{"artist", "track", "timestamp"},
	})
	return
}

//library.removeTrack
func (api libraryApi) RemoveTrack(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Library.RemoveTrack") }()
	err = callPost("library.removetrack", api.params, args, nil, P{
		"plain": []string{"artist", "track"},
	})
	return
}
