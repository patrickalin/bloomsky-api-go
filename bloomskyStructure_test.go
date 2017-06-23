package bloomskyStructure

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

var mybloomskyTest1 BloomskyStructure
var mybloomskyTest2 BloomskyStructure

const testFile1 = "testCase/test1.json"
const testFile2 = "testCase/test2.json"

func readTestFile(fMock string) []byte {
	testFile, err := ioutil.ReadFile(fMock)
	if err != nil {
		logrus.Fatalf("Error in reading the file %s Err:  %v", fMock, err)
	}
	return testFile
}

func TestMain(m *testing.M) {
	mybloomskyTest1 = NewBloomskyFromBody(readTestFile(testFile1))
	mybloomskyTest2 = NewBloomskyFromBody(readTestFile(testFile2))

	os.Exit(m.Run())
}

func TestTimestamp(t *testing.T) {
	if tt := mybloomskyTest1.GetTimeStamp().Truncate(time.Minute); tt.Equal(time.Date(
		2017, 06, 01, 21, 26, 0, 0, time.UTC).Truncate(time.Minute)) {
		t.Errorf("Expected %s, but it was %s instead.", time.Date(
			2017, 06, 01, 19, 26, 0, 0, time.UTC).Truncate(time.Minute), tt)
	}
}

func TestNbrFollowers(t *testing.T) {
	if f := mybloomskyTest1.GetNumOfFollowers(); f != 2 {
		t.Errorf("Expected 2, but it was %d instead.", f)
	}
}

