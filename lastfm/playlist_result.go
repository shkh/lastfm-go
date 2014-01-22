package lastfm

import "encoding/xml"

//playlist.addTrack (empty)

//playlist.create
type PlaylistCreate struct {
	XMLName  xml.Name `xml:"playlists"`
	User     string   `xml:"user,attr"`
	Playlist struct {
		Id          string `xml:"id"`
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Date        string `xml:"date"`
		Size        string `xml:"size"`
		Duration    string `xml:"duration"`
		Streamable  string `xml:"streamable"`
		Creator     string `xml:"creator"`
		Url         string `xml:"url"`
		Images      []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"playlist"`
}
