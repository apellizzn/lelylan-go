package main

import (
	"fmt"
  "os"
	"github.com/bradrydzewski/go.auth/oauth2"
	"github.com/joho/godotenv"
  "github.com/apellizzn/test/devices"
)

type Device struct {
		Id string
		Name string
		Type string
}

func main(){

	err := godotenv.Load()
  if err != nil {
    panic(err)
  }

	client := oauth2.Client {
		ClientId: os.Getenv("CLIENT_ID"),
    ClientSecret: os.Getenv("CLIENT_SECRET"),
    AccessTokenURL: "http://people.lelylan.com/oauth/token"}

  username := os.Getenv("USERNAME")
  password := os.Getenv("PASSWORD")

	t, err := client.GrantTokenPassword(username, password, "resources")
	if err == nil{
		devices.All(t)
		fmt.Println("Creating new device...")
		deletable := devices.Create("Go light","518be84900045e1521000007",t)
		devices.All(t)
		device := devices.One("54576b660dd0a49c9b000005", t)
		device.Update("Test name", t)
		fmt.Println("Deleting a device...")
		deletable.Delete(t)
		devices.All(t)
	} else {
		panic(err)
	}
}

