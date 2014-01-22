package lastfm

type tasteometerApi struct {
	params *apiParams
}

//tasteometer.compare
func (api tasteometerApi) Compare(args map[string]interface{}) (result TasteometerCompare, err error) {
	defer func() { appendCaller(err, "lastfm.Tasteometer.Compare") }()
	err = callGet("tasteometer.compare", api.params, args, &result, P{
		"plain": []string{"type1", "type2", "value1", "value2", "limit"},
	})
	return
}

//tasteometer.compareGroup (deprecated)
