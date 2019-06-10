package main

import (
	"fmt"

	bloomskyStructure "github.com/patrickalin/bloomsky-api-go"
)

func main() {
	fmt.Println("Example")
	// No url, no tocken, mock, no log
	mybloomskyExample := bloomskyStructure.New("", "", true, nil)

	// Refresh Data
	mybloomskyExample.Refresh()

	// use Getter
	fmt.Println("DeviceId from mock file :> ", mybloomskyExample.GetDeviceID())

	// or GetStructure
	mybloomskyExampleStruct := mybloomskyExample.GetBloomskyStruct()
	fmt.Println("DeviceId from bloomsky API :> ", mybloomskyExampleStruct.DeviceID)
	fmt.Println("Win direction from Struct bloomsky API :> ", mybloomskyExampleStruct.Storm.WindDirection)
	fmt.Println("Win direction degree from API bloomsky API :> ", mybloomskyExample.GetWindDirectionDeg())
	fmt.Println("Win direction from API bloomsky API :> ", mybloomskyExample.GetWindDirection())

}
