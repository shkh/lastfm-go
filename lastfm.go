package lastfm

const (
	UriApiSecBase  = "https://ws.audioscrobbler.com/2.0/"
	UriApiBase     = "http://ws.audioscrobbler.com/2.0/"
	UriBrowserBase = "http://www.last.fm/api/auth/"
	ResponseFormat = "xml" //Supports only xml so far
)

type P map[string]interface{}

type Api struct {
	creds *credentials
    Album       *albumApi
    Artist      *artistApi
	//Event       *eventApi
	//Geo         *geoApi
	//Group       *groupApi
	Library *libraryApi
	//Tasteometer *tasteometerApi
	Track *trackApi
	User  *userApi
	//Venue       *venueApi
}

type credentials struct {
	apikey   string
	secret   string
	username string
	sk       string
}

func New(key, secret string) (api *Api) {
	c := credentials{key, secret, "", ""}
	api = &Api{
		creds: &c,
        Album:       &albumApi{&c},
        Artist:      &artistApi{&c},
		//Event:       &eventApi{&c},
		//Geo:         &geoApi{&c},
		//Group:       &groupApi{&c},
		Library: &libraryApi{&c},
		Track:   &trackApi{&c},
		//Tasteometer: &tasteometerApi{&c},
		User: &userApi{&c},
		//Venue:       &venueApi{&c},
	}
	return
}

func (api *Api) SetSession(username, sessionkey string) {
	api.creds.username = username
	api.creds.sk = sessionkey
}

func (api Api) GetSessionKey() (sk string) {
	sk = api.creds.sk
	return
}

func (api Api) GetUserName() (username string) {
	username = api.creds.username
	return
}
