// Package bloomskyStructure calls rest API Bloomsky, puts it in the structure and gives somes functions
package bloomskyStructure

import (
	"encoding/json"
	"math"
	"os"
	"time"

	http "github.com/patrickalin/http-go"
	"github.com/sirupsen/logrus"
)

type bloomsky struct {
	url               string
	token             string
	BloomskyStructure BloomskyStructure
}

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
	LastCall         string
}

// BloomskyStormStructure represents the structure STORM of the JSON return by the API
type BloomskyStormStructure struct {
	UVIndex               string  `json:"UVIndex"`
	WindDirection         string  `json:"WindDirection"`
	WindGust              float64 `json:"WindGust"`
	WindGustms            float64
	WindGustkmh           float64
	SustainedWindSpeed    float64 `json:"SustainedWindSpeed"`
	SustainedWindSpeedms  float64
	SustainedWindSpeedkmh float64
	Rain                  float64
	RainDaily             float64 `json:"RainDaily"`
	RainDailymm           float64
	RainRate              float64 `json:"RainRate"`
	RainRatemm            float64
	Rainin                float64 `json:"24hRain"`
	Rainmm                float64
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

// Bloomsky is the interface BloomskyStructure
type Bloomsky interface {
	GetDeviceID() string
	GetHumidity() float64
	GetCity() string
	RefreshFromRest()
	RefreshFromBody(body []byte)
	GetNumOfFollowers() int
	IsNight() bool
	GetPressureHPa() float64
	GetWindDirection() string
	GetTimeStamp() time.Time
	GetIndexUV() string
	GetTemperatureFahrenheit() float64
	GetTemperatureCelsius() float64
	GetPressureInHg() float64
	GetWindGustMph() float64
	GetWindGustMs() float64
	GetSustainedWindSpeedMs() float64
	GetSustainedWindSpeedMph() float64
	IsRain() bool
	GetRainDailyIn() float64
	GetRainRateIn() float64
	GetRainIn() float64
	GetRainDailyMm() float64
	GetRainRateMm() float64
	GetWindGustkmh() float64
	GetSustainedWindSpeedkmh() float64
	GetRainMm() float64
	GetBloomskyStruct() BloomskyStructure
	GetLastCall() string
	GetTS() float64
}

const logFile = "bloomskyapi.log"

var log *logrus.Logger

var rest http.HTTP

// New calls bloomsky and get structurebloomsky
func New(bloomskyURL, bloomskyToken string, l *logrus.Logger) Bloomsky {
	initLog(l)

	var b bloomsky

	log.WithFields(logrus.Fields{
		"url": bloomskyURL,
		"fct": "bloomskyStructure.New",
	}).Debug("New bloomsky")

	b.token = bloomskyToken
	b.url = bloomskyURL

	rest = http.New(log)

	return &b
}

//Init the logger
func initLog(l *logrus.Logger) {
	if l != nil {
		log = l
		log.WithFields(logrus.Fields{
			"fct": "bloomskyStructure.initLog",
		}).Debug("Use the logger pass in New")
		return
	}

	log = logrus.New()

	log.WithFields(logrus.Fields{
		"fct": "bloomskyStructure.initLog",
	}).Debug("Create new logger")

	log.Formatter = new(logrus.TextFormatter)

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": err,
			"fct": "bloomskyStructure.initLog",
		}).Fatal("Failed to log to file, using default stderr")
	}
	log.Out = file
}

//Call rest and refresh the structure
func (bloomsky *bloomsky) RefreshFromRest() {
	tock := []string{bloomsky.token}

	var headers map[string][]string
	headers = make(map[string][]string)
	headers["Authorization"] = tock

	var retry = 0
	for retry < 5 {
		if err := rest.GetWithHeaders(bloomsky.url, headers); err != nil {
			log.WithFields(logrus.Fields{
				"url":   bloomsky.url,
				"Error": err,
			}).Error("Problem with call rest, check the URL and the secret ID in the config file")
			retry++
			time.Sleep(time.Minute * 5)
		} else {
			retry = 5
		}
	}

	bloomsky.RefreshFromBody(rest.GetBody())
}

