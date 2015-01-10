package lelylan

import "github.com/bradrydzewski/go.auth/oauth2"
import "github.com/apellizzn/lelylan-go/lelylan/api"
import "github.com/apellizzn/lelylan-go/lelylan/devices"
import "fmt"

const ENDPOINT = "http://api.lelylan.com/devices/" 

type Client struct{
	Token *oauth2.Token
}

func (client *Client) Devices() []devices.Device{
	headers := map[string]string{
		"Authorization": "Bearer "+ client.Token.Token() }
	response, err := api.Get(ENDPOINT, headers)
	if err.Status != 0 {
  	fmt.Println(err)
  }
	var devices []devices.Device
	response.Unmarshal(&devices)
	return devices
}

func (client *Client) Device(id string) devices.Device{
	headers := map[string]string{
		"Authorization": "Bearer "+ client.Token.Token() }
	response, err := api.Get(ENDPOINT + id, headers)
	if err.Status != 0 {
  	fmt.Println(err)
  }
	var device devices.Device
	response.Unmarshal(&device)
	return device
}

func (client *Client) CreateDevice(name string, typeId string) devices.Device{
	headers := map[string]string{
		"Authorization": "Bearer "+ client.Token.Token(),
		"Content-Type": "application/json",
		"Accept": "application/json" }
	device := devices.Device{ Name: name, Type: devices.Type{ Id: typeId }}
	response, err := api.Post(ENDPOINT, headers, device)
	if err.Status != 0 {
  	fmt.Println(err)
  }
	response.Unmarshal(&device)
	fmt.Println(device)
	return device
}

func (client *Client) UpdateDevice(name string, id string) devices.Device{
	headers := map[string]string{
		"Authorization": "Bearer "+ client.Token.Token(),
		"Content-Type": "application/json",
		"Accept": "application/json" }
	device := devices.Device{ Name: name }
	response, err := api.Put(ENDPOINT + id, headers, device)
	if err.Status != 0 {
  	fmt.Println(err)
  }
	response.Unmarshal(&device)
	return device
}

func (client *Client) DeleteDevice(id string) devices.Device{
	headers := map[string]string{
		"Authorization": "Bearer "+ client.Token.Token()}
	var device devices.Device
	response, err := api.Delete(ENDPOINT + id, headers)
	if err.Status != 0 {
  	fmt.Println(err)
  }
	response.Unmarshal(&device)
	return device
}