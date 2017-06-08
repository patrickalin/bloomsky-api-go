// Package bloomskyStructure calls rest API Bloomsky, puts it in the structure and gives somes functions
package bloomskyStructure

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	rest "github.com/patrickalin/myrest-go"
)

// BloomskyStructure represents the structure of the JSON return by the API
type BloomskyStructure struct {
	UTC              float64                `json:"UTC"`
	CityName         string                 `json:"CityName"`
	Storm            BloomskyStormStructure `json:"Storm"`
	Searchable       bool                   `json:"Searchable"`
	DeviceName       string                 `json:"DeviceName"`
	RegisterTime     float64                `json:"RegisterTime"`
	DST              float64                `json:"DST"`
	BoundedPoint     string                 `json:"BoundedPoint"`
	LON              float64                `json:"LON"`
	Point            interface{}            `json:"Point"`
	VideoList        []string               `json:"VideoList"`
	VideoListC       []string               `json:"VideoList_C"`
	DeviceID         string                 `json:"DeviceID"`
	NumOfFollowers   float64                `json:"NumOfFollowers"`
	LAT              float64                `json:"LAT"`
	ALT              float64                `json:"ALT"`
	Data             BloomskyDataStructure  `json:"Data"`
	FullAddress      string                 `json:"FullAddress"`
	StreetName       string                 `json:"StreetName"`
	PreviewImageList []string               `json:"PreviewImageList"`
}

// BloomskyStormStructure represents the structure STORM of the JSON return by the API
type BloomskyStormStructure struct {
	UVIndex               string  `json:"UVIndex"`
	WindDirection         string  `json:"WindDirection"`
	WindGust              float64 `json:"WindGust"`
	SustainedWindSpeed    float64 `json:"SustainedWindSpeed"`
	WindGustms            float64
	SustainedWindSpeedms  float64
	WindGustkmh           float64
	SustainedWindSpeedkmh float64
	Rain                  float64
	RainDaily             float64 `json:"RainDaily"`
	RainDailymm           float64
	RainRate              float64 `json:"RainRate"`
	RainRatemm            float64
	Rainin                float64 `json:"24hRain"`
}

// BloomskyDataStructure represents the structure SKY of the JSON return by the API
type BloomskyDataStructure struct {
	Luminance    float64 `json:"Luminance"`
	TemperatureF float64 `json:"Temperature"`
	TemperatureC float64
	ImageURL     string  `json:"ImageURL"`
	TS           float64 `json:"TS"`
	Rain         bool    `json:"Rain"`
	Humidity     float64 `json:"Humidity"`
	Pressure     float64 `json:"Pressure"`
	Pressurehpa  float64
	DeviceType   string  `json:"DeviceType"`
	Voltage      float64 `json:"Voltage"`
	Night        bool    `json:"Night"`
	UVIndex      float64 `json:"UVIndex"`
	ImageTS      float64 `json:"ImageTS"`
}

// bloomskyStructure is the interface bloomskyStructure
type bloomskyStructure interface {
	GetDeviceID() string
	GetSoftwareVersion() string
	GetAmbientTemperatureC() float64
	GetTargetTemperatureC() float64
	GetAmbientTemperatureF() float64
	GetTargetTemperatureF() float64
	GetHumidity() float64
	GetAway() string
	GetCity() string
	ShowPrettyAll() int
}

var debug = false

// ShowPrettyAll prints to the console the JSON
func (bloomskyInfo BloomskyStructure) ShowPrettyAll() {
	out, err := json.Marshal(bloomskyInfo)
	if err != nil {
		log.Fatal(fmt.Errorf("Error with parsing Json"))
	}
	if debug {
		fmt.Printf("Decode:> \n %s \n\n", out)
	}
}

//GetTimeStamp returns the timestamp give by Bloomsky
func (bloomskyInfo BloomskyStructure) GetTimeStamp() time.Time {
	tm := time.Unix(int64(bloomskyInfo.Data.TS), 0)
	return tm
}

//GetCity returns the city name
func (bloomskyInfo BloomskyStructure) GetCity() string {
	return bloomskyInfo.CityName
}

//GetDeviceID returns the Device Id
func (bloomskyInfo BloomskyStructure) GetDeviceID() string {
	return bloomskyInfo.DeviceID
}

//GetNumOfFollowers returns the number of followers
func (bloomskyInfo BloomskyStructure) GetNumOfFollowers() int {
	return int(bloomskyInfo.NumOfFollowers)
}

//GetIndexUV returns the UV index from 1 to 11
func (bloomskyInfo BloomskyStructure) GetIndexUV() string {
	return bloomskyInfo.Storm.UVIndex
}

//IsNight returns true if it's the night
func (bloomskyInfo BloomskyStructure) IsNight() bool {
	return bloomskyInfo.Data.Night
}

//GetTemperatureFahrenheit returns temperature in Fahrenheit
func (bloomskyInfo BloomskyStructure) GetTemperatureFahrenheit() float64 {
	return bloomskyInfo.Data.TemperatureF
}

//GetTemperatureCelsius returns temperature in Celsius
func (bloomskyInfo BloomskyStructure) GetTemperatureCelsius() float64 {
	return bloomskyInfo.Data.TemperatureC
}

//GetHumidity returns hulidity %
func (bloomskyInfo BloomskyStructure) GetHumidity() float64 {
	return bloomskyInfo.Data.Humidity
}

