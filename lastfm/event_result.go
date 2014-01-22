package lastfm

import "encoding/xml"

//event.attend (empty)

//event.getAttendees
type EventGetAttendees struct {
	XMLName    xml.Name `xml:"attendees"`
	Geo        string   `xml:"geo,attr"`
	Event      string   `xml:"event"`
	Total      string   `xml:"total,attr"`
	Page       string   `xml:"page,attr"`
	PerPage    string   `xml:"perPage,attr"`
	TotalPages string   `xml:"totalPages,attr"`
	Attendees  []struct {
		Name     string `xml:"name"`
		RealName string `xml:"realname"`
		Url      string `xml:"url"`
		Images   []struct {
			Size string `xml:"size,attr"`
			Url  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"user"`
}

//event.getInfo
type EventGetInfo struct {
	XMLName xml.Name `xml:"event"`
	Geo     string   `xml:"geo,attr"`
	Id      string   `xml:"id"`
	Title   string   `xml:"title"`
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
				Lat  float64 `xml:"lat"`
				Long float64 `xml:"long"`
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
	StartDate   string `xml:"startdate"`
	Description string `xml:"description"`
	Images      []struct {
		Size string `xml:"size,attr"`
		Url  string `xml:",chardata"`
	} `xml:"image"`
	Attendance string `xml:"attendance"`
	Reviews    string `xml:"reviews"`
	Tag        string `xml:"tag"`
	Url        string
	Website    string `xml:"website"`
	Tickets    []struct {
		Supplier string `xml:"supplier,attr"`
		Url      string `xml",chardata"`
	} `xml:"tickets>ticket"`
	Tags []string `xml:"tags>tag"`
}

//event.getShouts
type EventGetShouts struct {
	XMLName    xml.Name `xml:"shouts"`
	Total      string   `xml:"total,attr"`
	Page       string   `xml:"page,attr"`
	PerPage    string   `xml:"perPage,attr"`
	TotalPages string   `xml:"totalPages,attr"`
	Shouts     []struct {
		Body   string `xml:"body"`
		Author string `xml:"author"`
		Date   string `xml:"date"`
	} `xml:"shout"`
}

//event.share (emptry)

//event.shout (empty)

//evenc.getInfo (empty)
