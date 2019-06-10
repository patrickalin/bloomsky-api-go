package main

import (
	"fmt"

	bloomskyStructure "github.com/patrickalin/bloomsky-api-go"
)

const testFile1 = "testcase/test1.json"

func main() {
	fmt.Println("Example")
	// No url, no tocken, mock, no log
	mybloomskyExample := bloomskyStructure.New("", "", true, nil)
	// Refresh Data
	mybloomskyExample.Refresh()
	// use Getter
	fmt.Println("From mock file :> ", mybloomskyExample.GetDeviceID())
	// or GetStructure
	mybloomskyExampleStruct := mybloomskyExample.GetBloomskyStruct()
	fmt.Println("From bloomsky API :> ", mybloomskyExampleStruct.DeviceID)
	// Output:
	// 442C05954A59
	// 442C05954A59
}
