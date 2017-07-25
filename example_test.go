package bloomskyStructure_test

import (
	"fmt"

	"github.com/patrickalin/bloomsky-api-go"
)

const testFile1 = "testcase/test1.json"

func ExampleNew() {
	// No url, no tocken, mock, no log
	mybloomskyExample := bloomskyStructure.New("", "", true, nil)
	// Refresh Data
	mybloomskyExample.Refresh()
	// use Getter
	fmt.Println(mybloomskyExample.GetDeviceID())
	// or GetStructure
	mybloomskyExampleStruct := mybloomskyExample.GetBloomskyStruct()
	fmt.Println(mybloomskyExampleStruct.DeviceID)
	// Output:
	// 442C05954A59
	// 442C05954A59
}