//GetPressureHPa returns pressure in HPa
func (bloomskyInfo BloomskyStructure) GetPressureHPa() float64 {
	return bloomskyInfo.Data.Pressurehpa
}

//GetPressureInHg returns pressure in InHg
func (bloomskyInfo BloomskyStructure) GetPressureInHg() float64 {
	return bloomskyInfo.Data.Pressure
}

//GetWindDirection returns wind direction (N,S,W,E, ...)
func (bloomskyInfo BloomskyStructure) GetWindDirection() string {
	return bloomskyInfo.Storm.WindDirection
}

//GetWindGustMph returns Wind in Mph
func (bloomskyInfo BloomskyStructure) GetWindGustMph() float64 {
	return bloomskyInfo.Storm.WindGust
}

//GetWindGustMs returns Wind in Ms
func (bloomskyInfo BloomskyStructure) GetWindGustMs() float64 {
	return (bloomskyInfo.Storm.WindGust * 1.61)
}

//GetSustainedWindSpeedMph returns Sustained Wind Speed in Mph
func (bloomskyInfo BloomskyStructure) GetSustainedWindSpeedMph() float64 {
	return bloomskyInfo.Storm.SustainedWindSpeed
}

//GetSustainedWindSpeedMs returns Sustained Wind Speed in Ms
func (bloomskyInfo BloomskyStructure) GetSustainedWindSpeedMs() float64 {
	return (bloomskyInfo.Storm.SustainedWindSpeed * 1.61)
}

//IsRain returns true if it's rain
func (bloomskyInfo BloomskyStructure) IsRain() bool {
	return bloomskyInfo.Data.Rain
}

//GetRainDailyIn returns rain daily in In
func (bloomskyInfo BloomskyStructure) GetRainDailyIn() float64 {
	return bloomskyInfo.Storm.RainDaily
}

//GetRainIn returns total rain in In
func (bloomskyInfo BloomskyStructure) GetRainIn() float64 {
	return bloomskyInfo.Storm.Rain
}

//GetRainRateIn returns rain in In
func (bloomskyInfo BloomskyStructure) GetRainRateIn() float64 {
	return bloomskyInfo.Storm.RainRate
}

//GetRainDailyMm returns rain daily in mm
func (bloomskyInfo BloomskyStructure) GetRainDailyMm() float64 {
	return bloomskyInfo.Storm.RainDaily
}

//GetRainMm returns total rain in mm
func (bloomskyInfo BloomskyStructure) GetRainMm() float64 {
	return bloomskyInfo.Storm.Rain
}

//GetRainRateMm returns rain in mm
func (bloomskyInfo BloomskyStructure) GetRainRateMm() float64 {
	return bloomskyInfo.Storm.RainRate
}

// NewBloomsky calls bloomsky and get structurebloomsky
func NewBloomsky(bloomskyURL, bloomskyToken string, debug bool) BloomskyStructure {

	var retry = 0
	var duration = time.Minute * 5

	if debug {
		debug = true
	}

	// get body from Rest API
	if debug {
		log.New(os.Stdout, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile).Printf("Get from Rest bloomsky API %s %s", bloomskyURL, bloomskyToken)
	}
	myRest := rest.MakeNew()

	b := []string{bloomskyToken}

	var headers map[string][]string
	headers = make(map[string][]string)
	headers["Authorization"] = b

	for retry < 5 {
		if err := myRest.GetWithHeaders(bloomskyURL, headers); err != nil {
			fmt.Printf("Problem with call rest, check the URL and the secret ID in the config file %v \n", err)
			retry++
			time.Sleep(duration)
		} else {
			retry = 5
		}
	}

	body := myRest.GetBody()
	return NewBloomskyFromBody(body)
}

// NewBloomskyFromBody to unit test with String
func NewBloomskyFromBody(body []byte) BloomskyStructure {
	var bloomskyInfo []BloomskyStructure
	//fmt.Println("Unmarshal the response")
	if err := json.Unmarshal(body, &bloomskyInfo); err != nil {
		log.Fatal(fmt.Errorf("Problem with json to struct, problem in the struct %v?", err))
	}
	bloomskyInfo[0].Data.TemperatureC = toFixed(((bloomskyInfo[0].Data.TemperatureF - 32.00) * 5.00 / 9.00), 2)
	bloomskyInfo[0].Data.Pressurehpa = toFixed((bloomskyInfo[0].Data.Pressure * 33.8638815), 2)

	bloomskyInfo[0].Storm.WindGustms = toFixed(bloomskyInfo[0].Storm.WindGust*0.44704, 2)
	bloomskyInfo[0].Storm.WindGustkmh = toFixed(bloomskyInfo[0].Storm.WindGust*1.60934, 2)
	bloomskyInfo[0].Storm.SustainedWindSpeedms = toFixed(bloomskyInfo[0].Storm.SustainedWindSpeed*0.44704, 2)
	bloomskyInfo[0].Storm.SustainedWindSpeedkmh = toFixed(bloomskyInfo[0].Storm.SustainedWindSpeed*1.60934, 2)

	bloomskyInfo[0].Storm.RainDailymm = toFixed(bloomskyInfo[0].Storm.SustainedWindSpeed*25.4, 2)
	bloomskyInfo[0].Storm.RainRatemm = toFixed(bloomskyInfo[0].Storm.SustainedWindSpeed*25.4, 2)

	bloomskyInfo[0].ShowPrettyAll()

	return bloomskyInfo[0]
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
