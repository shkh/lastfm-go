package lastfm

type eventApi struct {
	creds *credentials
}

//Event.attend
func (api eventApi) Attend(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Event.Attend") }()
	err = callPost("event.attend", api.creds, args, nil, P{
		"normal": []string{"event", "status"},
	})
	return
}

//Event.getAttendees
func (api eventApi) GetAttendees(args P) (result EventGetAttendees, err error) {
	defer func() { appendCaller(err, "lastfm.Event.GetAttendees") }()
	err = callPost("event.attendees", api.creds, args, nil, P{
		"normal": []string{"event", "page", "limit"},
	})
	return
}

//Event.getInfo
func (api eventApi) GetInfo(args P) (result EventGetInfo, err error) {
	defer func() { appendCaller(err, "lastfm.Event.GetInfo") }()
	err = callGet("event.getinfo", api.creds, args, &result, P{
		"normal": []string{"event"},
	})
	return
}

//Event.getShouts
func (api eventApi) GetShouts(args P) (result EventGetShouts, err error) {
	defer func() { appendCaller(err, "lastfm.Event.GetShouts") }()
	err = callGet("event.getshouts", api.creds, args, &result, P{
		"normal": []string{"event", "page", "limit"},
	})
	return
}

//Event.share
func (api eventApi) Share(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Event.Share") }()
	err = callPost("event.share", api.creds, args, nil, P{
		"normal": []string{"event", "public", "message", "recipient"},
	})
	return
}

//Event.shout
func (api eventApi) Shout(args P) (err error) {
	defer func() { appendCaller(err, "lastfm.Event.Shout") }()
	err = callPost("event.shout", api.creds, args, nil, P{
		"normal": []string{"event", "message"},
	})
	return
}
