package lastfm

type venueApi struct {
	creds *credentials
}

//venue.getEvents
func (api venueApi) GetEvents(args P) (result VenueGetEvents, err error) {
	defer func() { appendCaller(err, "lastfm.Venue.GetEvents") }()
	err = callGet("venue.getevents", api.creds, args, &result, P{
		"normal": []string{"venue", "festivalsonly"},
	})
	return
}

//venue.getPastEvents
func (api venueApi) GetPastEvents(args P) (result VenueGetPastEvents, err error) {
	defer func() { appendCaller(err, "lastfm.Venue.GetPastEvents") }()
	err = callGet("venue.getpastevents", api.creds, args, &result, P{
		"normal": []string{"venue", "festivalsonly", "page", "limit"},
	})
	return
}

//venue.search
func (api venueApi) Search(args P) (result VenueSearch, err error) {
	defer func() { appendCaller(err, "lastfm.Venue.Search") }()
	err = callGet("venue.search", api.creds, args, &result, P{
		"normal": []string{"venue", "page", "limit", "country"},
	})
	return
}
