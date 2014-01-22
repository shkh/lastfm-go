package lastfm

import "encoding/xml"

//group.getHype
type GroupGetHype struct {
	XMLName xml.Name `xml:"weeklyartistchart"`
	From    string   `xml:"from,attr"`
	To      string   `xml:"to,attr"`
	Artist  struct {
		Rank             string `xml:"rank,attr"`
		Name             string `xml:"name"`
		PercentageChange string `xml:"percentagechange"`
		Mbid             string `xml:"mbid"`
		Url              string `xml:"url"`
		Streamable       string `xml:"streamable"`
		Images           []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"artist"`
}

//group.getMembers
type GroupGetMembers struct {
	XMLName    xml.Name `xml:"members"`
	For        string   `xml:"for,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Members    []struct {
		Name     string `xml:"name"`
		RealName string `xml:"realname"`
		Images   []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"user"`
}

//group.getWeeklyAlbumChart
type GroupGetWeeklyAlbumChart struct {
	XMLName xml.Name `xml:"weeklyalbumchart"`
	Group   string   `xml:"group,attr"`
	From    string   `xml:"from,attr"`
	To      string   `xml:"to,attr"`
	Albums  []struct {
		Rank   string `xml:"rank,attr"`
		Artist struct {
			Mbid string `xml:"mbid,attr"`
			Name string `xml:",chardata"`
		} `xml:"artist"`
		Name      string `xml:"name"`
		Mbid      string `xml:"mbid"`
		PlayCount string `xml:"playcount"`
		Url       string `xml:"url"`
	} `xml:"album"`
}

//group.getWeeklyArtistChart
type GroupGetWeeklyArtistChart struct {
	XMLName xml.Name `xml:"weeklyartistchart"`
	Group   string   `xml:"group,attr"`
	From    string   `xml:"from,attr"`
	To      string   `xml:"to,attr"`
	Artists []struct {
		Rank       string `xml:"rank,attr"`
		Name       string `xml:"name"`
		PlayCount  string `xml:"playcount"`
		Mbid       string `xml:"mbid"`
		Url        string `xml:"url"`
		Streamable string `xml:"streamable"`
		Images     []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"artist"`
}

//group.getWeeklyChartList
type GroupGetWeeklyChartlist struct {
	XMLName xml.Name `xml:"weeklychartlist"`
	Group   string   `xml:"group,attr"`
	Charts  []struct {
		From string `xml:"from,attr"`
		To   string `xml:"to,attr"`
	} `xml:"chart"`
}

//group.getWeeklyTrackChart
type GroupGetWeeklyTrackChart struct {
	XMLName xml.Name `xml:"weeklytrackchart"`
	Group   string   `xml:"group,attr"`
	From    string   `xml:"from,attr"`
	To      string   `xml:"to,attr"`
	Tracks  []struct {
		Rank   string `xml:"rank,attr"`
		Artist struct {
			Mbid string `xml:"mbid,attr"`
			Name string `xml:",chardata"`
		} `xml:"artist"`
		Name      string `xml:"name"`
		Mbid      string `xml:"mbid"`
		PlayCount string `xml:"playcount"`
		Images    []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
		Url string `xml:"url"`
	} `xml:"track"`
}
