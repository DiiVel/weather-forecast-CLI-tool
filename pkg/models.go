package pkg

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}

type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Condition struct {
	Text string `json:"text"`
}

type Current struct {
	TempC     float64 `json:"temp_c"`
	Condition Condition
}

type Hour []struct {
	TimeEpoch    int64   `json:"time_epoch"`
	TempC        float64 `json:"temp_c"`
	Condition    Condition
	ChanceOfRain float64 `json:"chance_of_rain"`
}

type ForecastDay []struct {
	Hour Hour `json:"hour"`
}

type Forecast struct {
	ForecastDay ForecastDay `json:"forecastday"`
}
