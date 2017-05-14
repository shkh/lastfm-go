package lastfm

import "encoding/xml"

type GeoGetEventsEvent struct {
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
}

//geo.getEvents
type GeoGetEvents struct {
	XMLName       xml.Name            `xml:"events"`
	Geo           string              `xml:"geo,attr"`
	Location      string              `xml:"location,attr"`
	Tag           string              `xml:"tag,attr"`
	FestivalsOnly string              `xml:"festivalsonly,attr"`
	Total         int                 `xml:"total,attr"`
	Page          int                 `xml:"page,attr"`
	PerPage       int                 `xml:"perPage,attr"`
	TotalPages    int                 `xml:"totalPages,attr"`
	Events        []GeoGetEventsEvent `xml:"event"`
}

//geo.getMetroArtistChart
type GeoGetMetroArtistChart struct {
	XMLName    xml.Name `xml:"topartists"`
	Metro      string   `xml:"metro,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Artists    []struct {
		Rank       string `xml:"rank,attr"`
		Name       string `Xml:"name"`
		Listeners  string `xml:"listeners"`
		Mbid       string `xml:"mbid"`
		Url        string `xml:"url"`
		Streamable string `xml:"streamable"`
		Images     []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"artist"`
}

//geo.getMetroHypeArtistChart
type GeoGetMetroHypeArtistChart struct {
	XMLName    xml.Name `xml:"topartists"`
	Metro      string   `xml:"metro,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Artists    []struct {
		Rank       string `xml:"rank,attr"`
		Name       string `Xml:"name"`
		Mbid       string `xml:"mbid"`
		Url        string `xml:"url"`
		Streamable string `xml:"streamable"`
		Images     []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"artist"`
}

//geo.getMetroHypeTrackChart
type GeoGetMetroHypeTrackChart struct {
	XMLName    xml.Name `xml:"toptracks"`
	Metro      string   `xml:"metro,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Tracks     []struct {
		Rank       string `xml:"rank,attr"`
		Name       string `xml:"name"`
		Duration   string `xml:"duration"`
		Mbid       string `xml:"mbid"`
		Url        string `xml:"url"`
		Streamable struct {
			FullTrack  string `xml:"fulltrack,attr"`
			Streamable string `xml:",chardata"`
		} `xml:"streamable"`
		Artist struct {
			Name string `xml:"artist"`
			Mbid string `xml:"mbid"`
			Url  string `xml:"url"`
		} `xml:"artist"`
		Images []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"track"`
}

//geo.getMetroTrackChart
type GeoGetMetroTrackChart struct {
	XMLName    xml.Name `xml:"toptracks"`
	Metro      string   `xml:"metro,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Tracks     []struct {
		Rank       string `xml:"rank,attr"`
		Name       string `xml:"name"`
		Duration   string `xml:"duration"`
		Listeners  string `xml:"listeners"`
		Mbid       string `xml:"mbid"`
		Url        string `xml:"url"`
		Streamable struct {
			FullTrack  string `xml:"fulltrack,attr"`
			Streamable string `xml:",chardata"`
		} `xml:"streamable"`
		Artist struct {
			Name string `xml:"artist"`
			Mbid string `xml:"mbid"`
			Url  string `xml:"url"`
		} `xml:"artist"`
		Images []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"track"`
}

//geo.getMetroUniqueArtistChart
type GeoGetMetroUniqueArtistChart struct {
	XMLName    xml.Name `xml:"topartists"`
	Metro      string   `xml:"metro,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Artists    []struct {
		Rank       string `xml:"rank,attr"`
		Name       string `Xml:"name"`
		Mbid       string `xml:"mbid"`
		Url        string `xml:"url"`
		Streamable string `xml:"streamable"`
		Images     []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"artist"`
}

//geo.getMetroUniqueTrackChart
type GeoGetMetroUniqueTrackChart struct {
	XMLName    xml.Name `xml:"toptracks"`
	Metro      string   `xml:"metro,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Tracks     []struct {
		Rank       string `xml:"rank,attr"`
		Name       string `xml:"name"`
		Duration   string `xml:"duration"`
		Mbid       string `xml:"mbid"`
		Url        string `xml:"url"`
		Streamable struct {
			FullTrack  string `xml:"fulltrack,attr"`
			Streamable string `xml:",chardata"`
		} `xml:"streamable"`
		Artist struct {
			Name string `xml:"artist"`
			Mbid string `xml:"mbid"`
			Url  string `xml:"url"`
		} `xml:"artist"`
		Images []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"track"`
}

//geo.getMetroWeeklyChartlist
type GeoGetMetroWeeklyChartlist struct {
	XMLName xml.Name `xml:"weeklychartlist"`
	Charts  []struct {
		From string `xml:"from,attr"`
		To   string `xml:"to,attr"`
	} `xml:"chart"`
}

//geo.getMetros
type GeoGetMetros struct {
	XMLName xml.Name `xml:"metros"`
	Metros  []struct {
		Name    string `xml:"name"`
		Country string `xml:"country"`
	} `xml:"metro"`
}

//geo.getTopArtists
type GeoGetTopArtists struct {
	XMLName    xml.Name `xml:"topartists"`
	Country    string   `xml:"country,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Artists    []struct {
		Rank       string `xml:"rank,attr"`
		Name       string `Xml:"name"`
		Listeners  string `xml:"listeners"`
		Mbid       string `xml:"mbid"`
		Url        string `xml:"url"`
		Streamable string `xml:"streamable"`
		Images     []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"artist"`
}

//geo.getTopTracks
type GeoGetTopTracks struct {
	XMLName    xml.Name `xml:"toptracks"`
	Country    string   `xml:"country,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Tracks     []struct {
		Rank       string `xml:"rank,attr"`
		Name       string `xml:"name"`
		Duration   string `xml:"duration"`
		Listeners  string `xml:"listeners"`
		Mbid       string `xml:"mbid"`
		Url        string `xml:"url"`
		Streamable struct {
			FullTrack  string `xml:"fulltrack,attr"`
			Streamable string `xml:",chardata"`
		} `xml:"streamable"`
		Artist struct {
			Name string `xml:"artist"`
			Mbid string `xml:"mbid"`
			Url  string `xml:"url"`
		} `xml:"artist"`
		Images []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"track"`
}
