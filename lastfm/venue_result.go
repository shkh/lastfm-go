package lastfm

import "encoding/xml"

//venue.getEvents
type VenueGetEvents struct {
	XMLName       xml.Name `xml:"events"`
	Geo           string   `xml:"geo,attr"`
	Venue         string   `xml:"venue,attr"`
	FestivalsOnly string   `xml:"festivalsonly"`
	Events        []struct {
		Geo     string `xml:"geo,attr"`
		Id      string `xml:"id"`
		Title   string `xml:"title"`
		Artists struct {
			Headliner string   `xml:"headliner"`
			Artists   []string `xml:"artist"`
		} `xml:"artists"`
		Venue struct {
			Id       string `xml:"id"`
			Name     string `xml:"name"`
			Location struct {
				City       string `xml:"city"`
				Country    string `xml:"country"`
				Street     string `xml:"street"`
				Postalcode string `xml:"postalcode"`
				Point      struct {
					Lat  string `xml:"lat"`
					Long string `xml:"long"`
				} `xml:"point"`
			} `xml:"location"`
			Url         string `xml:"url"`
			Website     string `xml:"website"`
			PhoneNumber string `xml:"phonenumber"`
			Images      []struct {
				Size string `xml:"size,attr"`
				Url  string `xml:",chardata"`
			} `xml:"image"`
		} `xml:"venue"`
		StartDate   string `xml:"startDate"`
		Description string `xml:"description"`
		Images      []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
		Attendance string `xml:"attendance"`
		Reviews    string `xml:"reviews"`
		Tag        string `xml:"tag"`
		Url        string `xml:"url"`
		Website    string `xml:"website"`
		Tickets    []struct {
			Supplier string `xml:"supplier,attr"`
			Url      string `xml:",chardata"`
		} `xml:"tickets>ticket"`
		Canceled string   `xml:"canceled"`
		Tags     []string `xml:"tags>tag"`
	} `xml:"event"`
}

//venue.getPastEvents
type VenueGetPastEvents struct {
	XMLName       xml.Name `xml:"events"`
	Geo           string   `xml:"geo,attr"`
	Venue         string   `xml:"venue,attr"`
	Url           string   `xml:"url,attr"`
	VenueTimezone string   `xml:"venuetimezone,attr"`
	FestivalsOnly string   `xml:"festivalsonly,attr"`
	Total         int      `xml:"total,attr"`
	Page          int      `xml:"page,attr"`
	PerPage       int      `xml:"perPage,attr"`
	TotalPages    int      `xml:"totalPages,attr"`
	Events        []struct {
		Geo     string `xml:"geo,attr"`
		Id      string `xml:"id"`
		Title   string `xml:"title"`
		Artists struct {
			Headliner string   `xml:"headliner"`
			Artists   []string `xml:"artist"`
		} `xml:"artists"`
		Venue struct {
			Id       string `xml:"id"`
			Name     string `xml:"name"`
			Location struct {
				City       string `xml:"city"`
				Country    string `xml:"country"`
				Street     string `xml:"street"`
				Postalcode string `xml:"postalcode"`
				Point      struct {
					Lat  string `xml:"lat"`
					Long string `xml:"long"`
				} `xml:"point"`
			} `xml:"location"`
			Url         string `xml:"url"`
			Website     string `xml:"website"`
			PhoneNumber string `xml:"phonenumber"`
			Images      []struct {
				Size string `xml:"size,attr"`
				Url  string `xml:",chardata"`
			} `xml:"image"`
		} `xml:"venue"`
		StartDate   string `xml:"startDate"`
		Description string `xml:"description"`
		Images      []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
		Attendance string   `xml:"attendance"`
		Reviews    string   `xml:"reviews"`
		Tag        string   `xml:"tag"`
		Url        string   `xml:"url"`
		Website    string   `xml:"website"`
		Canceled   string   `xml:"canceled"`
		Tags       []string `xml:"tags>tag"`
	} `xml:"event"`
}

//venue.search
type VenueSearch struct {
	XMLName    xml.Name `xml:"results"`
	For        string   `xml:"for,attr"`
	OpenSearch string   `xml:"opensearch,attr"`
	Geo        string   `xml:"geo,attr"`
	Query      struct {
		Role        string `xml:"role,attr"`
		SearchTerms string `xml:"searchTrems,attr"`
		StartPage   int    `xml:"startPage,attr"`
	} `xml:"Query"`
	TotalResults int `xml:"totalResults"`
	StartIndex   int `xml:"startIndex"`
	ItemsPerPage int `xml:"itemsPerPage"`
	Venues       []struct {
		Id       string `xml:"id"`
		Name     string `xml:"name"`
		Location struct {
			City       string `xml:"city"`
			Country    string `xml:"country"`
			Street     string `xml:"street"`
			Postalcode string `xml:"postalcode"`
			Point      struct {
				Lat  string `xml:"lat"`
				Long string `xml:"long"`
			} `xml:"point"`
		} `xml:"location"`
		Url         string `xml:"url"`
		Website     string `xml:"website"`
		Phonenumber string `xml:"phonenumber"`
		Images      []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"venuematches>venue"`
}
