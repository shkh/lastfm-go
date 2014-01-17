package lastfm

import "encoding/xml"

type RadioGetPlaylist struct {
	XMLName xml.Name `xml:"playlist"`
	Version int      `xml:"version,attr"`
	Title   string   `xml:"title"`
	Creator string   `xml:"creator"`
	Date    string   `xml:"date"`
	Link    struct {
		Rel     string `xml:"rel,attr"`
		Content string `xml:",chardata"`
	}
	Tracks []struct {
		Location  string `xml:"location"`
		Title     string `xml:"title"`
		Id        string `xml:"identifier"`
		Album     string `xml:"album"`
		Creator   string `xml:"creator"`
		Duration  int    `xml:"duration"`
		Image     string `xml:"image"`
		Extension struct {
			Application  string `xml:"application,attr"`
			ArtistPage   string `xml:"artistpage"`
			AlbumPage    string `xml:"albumpage"`
			TrackPage    string `xml:"trackpage"`
			BuyTrackUrl  string `xml:"buyTrackUrl"`
			BuyAlbumUrl  string `xml:"buyAlbumUrl"`
			FreeTrackUrl string `xml:"freeTrackUrl"`
		} `xml:"extension"`
	} `xml:"tracklist"`
}
