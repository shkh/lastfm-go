package lastfm

import "encoding/xml"

//tasteometer.compare
type TasteometerCompare struct {
	XMLName xml.Name `xml:"comparison"`
	Result  struct {
		Score   string `xml:"score"`
		Artists struct {
			Matches string `xml:"matches,attr"`
			Artists []struct {
				Name   string `xml:"name"`
				Url    string `xml:"url"`
				Images []struct {
					Size string `xml:"size,attr"`
					Url  string `xml:",chardata"`
				} `xml:"image"`
			} `xml:"artist"`
		} `xml:"artists"`
	} `xml:"result"`
}
