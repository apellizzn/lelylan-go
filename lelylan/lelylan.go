package lelylan

import "github.com/bradrydzewski/go.auth/oauth2"
import "github.com/apellizzn/lelylan-go/api"
import "github.com/apellizzn/lelylan-go/devices"
import "fmt"

const ENDPOINT = "http://api.lelylan.com/devices/" 

type Client struct{
	Token *oauth2.Token
}

func (client *Client) Devices() []devices.Device{
	headers := map[string]string{
		"Authorization": "Bearer "+ client.Token.Token() }
	response := api.Get(ENDPOINT,headers)
	var devices []devices.Device
	response.Unmarshal(&devices)
	fmt.Println(devices)
	return devices
}

func (client *Client) Device(id string) devices.Device{
	headers := map[string]string{
		"Authorization": "Bearer "+ client.Token.Token() }
	response := api.Get(ENDPOINT + id, headers)
	var device devices.Device
	response.Unmarshal(&device)
	fmt.Println(device)
	return device
}

func (client *Client) CreateDevice(name string, typeId string) devices.Device{
	headers := map[string]string{
		"Authorization": "Bearer "+ client.Token.Token(),
		"Content-Type": "application/json",
		"Accept": "application/json" }
	device := devices.Device{ Name: name, Type: devices.Type{ Id: typeId }}
	response := api.Post(ENDPOINT, headers, device)
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
	response := api.Put(ENDPOINT + id, headers, device)
	response.Unmarshal(&device)
	fmt.Println(device)
	return device
}

func (client *Client) DeleteDevice(id string) devices.Device{
	headers := map[string]string{
		"Authorization": "Bearer "+ client.Token.Token()}
	var device devices.Device
	response := api.Delete(ENDPOINT + id, headers)
	response.Unmarshal(&device)
	fmt.Println(device)
	return device
}