package lastfm

type tasteometerApi struct {
	creds *credentials
}

//tasteometer.compare
func (api tasteometerApi) Compare (args P) (result TasteometerCompare, err error) {
    defer func () { appendCaller (err, "lastfm.Tasteometer.Compare") } ()
	err = callGet("tasteometer.compare", api.creds, args, &result, P{
		"normal": []string{"type1", "type2", "value1", "value2", "limit"},
	})
    return
}

//tasteometer.compareGroup (deprecated)
