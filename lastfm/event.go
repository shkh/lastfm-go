package lastfm

type eventApi struct {
	params *apiParams
}

//Event.attend
func (api eventApi) Attend(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Event.Attend") }()
	err = callPost("event.attend", api.params, args, nil, P{
		"plain": []string{"event", "status"},
	})
	return
}

//Event.getAttendees
func (api eventApi) GetAttendees(args map[string]interface{}) (result EventGetAttendees, err error) {
	defer func() { appendCaller(err, "lastfm.Event.GetAttendees") }()
	err = callPost("event.attendees", api.params, args, nil, P{
		"plain": []string{"event", "page", "limit"},
	})
	return
}

//Event.getInfo
func (api eventApi) GetInfo(args map[string]interface{}) (result EventGetInfo, err error) {
	defer func() { appendCaller(err, "lastfm.Event.GetInfo") }()
	err = callGet("event.getinfo", api.params, args, &result, P{
		"plain": []string{"event"},
	})
	return
}

//Event.getShouts
func (api eventApi) GetShouts(args map[string]interface{}) (result EventGetShouts, err error) {
	defer func() { appendCaller(err, "lastfm.Event.GetShouts") }()
	err = callGet("event.getshouts", api.params, args, &result, P{
		"plain": []string{"event", "page", "limit"},
	})
	return
}

//Event.share
func (api eventApi) Share(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Event.Share") }()
	err = callPost("event.share", api.params, args, nil, P{
		"plain": []string{"event", "public", "message", "recipient"},
	})
	return
}

//Event.shout
func (api eventApi) Shout(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Event.Shout") }()
	err = callPost("event.shout", api.params, args, nil, P{
		"plain": []string{"event", "message"},
	})
	return
}
