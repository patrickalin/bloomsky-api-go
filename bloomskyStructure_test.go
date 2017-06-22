package bloomskyStructure

import (
	"os"
	"testing"
	"time"
)

var mybloomsky BloomskyStructure
var mybloomsky1 BloomskyStructure

func TestMain(m *testing.M) {
	body := []byte("[{\"UTC\":2,\"CityName\":\"Thuin\",\"Storm\":{\"UVIndex\":\"1\",\"WindDirection\":\"E\",\"RainDaily\":0,\"WindGust\":0,\"SustainedWindSpeed\":0,\"RainRate\":0,\"24hRain\":0},\"Searchable\":true,\"DeviceName\":\"skyThuin\",\"RegisterTime\":1486905295,\"DST\":1,\"BoundedPoint\":\"\",\"LON\":4.3101,\"Point\":{},\"VideoList\":[\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-27.mp4\",\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-28.mp4\",\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-29.mp4\",\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-30.mp4\",\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-31.mp4\"],\"VideoList_C\":[\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-27_C.mp4\",\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-28_C.mp4\",\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-29_C.mp4\",\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-30_C.mp4\",\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-31_C.mp4\"],\"DeviceID\":\"442C05954A59\",\"NumOfFollowers\":2,\"LAT\":50.3394,\"ALT\":195,\"Data\":{\"Luminance\":9999,\"Temperature\":70.79,\"ImageURL\":\"http://s3-us-west-1.amazonaws.com/bskyimgs/faBiuZWsnpaoqZqrqJ1kr5uqmZammJw=.jpg\",\"TS\":1496345207,\"Rain\":false,\"Humidity\":64,\"Pressure\":29.41,\"DeviceType\":\"SKY2\",\"Voltage\":2611,\"Night\":false,\"UVIndex\":9999,\"ImageTS\":1496345207},\"FullAddress\":\"Drève des Alliés, Thuin, Wallonie, BE\",\"StreetName\":\"Drève des Alliés\",\"PreviewImageList\":[\"http://s3-us-west-1.amazonaws.com/bskyimgs/faBiuZWsnpaoqZqrqJ1kr5qwlZOmn5c=.jpg\",\"http://s3-us-west-1.amazonaws.com/bskyimgs/faBiuZWsnpaoqZqrqJ1kr5qwnZmqmZw=.jpg\",\"http://s3-us-west-1.amazonaws.com/bskyimgs/faBiuZWsnpaoqZqrqJ1kr5unnJakmZg=.jpg\",\"http://s3-us-west-1.amazonaws.com/bskyimgs/faBiuZWsnpaoqZqrqJ1kr5uom5Kkm50=.jpg\",\"http://s3-us-west-1.amazonaws.com/bskyimgs/faBiuZWsnpaoqZqrqJ1kr5upmZiqnps=.jpg\"]}]")

	mybloomsky = NewBloomskyFromBody(body)

	body = []byte("[{\"UTC\":2,\"CityName\":\"Thuin\",\"Storm\":{\"UVIndex\":\"1\",\"WindDirection\":\"E\",\"RainDaily\":0,\"WindGust\":0,\"SustainedWindSpeed\":0,\"RainRate\":0,\"24hRain\":0},\"Searchable\":true,\"DeviceName\":\"skyThuin\",\"RegisterTime\":1486905295,\"DST\":1,\"BoundedPoint\":\"\",\"LON\":4.3101,\"Point\":{},\"VideoList\":[\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-27.mp4\",\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-28.mp4\",\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-29.mp4\",\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-30.mp4\",\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-31.mp4\"],\"VideoList_C\":[\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-27_C.mp4\",\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-28_C.mp4\",\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-29_C.mp4\",\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-30_C.mp4\",\"http://s3.amazonaws.com/bskytimelapses/faBiuZWsnpaoqZqr_2_2017-05-31_C.mp4\"],\"DeviceID\":\"442C05954A59\",\"NumOfFollowers\":2,\"LAT\":50.3394,\"ALT\":195,\"Data\":{\"Luminance\":9999,\"Temperature\":70.79,\"ImageURL\":\"http://s3-us-west-1.amazonaws.com/bskyimgs/faBiuZWsnpaoqZqrqJ1kr5uqmZammJw=.jpg\",\"TS\":1496345207,\"Rain\":false,\"Humidity\":64,\"Pressure\":29.41,\"DeviceType\":\"SKY2\",\"Voltage\":2611,\"Night\":true,\"UVIndex\":9999,\"ImageTS\":1496345207},\"FullAddress\":\"Drève des Alliés, Thuin, Wallonie, BE\",\"StreetName\":\"Drève des Alliés\",\"PreviewImageList\":[\"http://s3-us-west-1.amazonaws.com/bskyimgs/faBiuZWsnpaoqZqrqJ1kr5qwlZOmn5c=.jpg\",\"http://s3-us-west-1.amazonaws.com/bskyimgs/faBiuZWsnpaoqZqrqJ1kr5qwnZmqmZw=.jpg\",\"http://s3-us-west-1.amazonaws.com/bskyimgs/faBiuZWsnpaoqZqrqJ1kr5unnJakmZg=.jpg\",\"http://s3-us-west-1.amazonaws.com/bskyimgs/faBiuZWsnpaoqZqrqJ1kr5uom5Kkm50=.jpg\",\"http://s3-us-west-1.amazonaws.com/bskyimgs/faBiuZWsnpaoqZqrqJ1kr5upmZiqnps=.jpg\"]}]")

	mybloomsky1 = NewBloomskyFromBody(body)

	os.Exit(m.Run())
}

func TestGetCity(t *testing.T) {
	if city := mybloomsky.GetCity(); city != "Thuin" {
		t.Errorf("Expected Thuin, but it was %s instead.", city)
	}
}

func TestDeviceId(t *testing.T) {
	if city := mybloomsky.GetDeviceID(); city != "442C05954A59" {
		t.Errorf("Expected 442C05954A59, but it was %s instead.", city)
	}
}

func TestTimestamp(t *testing.T) {
	if tt := mybloomsky.GetTimeStamp().Truncate(time.Minute); tt.Equal(time.Date(
		2017, 06, 01, 21, 26, 0, 0, time.UTC).Truncate(time.Minute)) {
		t.Errorf("Expected %s, but it was %s instead.", time.Date(
			2017, 06, 01, 19, 26, 0, 0, time.UTC).Truncate(time.Minute), tt)
	}
}

func TestNbrFollowers(t *testing.T) {
	if f := mybloomsky.GetNumOfFollowers(); f != 2 {
		t.Errorf("Expected 2, but it was %d instead.", f)
	}
}

func TestUV(t *testing.T) {
	if uv := mybloomsky.GetIndexUV(); uv != "1" {
		t.Errorf("Expected 1, but it was %s instead.", uv)
	}
}

func TestBloomskyStructure_IsNight(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   bool
	}{
		{"Day", mybloomsky, false},
		{"Night", mybloomsky1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bloomskyInfo := BloomskyStructure{
				Data: tt.fields.Data,
			}
			if got := bloomskyInfo.IsNight(); got != tt.want {
				t.Errorf("BloomskyStructure.IsNight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_IsNight2(t *testing.T) {
	tests := []struct {
		name  string
		night bool
		want  bool
	}{
		{"Day", false, false},
		{"Night", true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := BloomskyDataStructure{
				Night: tt.night,
			}
			bloomskyInfo := BloomskyStructure{
				Data: data,
			}
			if got := bloomskyInfo.IsNight(); got != tt.want {
				t.Errorf("BloomskyStructure.IsNight() = %v, want %v", got, tt.want)
			}
		})
	}
}
