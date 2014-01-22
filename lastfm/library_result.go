package lastfm

import "encoding/xml"

//library.addAlbum (emptry)

//library.addArtist (empty)

//library.addTrack (empty)

//library.getAlbums
type LibraryGetAlbums struct {
	XMLName    xml.Name `xml:"albums"`
	User       string   `xml:"user,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Albums     []struct {
		Name      string `xml:"name"`
		PlayCount string `xml:"playcount"`
		TagCount  string `xml:"tagcount"`
		Mbid      string `xml:"mbid"`
		Url       string `xml:"url"`
		Artist    struct {
			Name string `xml:"name"`
			Mbid string `xml:"mbid"`
			Url  string `xml:"url"`
		} `xml:"artist"`
		Images []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"album"`
}

//library.getArtists
type LibraryGetArtists struct {
	XMLName    xml.Name `xml:"artists"`
	User       string   `xml:"user,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Artists    []struct {
		Name       string `xml:"name"`
		PlayCount  string `xml:"playcount"`
		TagCount   string `xml:"tagcount"`
		Mbid       string `xml:"mbid"`
		Url        string `xml:"url"`
		Streamable string `xml:"streamable"`
		Images     []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"artist"`
}

//library.getTracks
type LibraryGetTracks struct {
	XMLName    xml.Name `xml:"tracks"`
	User       string   `xml:"user,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Tracks     []struct {
		Name       string `xml:"name"`
		Duration   string `xml:"duration"`
		PlayCount  string `xml:"playcount"`
		TagCount   string `xml:"tagcount"`
		Mbid       string `xml:"mbid"`
		Url        string `xml:"url"`
		Streamable struct {
			FullTrack  string `xml:"fulltrack,attr"`
			Streamable string `xml:"streamable"`
		} `xml:"streamable"`
		Artist struct {
			Name string `xml:"name"`
			Mbid string `xml:"mbid"`
			Url  string `xml:"url"`
		} `xml:"artist"`
		Album struct {
			Name     string `xml:"name"`
			Position string `xml:"position"`
		} `xml:"album"`
		Images []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"track"`
}

//library.removeAlbum (empty)

//library.removeArtist (empty)

//library.removeScrobble (empty)

//library.removeTrack (empty)
