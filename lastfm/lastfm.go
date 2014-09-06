package lastfm

const (
	UriApiSecBase  = "https://ws.audioscrobbler.com/2.0/"
	UriApiBase     = "http://ws.audioscrobbler.com/2.0/"
	UriBrowserBase = "https://www.last.fm/api/auth/"
)

type P map[string]interface{}

type Api struct {
	params      *apiParams
	Album       *albumApi
	Artist      *artistApi
	Chart       *chartApi
	Event       *eventApi
	Geo         *geoApi
	Group       *groupApi
	Library     *libraryApi
	Playlist    *playlistApi
	Radio       *radioApi
	Tag         *tagApi
	Tasteometer *tasteometerApi
	Track       *trackApi
	User        *userApi
	Venue       *venueApi
}

type apiParams struct {
	apikey    string
	secret    string
	sk        string
	useragent string
}

func New(key, secret string) (api *Api) {
	params := apiParams{key, secret, "", ""}
	api = &Api{
		params:      &params,
		Album:       &albumApi{&params},
		Artist:      &artistApi{&params},
		Chart:       &chartApi{&params},
		Event:       &eventApi{&params},
		Geo:         &geoApi{&params},
		Group:       &groupApi{&params},
		Library:     &libraryApi{&params},
		Playlist:    &playlistApi{&params},
		Radio:       &radioApi{&params},
		Tag:         &tagApi{&params},
		Tasteometer: &tasteometerApi{&params},
		Track:       &trackApi{&params},
		User:        &userApi{&params},
		Venue:       &venueApi{&params},
	}
	return
}

func (api *Api) SetSession(sessionkey string) {
	api.params.sk = sessionkey
}

func (api Api) GetSessionKey() (sk string) {
	sk = api.params.sk
	return
}

func (api *Api) SetUserAgent(useragent string) {
	api.params.useragent = useragent
}
