package lastfm

type trackApi struct {
	params *apiParams
}

//track.addTags
func (api trackApi) AddTags(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.AddTags") }()
	err = callPost("track.addtags", api.params, args, nil, P{
		"plain": []string{"artist", "track", "tags"},
	})
	return
}

//track.ban
func (api trackApi) Ban(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.Ban") }()
	err = callPost("track.ban", api.params, args, nil, P{
		"plain": []string{"artist", "track"},
	})
	return
}

//track.getBuylinks
func (api trackApi) GetBuylinks(args map[string]interface{}) (result TrackGetBuylinks, err error) {
	defer func() { appendCaller(err, "lastfm.Track.GetBuylinks") }()
	err = callGet("track.getbuylinks", api.params, args, &result, P{
		"plain": []string{"artist", "track", "mbid", "autocorrect", "country"},
	})
	return
}

//track.getCorrection
func (api trackApi) GetCorrection(args map[string]interface{}) (result TrackGetCorrection, err error) {
	defer func() { appendCaller(err, "lastfm.Track.Correction") }()
	err = callGet("track.getcorrection", api.params, args, &result, P{
		"plain": []string{"artist", "track"},
	})
	return
}

//track.getFingerprintMetadata
func (api trackApi) GetFingerprintMetadata(args map[string]interface{}) (result TrackGetFingerprintMetadata, err error) {
	defer func() { appendCaller(err, "lastfm.Track.Correction") }()
	err = callGet("track.getfingerprintmetadata", api.params, args, &result, P{
		"plain": []string{"fingerprintid"},
	})
	return
}

//track.getInfo
func (api trackApi) GetInfo(args map[string]interface{}) (result TrackGetInfo, err error) {
	defer func() { appendCaller(err, "lastfm.Track.GetInfo") }()
	err = callGet("track.getinfo", api.params, args, &result, P{
		"plain": []string{"artist", "track", "mbid", "username", "autocorrect"},
	})
	return
}

//track.getShouts
func (api trackApi) GetShouts(args map[string]interface{}) (result TrackGetShouts, err error) {
	defer func() { appendCaller(err, "lastfm.Track.GetShouts") }()
	err = callGet("track.getshouts", api.params, args, &result, P{
		"plain": []string{"artist", "track", "mbid", "limit", "page", "autocorrect"},
	})
	return
}

//track.getSimilar
func (api trackApi) GetSimilar(args map[string]interface{}) (result TrackGetSimilar, err error) {
	defer func() { appendCaller(err, "lastfm.Track.GetSimilar") }()
	err = callGet("track.getsimilar", api.params, args, &result, P{
		"plain": []string{"artist", "track", "mbid", "limit", "autocorrect"},
	})
	return
}

//track.getTags
func (api trackApi) GetTags(args map[string]interface{}) (result TrackGetTags, err error) {
	defer func() { appendCaller(err, "lastfm.Track.GetTags") }()
	if _, ok := args["users"]; !ok && api.params.sk != "" {
		err = callPost("track.gettags", api.params, args, &result, P{
			"plain": []string{"artist", "track", "mbid", "autocorrect"},
		})
	} else {
		err = callGet("track.gettags", api.params, args, &result, P{
			"plain": []string{"artist", "track", "mbid", "user", "autocorrect"},
		})
	}
	return
}

//track.getTopFans
func (api trackApi) GetTopFans(args map[string]interface{}) (result TrackGetTopFans, err error) {
	defer func() { appendCaller(err, "lastfm.Track.GetTopFans") }()
	err = callGet("track.gettopfans", api.params, args, &result, P{
		"plain": []string{"artist", "track", "mbid", "autocorrect"},
	})
	return
}

//track.getTopTags
func (api trackApi) GetTopTags(args map[string]interface{}) (result TrackGetTopTags, err error) {
	defer func() { appendCaller(err, "lastfm.Track.GetTopTags") }()
	err = callGet("track.gettoptags", api.params, args, &result, P{
		"plain": []string{"artist", "track", "mbid", "autocorrect"},
	})
	return
}

//track.love
func (api trackApi) Love(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.Love") }()
	err = callPost("track.love", api.params, args, nil, P{
		"plain": []string{"artist", "track"},
	})
	return
}

//track.removeTag
func (api trackApi) RemoveTag(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.RemoveTag") }()
	err = callPost("track.removetag", api.params, args, nil, P{
		"plain": []string{"artist", "track", "tag"},
	})
	return
}

//track.scrobble
func (api trackApi) Scrobble(args map[string]interface{}) (result TrackScrobble, err error) {
	defer func() { appendCaller(err, "lastfm.Track.Scrobble") }()
	err = callPost("track.scrobble", api.params, args, &result, P{
		"indexing": []string{"artist", "track", "timestamp", "album", "context", "streamId", "chosenByUser", "trackNumber", "mbid", "albumArtist", "duration"},
	})
	return
}

//track.search
func (api trackApi) Search(args map[string]interface{}) (result TrackSearch, err error) {
	defer func() { appendCaller(err, "lastfm.Track.Search") }()
	err = callGet("track.search", api.params, args, &result, P{
		"plain": []string{"artist", "track", "limit", "page"},
	})
	return
}

//track.share
func (api trackApi) Share(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.Share") }()
	err = callPost("track.share", api.params, args, nil, P{
		"plain": []string{"artist", "track", "public", "message", "recipient"},
	})
	return
}

//track.unban
func (api trackApi) UnBan(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.UnBan") }()
	err = callPost("track.unban", api.params, args, nil, P{
		"plain": []string{"artist", "track"},
	})
	return
}

//track.unlove
func (api trackApi) UnLove(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.UnLove") }()
	err = callPost("track.unlove", api.params, args, nil, P{
		"plain": []string{"artist", "track"},
	})
	return
}

//track.updateNowPlaying
func (api trackApi) UpdateNowPlaying(args map[string]interface{}) (result TrackUpdateNowPlaying, err error) {
	defer func() { appendCaller(err, "lastfm.Track.UpdateNowPlaying") }()
	err = callPost("track.updatenowplaying", api.params, args, &result, P{
		"plain": []string{"artist", "track", "album", "trackNumber", "context", "mbid", "duration", "albumArtist"},
	})
	return
}