//Refresh from body without call rest
func (bloomsky *bloomsky) RefreshFromBody(body []byte) {
	var bloomskyArray []BloomskyStructure
	if err := json.Unmarshal(body, &bloomskyArray); err != nil {
		log.WithFields(logrus.Fields{
			"body": body,
			"msg":  err,
		}).Fatal("Problem with json to struct")
	}
	bloomsky.BloomskyStructure = bloomskyArray[0]
	bloomsky.BloomskyStructure.Data.TemperatureC = toFixed(((bloomsky.BloomskyStructure.Data.TemperatureF - 32.00) * 5.00 / 9.00), 2)
	bloomsky.BloomskyStructure.Data.Pressurehpa = toFixed((bloomsky.BloomskyStructure.Data.Pressure * 33.8638815), 2)

	bloomsky.BloomskyStructure.Storm.WindGustms = toFixed(bloomsky.BloomskyStructure.Storm.WindGust*0.44704, 2)
	bloomsky.BloomskyStructure.Storm.WindGustkmh = toFixed(bloomsky.BloomskyStructure.Storm.WindGust*1.60934, 2)
	bloomsky.BloomskyStructure.Storm.SustainedWindSpeedms = toFixed(bloomsky.BloomskyStructure.Storm.SustainedWindSpeed*0.44704, 2)
	bloomsky.BloomskyStructure.Storm.SustainedWindSpeedkmh = toFixed(bloomsky.BloomskyStructure.Storm.SustainedWindSpeed*1.60934, 2)

	bloomsky.BloomskyStructure.Storm.RainDailymm = toFixed(bloomsky.BloomskyStructure.Storm.RainDaily*25.4, 2)
	bloomsky.BloomskyStructure.Storm.RainRatemm = toFixed(bloomsky.BloomskyStructure.Storm.RainRate*25.4, 2)
	bloomsky.BloomskyStructure.Storm.Rainmm = toFixed(bloomsky.BloomskyStructure.Storm.Rainin*25.4, 2)
	bloomsky.BloomskyStructure.LastCall = time.Now().Format("2006-01-02 15:04:05")

	log.WithFields(logrus.Fields{
		"time": bloomsky.BloomskyStructure.LastCall,
		"fct":  "bloomskyStructure.RefreshFromBody",
	}).Debug("Refresh From Body")

	bloomsky.ShowPrettyAll()

}

func (bloomsky *bloomsky) GetBloomskyStruct() BloomskyStructure {
	return bloomsky.BloomskyStructure
}

//GetTimeStamp returns the timestamp give by Bloomsky
func (bloomsky *bloomsky) GetTimeStamp() time.Time {
	return time.Unix(int64(bloomsky.BloomskyStructure.Data.TS), 0)
}

//GetCity returns the city name
func (bloomsky *bloomsky) GetCity() string {
	return bloomsky.BloomskyStructure.CityName
}

//GetDeviceID returns the Device Id
func (bloomsky *bloomsky) GetDeviceID() string {
	return bloomsky.BloomskyStructure.DeviceID
}

//GetNumOfFollowers returns the number of followers
func (bloomsky *bloomsky) GetNumOfFollowers() int {
	return int(bloomsky.BloomskyStructure.NumOfFollowers)
}

//GetIndexUV returns the UV index from 1 to 11
func (bloomsky *bloomsky) GetIndexUV() string {
	return bloomsky.BloomskyStructure.Storm.UVIndex
}

//IsNight returns true if it's the night
func (bloomsky *bloomsky) IsNight() bool {
	return bloomsky.BloomskyStructure.Data.Night
}

//GetTemperatureFahrenheit returns temperature in Fahrenheit
func (bloomsky *bloomsky) GetTemperatureFahrenheit() float64 {
	return bloomsky.BloomskyStructure.Data.TemperatureF
}

//GetTemperatureCelsius returns temperature in Celsius
func (bloomsky *bloomsky) GetTemperatureCelsius() float64 {
	return bloomsky.BloomskyStructure.Data.TemperatureC
}

