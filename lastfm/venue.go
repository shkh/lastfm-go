package lastfm

type venueApi struct {
	params *apiParams
}

//venue.getEvents
func (api venueApi) GetEvents(args map[string]interface{}) (result VenueGetEvents, err error) {
	defer func() { appendCaller(err, "lastfm.Venue.GetEvents") }()
	err = callGet("venue.getevents", api.params, args, &result, P{
		"plain": []string{"venue", "festivalsonly"},
	})
	return
}

//venue.getPastEvents
func (api venueApi) GetPastEvents(args map[string]interface{}) (result VenueGetPastEvents, err error) {
	defer func() { appendCaller(err, "lastfm.Venue.GetPastEvents") }()
	err = callGet("venue.getpastevents", api.params, args, &result, P{
		"plain": []string{"venue", "festivalsonly", "page", "limit"},
	})
	return
}

//venue.search
func (api venueApi) Search(args map[string]interface{}) (result VenueSearch, err error) {
	defer func() { appendCaller(err, "lastfm.Venue.Search") }()
	err = callGet("venue.search", api.params, args, &result, P{
		"plain": []string{"venue", "page", "limit", "country"},
	})
	return
}
