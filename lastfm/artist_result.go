package lastfm

import "encoding/xml"

//////////////
// artist.x //
//////////////

//artist.addTags (empty)

//artist.getCorrection
type ArtistGetCorrection struct {
	XMLName    xml.Name `xml:"corrections"`
	Correction struct {
		Index  string `xml:"index,attr"`
		Artist struct {
			Name string `xml:"name"`
			Mbid string `xml:"mbid"`
			Url  string `xml:"url"`
		} `xml:"artist"`
	} `xml:"correction"`
}

//artist.getEvents
type ArtistGetEvents struct {
	XMLName       xml.Name `xml:"events"`
	Geo           string   `xml:"geo,attr"`
	Artist        string   `xml:"artist,attr"`
	FestivalsOnly string   `xml:"festivalsonly,attr"`
	Total         int      `xml:"total,attr"`
	Page          int      `xml:"page,attr"`
	PerPage       int      `xml:"perPage,attr"`
	TotalPages    int      `xml:"totalPages,attr"`
	Events        []struct {
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

//artist.getInfo
type ArtistGetInfo struct {
	Name   string `xml:"name"`
	Mbid   string `xml:"mbid"`
	Url    string `xml:"url"`
	Images []struct {
		Size string `xml:"size,attr"`
		Url  string `xml:",chardata"`
	} `xml:"image"`
	Streamable string `xml:"streamable"`
	Stats      struct {
		Listeners string `xml:"listeners"`
		Plays     string `xml:"plays"`
	} `xml:"stats"`
	//Similar struct {
	//Artists []struct {
	//Name   string `xml:"name"`
	//Url    string `xml:"url"`
	//Images []struct {
	//Size string `xml:"size,attr"`
	//Url  string `xml:",chardata"`
	//} `xml:"image"`
	//} `xml:"artist"`
	//} `xml:"similar"`
	Similars []struct {
		Name   string `xml:"name"`
		Url    string `xml:"url"`
		Images []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"similar>artist"`
	Tags []struct {
		Name string `xml:"name"`
		Url  string `xml:"url"`
	} `xml:"tags>tag"`
	Bio struct {
		Links []struct {
			Rel string `xml:"rel,attr"`
			Url string `xml:"href,attr"`
		} `xml:"links>link"`
		Published  string `xml:"published"`
		Summary    string `xml:"summary"`
		Content    string `xml:"content"`
		YearFormed string `xml:"yearformed"`
	} `xml:"bio"`
}

//artist.getPastEvents
type ArtistGetPastEvents struct {
	XMLName       xml.Name `xml:"events"`
	Artist        string   `xml:"artist,attr"`
	Url           string   `xml:"url,attr"`
	Geo           string   `xml:"geo,attr"`
	FestivalsOnly string   `xml:"festivalsonly,attr"`
	Total         int      `xml:"total,attr"`
	Page          int      `xml:"page,attr"`
	PerPage       int      `xml:"perPage,attr"`
	TotalPages    int      `xml:"totalPages,attr"`
	Events        []struct {
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

//artist.getPodcast
type ArtistGetPodcast struct {
	XMLName xml.Name `xml:"rss"`
	Itunes  string   `xml:"itunes,attr"`
	Version string   `xml:"version,attr"`
	Channel struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Items       []struct {
			Title string `xml:"title"`
			Guid  struct {
				IdPermalink string `xml:"isPermalink,attr"`
				Guid        string `xml:",chardata"`
			} `xml:"guid"`
			Author    string `xml:"author"`
			Enclosure struct {
				Url    string `xml:"url,attr"`
				Length string `xml:"length,attr"`
				Type   string `xml:"type,attr"`
			}
			Description string `xml:"description"`
			Link        string `xml:"link"`
		} `xml:"item"`
	} `xml:"channel"`
}

//artist.getShouts
type ArtistGetShouts struct {
	XMLName    xml.Name `xml:"shouts"`
	Artist     string   `xml:"artist,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Shouts     []struct {
		Body   string `xml:"body"`
		Author string `xml:"author"`
		Date   string `xml:"date"`
	} `xml:"shout"`
}

//artist.getSimilar
type ArtistGetSimilar struct {
	XMLName  xml.Name `xml:"similarartists"`
	Artist   string   `xml:"artist,attr"`
	Similars []struct {
		Name   string `xml:"name"`
		Mbid   string `xml:"mbid"`
		Match  string `xml:"match"`
		Url    string `xml:"url"`
		Images []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
		Streamable string `xml:"streamable"`
	} `xml:"artist"`
}

//artist.getTags
type ArtistGetTags struct {
	XMLName xml.Name `xml:"tags"`
	Artist  string   `xml:"artist,attr"`
	Tags    []struct {
		Name string `xml:"name"`
		Url  string `xml:"url"`
	} `xml:"tag"`
}

//artist.getTopAlbums
type ArtistGetTopAlbums struct {
	XMLName    xml.Name `xml:"topalbums"`
	Artist     string   `xml:"artist,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Albums     []struct {
		Rank   string `xml:"rank,attr"`
		Name   string `xml:"name"`
		Mbid   string `xml:"mbid"`
		Artist struct {
			Name string `xml:"name"`
			Mbid string `xml:"mbid"`
			Url  string `xml:"url"`
		} `xml:"artist"`
		PlayCount string `xml:"playcount"`
		Url       string `xml:"url"`
		Images    []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"album"`
}

//artist.getTopFans
type ArtistGetTopFans struct {
	XMLName xml.Name `xml:"topfans"`
	Artist  string   `xml:"artist,attr"`
	Users   []struct {
		Name     string `xml:"name"`
		RealName string `xml:"realname"`
		Url      string `xml:"url"`
		Images   []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
		Weight string `xml:"weight"`
	} `xml:"user"`
}

//artist.getTopTags
type ArtistGetTopTags struct {
	XMLName xml.Name `xml:"toptags"`
	Artist  string   `xml:"artist,attr"`
	Tags    []struct {
		Name  string `xml:"name"`
		Count string `xml:"count"`
		Url   string `xml:"url"`
	} `xml:"tag"`
}

//artist.getTopTracks
type ArtistGetTopTracks struct {
	XMLName    xml.Name `xml:"toptracks"`
	Artist     string   `xml:"artist,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Tracks     []struct {
		Rank      string `xml:"rank,attr"`
		Name      string `xml:"name"`
		Duration  string `xml:"duration"`
		PlayCount string `xml:"playcount"`
	} `xml:"track"`
}

//artist.removeTag (empty)

//artist.search
type ArtistSearch struct {
	XMLName    xml.Name `xml:"results"`
	OpenSearch string   `xml:"opensearch,attr"`
	For        string   `xml:"for,attr"`
	Query      struct {
		Role        string `xml:"role,attr"`
		SearchTerms string `xml:"searchTrems,attr"`
		StartPage   int    `xml:"startPage,attr"`
	} `xml:"Query"`
	TotalResults  int `xml:"totalResults"`
	StartIndex    int `xml:"startIndex"`
	ItemsPerPage  int `xml:"itemsPerPage"`
	ArtistMatches []struct {
		Name   string `xml:"name"`
		Mbid   string `xml:"mbid"`
		Url    string `xml:"url"`
		Images []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
		Listeners string `xml:"listeners"`
	} `xml:"artistmatches>artist"`
}

//artist.share (empty)

//artist.shout (empty)
