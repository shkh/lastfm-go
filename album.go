package lastfm

type albumApi struct {
	creds *credentials
}

//album.addTags
func (api albumApi) AddTags(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Album.AddTags") }()
	if err = requireAuth(api.creds); err != nil {
		return
	}
	err = callPost("album.addtags", api.creds, args, nil, P{
		"normal": []string{"artist", "album", "tags"},
	})
	return
}

//album.getBuylinks
func (api albumApi) GetBuyLinks(args P) (result AlbumGetBuyLinks, err error) {
	defer func() { appendCaller(err, "lastfm.Album.GetBuyLinks") }()
	err = callGet("album.getbuylinks", api.creds, args, &result, P{
		"normal": []string{"artist", "album", "mbid", "autocorrect", "country"},
	})
	return
}

//album.getInfo
func (api albumApi) GetInfo(args P) (result AlbumGetInfo, err error) {
	defer func() { appendCaller(err, "lastfm.Album.GetInfo") }()
	err = callGet("album.getinfo", api.creds, args, &result, P{
		"normal": []string{"artist", "album", "mbid", "autocorrect", "username", "lang"},
	})
	return
}

//album.getShouts
func (api albumApi) GetShouts(args P) (result AlbumGetShouts, err error) {
	defer func() { appendCaller(err, "lastfm.Album.GetShouts") }()
	err = callGet("album.getshouts", api.creds, args, &result, P{
		"normal": []string{"artist", "album", "mbid", "limit", "autocorrect", "page"},
	})
	return
}

//album.getTags
func (api albumApi) GetTags(args P) (result AlbumGetTags, err error) {
	defer func() { appendCaller(err, "lastfm.Album.GetTags") }()
	if _, ok := args["user"]; !ok {
		if err = requireAuth(api.creds); err != nil {
			return
		}
		args["user"] = api.creds.username
	}
	err = callGet("album.gettags", api.creds, args, &result, P{
		"normal": []string{"artist", "album", "mbid", "autocorrect", "user"},
	})
	return
}

//album.getTopTags
func (api albumApi) GetTopTags(args P) (result AlbumGetTopTags, err error) {
	defer func() { appendCaller(err, "lastfm.Album.GetTopTags") }()
	err = callGet("album.gettoptags", api.creds, args, &result, P{
		"normal": []string{"artist", "album", "autocorrect", "mbid"},
	})
	return
}

//album.removeTag
func (api albumApi) RemoveTag(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Album.RemoveTag") }()
	if err = requireAuth(api.creds); err != nil {
		return
	}
	err = callPost("album.removetag", api.creds, args, nil, P{
		"normal": []string{"artist", "album", "tag"},
	})
	return
}

//album.search
func (api albumApi) Search(args P) (result AlbumSearch, err error) {
	defer func() { appendCaller(err, "lastfm.Album.Search") }()
	err = callGet("album.search", api.creds, args, &result, P{
		"normal": []string{"limit", "page", "album"},
	})
	return
}

//album.share
func (api albumApi) Share(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Album.Share") }()
	if err = requireAuth(api.creds); err != nil {
		return
	}
	err = callPost("album.share", api.creds, args, nil, P{
		"normal": []string{"artist", "album", "public", "message", "recipient"},
	})
	return
}
