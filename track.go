package lastfm

type trackApi struct {
	creds *credentials
}

//track.addTags
func (api trackApi) AddTags(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.AddTags") }()
	err = callPost("track.addtags", api.creds, args, nil, P{
		"normal": []string{"artist", "track", "tags"},
	})
	return
}

//track.ban
func (api trackApi) Ban(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.Ban") }()
	err = callPost("track.ban", api.creds, args, nil, P{
		"normal": []string{"artist", "track"},
	})
	return
}

//track.getBuylinks
func (api trackApi) GetBuylinks(args P) (result TrackGetBuylinks, err error) {
	defer func() { appendCaller(err, "lastfm.Track.GetBuylinks") }()
	err = callGet("track.getbuylinks", api.creds, args, &result, P{
		"normal": []string{"artist", "track", "mbid", "autocorrect", "country"},
	})
	return
}

//track.getCorrection
func (api trackApi) GetCorrection(args P) (result TrackGetCorrection, err error) {
	defer func() { appendCaller(err, "lastfm.Track.Correction") }()
	err = callGet("track.getcorrection", api.creds, args, &result, P{
		"normal": []string{"artist", "track"},
	})
	return
}

//track.getFingerprintMetadata
func (api trackApi) GetFingerprintMetadata(args P) (result TrackGetFingerprintMetadata, err error) {
	defer func() { appendCaller(err, "lastfm.Track.Correction") }()
	err = callGet("track.getfingerprintmetadata", api.creds, args, &result, P{
		"normal": []string{"fingerprintid"},
	})
	return
}

//track.getInfo
func (api trackApi) GetInfo(args P) (result TrackGetInfo, err error) {
	defer func() { appendCaller(err, "lastfm.Track.GetInfo") }()
	err = callGet("track.getinfo", api.creds, args, &result, P{
		"normal": []string{"artist", "track", "mbid", "username", "autocorrect"},
	})
	return
}

//track.getShouts
func (api trackApi) GetShouts(args P) (result TrackGetShouts, err error) {
	defer func() { appendCaller(err, "lastfm.Track.GetShouts") }()
	err = callGet("track.getshouts", api.creds, args, &result, P{
		"normal": []string{"artist", "track", "mbid", "limit", "page", "autocorrect"},
	})
	return
}

//track.getSimilar
func (api trackApi) GetSimilar(args P) (result TrackGetSimilar, err error) {
	defer func() { appendCaller(err, "lastfm.Track.GetSimilar") }()
	err = callGet("track.getsimilar", api.creds, args, &result, P{
		"normal": []string{"artist", "track", "mbid", "limit", "autocorrect"},
	})
	return
}

//track.getTags
func (api trackApi) GetTags(args P) (result TrackGetTags, err error) {
	defer func() { appendCaller(err, "lastfm.Track.GetTags") }()
	if _, ok := args["user"]; !ok {
		args["user"] = api.creds.username
	}
	err = callGet("track.gettags", api.creds, args, &result, P{
		"normal": []string{"artist", "track", "mbid", "user", "autocorrect"},
	})
	return
}

//track.getTopFans
func (api trackApi) GetTopFans(args P) (result TrackGetTopFans, err error) {
	defer func() { appendCaller(err, "lastfm.Track.GetTopFans") }()
	err = callGet("track.gettopfans", api.creds, args, &result, P{
		"normal": []string{"artist", "track", "mbid", "autocorrect"},
	})
	return
}

//track.getTopTags
func (api trackApi) GetTopTags(args P) (result TrackGetTopTags, err error) {
	defer func() { appendCaller(err, "lastfm.Track.GetTopTags") }()
	err = callGet("track.gettoptags", api.creds, args, &result, P{
		"normal": []string{"artist", "track", "mbid", "autocorrect"},
	})
	return
}

//track.love
func (api trackApi) Love(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.Love") }()
	err = callPost("track.love", api.creds, args, nil, P{
		"normal": []string{"artist", "track"},
	})
	return
}

//track.removeTag
func (api trackApi) RemoveTag(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.RemoveTag") }()
	err = callPost("track.removetag", api.creds, args, nil, P{
		"normal": []string{"artist", "track", "tag"},
	})
	return
}

//track.scrobble
func (api trackApi) Scrobble(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.Scrobble") }()
	err = callPost("track.scrobble", api.creds, args, nil, P{
		"indexing": []string{"artist", "track", "timestamp", "album", "context", "streamId", "chosenByUser", "trackNumber", "mbid", "albumArtist", "duration"},
	})
	return
}

//track.search
func (api trackApi) Search(args P) (result TrackSearch, err error) {
	defer func() { appendCaller(err, "lastfm.Track.Search") }()
	err = callGet("track.search", api.creds, args, &result, P{
		"normal": []string{"artist", "track", "limit", "page"},
	})
	return
}

//track.share
func (api trackApi) Share(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.Share") }()
	err = callPost("track.share", api.creds, args, nil, P{
		"normal": []string{"artist", "track", "public", "message", "recipient"},
	})
	return
}

//track.unban
func (api trackApi) UnBan(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.UnBan") }()
	err = callPost("track.unban", api.creds, args, nil, P{
		"normal": []string{"artist", "track"},
	})
	return
}

//track.unlove
func (api trackApi) UnLove(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.UnLove") }()
	err = callPost("track.unlove", api.creds, args, nil, P{
		"normal": []string{"artist", "track"},
	})
	return
}

//track.updateNowPlaying
func (api trackApi) UpdateNowPlaying(args P) (result TrackUpdateNowPlaying, err error) {
	defer func() { appendCaller(err, "lastfm.Track.UpdateNowPlaying") }()
	err = callPost("track.updatenowplaying", api.creds, args, &result, P{
		"normal": []string{"artist", "track", "album", "trackNumber", "context", "mbid", "duration", "albumArtist"},
	})
	return
}