func TestBloomskyStructure_IsNight(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   bool
	}{
		{"Day", mybloomskyTest1, false},
		{"Night", mybloomskyTest2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bloomskyInfo := tt.fields
			if got := bloomskyInfo.IsNight(); got != tt.want {
				t.Errorf("BloomskyStructure.IsNight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetTimeStamp(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   time.Time
	}{
		{"Test1", mybloomskyTest1, time.Unix(int64(1496365207), 0)},
		{"Test2", mybloomskyTest2, time.Unix(int64(1496345207), 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bloomskyInfo := tt.fields
			if got := bloomskyInfo.GetTimeStamp(); got != tt.want {
				t.Errorf("BloomskyStructure.GetTimeStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetCity(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   string
	}{
		{"Test1", mybloomskyTest1, "Thuin"},
		{"Test2", mybloomskyTest2, "Paris"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bloomskyInfo := tt.fields
			if got := bloomskyInfo.GetCity(); got != tt.want {
				t.Errorf("BloomskyStructure.GetCity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetDeviceId(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   string
	}{
		{"Test1", mybloomskyTest1, "442C05954A59"},
		{"Test2", mybloomskyTest2, "442C05954A58"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bloomskyInfo := tt.fields
			if got := bloomskyInfo.GetDeviceID(); got != tt.want {
				t.Errorf("BloomskyStructure.GetDeviceID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetNumOfFollowers(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   int
	}{
		{"Test1", mybloomskyTest1, 2},
		{"Test2", mybloomskyTest2, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetNumOfFollowers(); got != tt.want {
				t.Errorf("BloomskyStructure.GetNumOfFollowers() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestNewBloomsky(t *testing.T) {
	type args struct {
		bloomskyURL   string
		bloomskyToken string
	}
	tests := []struct {
		name string
		args args
		want BloomskyStructure
	}{
		{"Error token", args{"https://api.bloomsky.com/api/skydata/", ""}, mybloomskyTest1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBloomsky(tt.args.bloomskyURL, tt.args.bloomskyToken); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBloomsky() = %v, want %v", got, tt.want)
			}
		})
	}
}*/

func TestNewBloomskyFromBody(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name string
		args args
		want BloomskyStructure
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBloomskyFromBody(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBloomskyFromBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUV(t *testing.T) {
	if uv := mybloomskyTest1.GetIndexUV(); uv != "1" {
		t.Errorf("Expected 1, but it was %s instead.", uv)
	}
}

func TestBloomskyStructure_GetIndexUV(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   string
	}{
		{"Test1", mybloomskyTest1, "1"},
		{"Test2", mybloomskyTest2, "3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetIndexUV(); got != tt.want {
				t.Errorf("BloomskyStructure.GetIndexUV() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestBloomskyStructure_GetTemperatureFahrenheit(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   float64
	}{
		{"Test1", mybloomskyTest1, 70.79},
		{"Test2", mybloomskyTest2, 65.79},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetTemperatureFahrenheit(); got != tt.want {
				t.Errorf("BloomskyStructure.GetTemperatureFahrenheit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetTemperatureCelsius(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   float64
	}{
		{"Test1", mybloomskyTest1, 21.55},
		{"Test2", mybloomskyTest2, 18.77},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetTemperatureCelsius(); got != tt.want {
				t.Errorf("BloomskyStructure.GetTemperatureCelsius() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetHumidity(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   float64
	}{
		{"Test1", mybloomskyTest1, 64},
		{"Test2", mybloomskyTest2, 43},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetHumidity(); got != tt.want {
				t.Errorf("BloomskyStructure.GetHumidity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetPressureInHg(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   float64
	}{
		{"Test1", mybloomskyTest1, 29.41},
		{"Test2", mybloomskyTest2, 49.41},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetPressureInHg(); got != tt.want {
				t.Errorf("BloomskyStructure.GetPressureInHg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetPressureHPa(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   float64
	}{
		{"Test1", mybloomskyTest1, 995.94},
		{"Test2", mybloomskyTest2, 1673.21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetPressureHPa(); got != tt.want {
				t.Errorf("BloomskyStructure.GetPressureHPa() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetWindDirection(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   string
	}{
		{"Test1", mybloomskyTest1, "E"},
		{"Test2", mybloomskyTest2, "W"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetWindDirection(); got != tt.want {
				t.Errorf("BloomskyStructure.GetWindDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetWindGustMph(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   float64
	}{
		{"Test1", mybloomskyTest1, 0},
		{"Test2", mybloomskyTest2, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetWindGustMph(); got != tt.want {
				t.Errorf("BloomskyStructure.GetWindGustMph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetWindGustMs(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   float64
	}{
		{"Test1", mybloomskyTest1, 0},
		{"Test2", mybloomskyTest2, 33.81},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetWindGustMs(); got != tt.want {
				t.Errorf("BloomskyStructure.GetWindGustMs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetSustainedWindSpeedMph(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   float64
	}{
		{"Test1", mybloomskyTest1, 0},
		{"Test2", mybloomskyTest2, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetSustainedWindSpeedMph(); got != tt.want {
				t.Errorf("BloomskyStructure.GetSustainedWindSpeedMph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetSustainedWindSpeedMs(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   float64
	}{
		{"Test1", mybloomskyTest1, 0},
		{"Test2", mybloomskyTest2, 19.32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetSustainedWindSpeedMs(); got != tt.want {
				t.Errorf("BloomskyStructure.GetSustainedWindSpeedMs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_IsRain(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   bool
	}{
		{"Test1", mybloomskyTest1, true},
		{"Test2", mybloomskyTest2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.IsRain(); got != tt.want {
				t.Errorf("BloomskyStructure.IsRain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetRainDailyIn(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   float64
	}{
		{"Test1", mybloomskyTest1, 0},
		{"Test2", mybloomskyTest2, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetRainDailyIn(); got != tt.want {
				t.Errorf("BloomskyStructure.GetRainDailyIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetRainIn(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   float64
	}{
		{"Test1", mybloomskyTest1, 0},
		{"Test2", mybloomskyTest2, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetRainIn(); got != tt.want {
				t.Errorf("BloomskyStructure.GetRainIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetRainRateIn(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   float64
	}{
		{"Test1", mybloomskyTest1, 0},
		{"Test2", mybloomskyTest2, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetRainRateIn(); got != tt.want {
				t.Errorf("BloomskyStructure.GetRainRateIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetRainDailyMm(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   float64
	}{
		{"Test1", mybloomskyTest1, 0},
		{"Test2", mybloomskyTest2, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetRainDailyMm(); got != tt.want {
				t.Errorf("BloomskyStructure.GetRainDailyMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetRainMm(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   float64
	}{
		{"Test1", mybloomskyTest1, 0},
		{"Test2", mybloomskyTest2, 406.4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetRainMm(); got != tt.want {
				t.Errorf("BloomskyStructure.GetRainMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetRainRateMm(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   float64
	}{
		{"Test1", mybloomskyTest1, 0},
		{"Test2", mybloomskyTest2, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetRainRateMm(); got != tt.want {
				t.Errorf("BloomskyStructure.GetRainRateMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetSustainedWindSpeedkmh(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   float64
	}{
		{"Test1", mybloomskyTest1, 0},
		{"Test2", mybloomskyTest2, 19.31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetSustainedWindSpeedkmh(); got != tt.want {
				t.Errorf("BloomskyStructure.GetSustainedWindSpeedkmh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomskyStructure_GetWindGustkmh(t *testing.T) {
	tests := []struct {
		name   string
		fields BloomskyStructure
		want   float64
	}{
		{"Test1", mybloomskyTest1, 0},
		{"Test2", mybloomskyTest2, 33.8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetWindGustkmh(); got != tt.want {
				t.Errorf("BloomskyStructure.GetWindGustkmh() = %v, want %v", got, tt.want)
			}
		})
	}
}