//GetHumidity returns humidity %
func (bloomsky *bloomsky) GetHumidity() float64 {
	return bloomsky.BloomskyStructure.Data.Humidity
}

//GetPressureHPa returns pressure in HPa
func (bloomsky *bloomsky) GetPressureHPa() float64 {
	return bloomsky.BloomskyStructure.Data.Pressurehpa
}

//GetPressureInHg returns pressure in InHg
func (bloomsky *bloomsky) GetPressureInHg() float64 {
	return bloomsky.BloomskyStructure.Data.Pressure
}

//GetWindDirection returns wind direction (N,S,W,E, ...)
func (bloomsky *bloomsky) GetWindDirection() string {
	return bloomsky.BloomskyStructure.Storm.WindDirection
}

//GetWindGustMph returns Wind in Mph
func (bloomsky *bloomsky) GetWindGustMph() float64 {
	return bloomsky.BloomskyStructure.Storm.WindGust
}

//GetWindGustMs returns Wind in Ms
func (bloomsky *bloomsky) GetWindGustMs() float64 {
	return (bloomsky.BloomskyStructure.Storm.WindGust * 1.61)
}

//GetSustainedWindSpeedMph returns Sustained Wind Speed in Mph
func (bloomsky *bloomsky) GetSustainedWindSpeedMph() float64 {
	return bloomsky.BloomskyStructure.Storm.SustainedWindSpeed
}

//GetSustainedWindSpeedMs returns Sustained Wind Speed in Ms
func (bloomsky *bloomsky) GetSustainedWindSpeedMs() float64 {
	return (bloomsky.BloomskyStructure.Storm.SustainedWindSpeed * 1.61)
}

//IsRain returns true if it's rain
func (bloomsky *bloomsky) IsRain() bool {
	return bloomsky.BloomskyStructure.Data.Rain
}

//GetRainDailyIn returns rain daily in In
func (bloomsky *bloomsky) GetRainDailyIn() float64 {
	return bloomsky.BloomskyStructure.Storm.RainDaily
}

//GetRainIn returns total rain in In
func (bloomsky *bloomsky) GetRainIn() float64 {
	return bloomsky.BloomskyStructure.Storm.Rainin
}

//GetRainRateIn returns rain in In
func (bloomsky *bloomsky) GetRainRateIn() float64 {
	return bloomsky.BloomskyStructure.Storm.RainRate
}

//GetRainDailyMm returns rain daily in mm
func (bloomsky *bloomsky) GetRainDailyMm() float64 {
	return bloomsky.BloomskyStructure.Storm.RainDaily
}

//GetRainMm returns total rain in mm
func (bloomsky *bloomsky) GetRainMm() float64 {
	return bloomsky.BloomskyStructure.Storm.Rainmm
}

//GetRainRateMm returns rain in mm
func (bloomsky *bloomsky) GetRainRateMm() float64 {
	return bloomsky.BloomskyStructure.Storm.RainRate
}

//GetSustainedWindSpeedkmh returns Sustained Wind in Km/h
func (bloomsky *bloomsky) GetSustainedWindSpeedkmh() float64 {
	return bloomsky.BloomskyStructure.Storm.SustainedWindSpeedkmh
}

//GetWindGustkmh returns Wind in Km/h
func (bloomsky *bloomsky) GetWindGustkmh() float64 {
	return bloomsky.BloomskyStructure.Storm.WindGustkmh
}

func (bloomsky *bloomsky) GetLastCall() string {
	return bloomsky.BloomskyStructure.LastCall
}

func (bloomsky *bloomsky) GetTS() float64 {
	return bloomsky.BloomskyStructure.Data.TS
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

// ShowPrettyAll prints to the console the JSON
func (bloomsky *bloomsky) ShowPrettyAll() {
	out, err := json.Marshal(bloomsky)
	if err != nil {
		log.WithFields(logrus.Fields{
			"bloomsky": out,
			"fct":      "bloomskyStructure.ShowPrettyAll",
		}).Fatal("Error with parsing Json")
	}
}
