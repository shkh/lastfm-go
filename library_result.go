package lastfm

import "encoding/xml"

//library.addAlbum
//library.addArtist
//library.addTrack

//library.getAlbums
type LibraryGetAlbums struct {
    XMLName xml.Name `xml:"albums"`
    User string `xml:"user,attr"`
	Total      string   `xml:"total,attr"`
	Page       string   `xml:"page,attr"`
	PerPage    string   `xml:"perPage,attr"`
	TotalPages string   `xml:"totalPages"`
    Albums []struct {
        Name string `xml:"name"`
        PlayCount string `xml:"playcount"`
        TagCount string `xml:"tagcount"`
        Mbid string `xml:"mbid"`
        Url string `xml:"url"`
        Artist struct {
            Name string `xml:"name"`
            Mbid string `xml:"mbid"`
            Url string `xml:"url"`
        } `xml:"artist"`
		Images []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`

    } `xml:"album"`
}

//library.getArtists
type LibraryGetArtists struct {
    XMLName xml.Name `xml:"artists"`
    User string `xml:"user,attr"`
	Total      string   `xml:"total,attr"`
	Page       string   `xml:"page,attr"`
	PerPage    string   `xml:"perPage,attr"`
	TotalPages string   `xml:"totalPages"`
    Artists []struct {
        Name string `xml:"name"`
        PlayCount string `xml:"playcount"`
        TagCount string `xml:"tagcount"`
        Mbid string `xml:"mbid"`
        Url string `xml:"url"`
        Streamable string `xml:"streamable"`
		Images []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
    } `xml:"artist"`
}

//library.getTracks
type LibraryGetTracks struct {
    XMLName xml.Name `xml:"tracks"`
    User string `xml:"user,attr"`
	Total      string   `xml:"total,attr"`
	Page       string   `xml:"page,attr"`
	PerPage    string   `xml:"perPage,attr"`
	TotalPages string   `xml:"totalPages"`
    Tracks []struct {
        Name string `xml:"name"`
        Duration string `xml:"duration"`
        PlayCount string `xml:"playcount"`
        TagCount string `xml:"tagcount"`
        Mbid string `xml:"mbid"`
        Url string `xml:"url"`
		Streamable struct {
			FullTrack  string `xml:"fulltrack,attr"`
			Streamable string `xml:"streamable"`
		} `xml:"streamable"`
        Artist struct {
            Name string `xml:"name"`
            Mbid string `xml:"mbid"`
            Url string `xml:"url"`
        } `xml:"artist"`
        Album struct {
            Name string `xml:"name"`
            Position string `xml:"position"`
        } `xml:"album"`
		Images []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
    } `xml:"track"`
}

//library.removeAlbum
//library.removeArtist
//library.removeScrobble
//library.removeTrack
