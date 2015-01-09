package devices

import "github.com/bradrydzewski/go.auth/oauth2"
import "github.com/apellizzn/test/comunication"
import "encoding/json"
import "fmt"

type Device struct {
	Name string 
	Id string
	Uri string
	Cathegory string
	Pending bool
	Activated bool
}

func (device *Device) Delete(token *oauth2.Token) bool{
	headers := map[string]string { "Authorization": "Bearer " + token.Token() }
	raw := comunication.Perform("DELETE", "http://api.lelylan.com/devices/" + device.Id, headers, nil)
	err := json.Unmarshal(raw, &device)
	if err == nil {
		fmt.Printf("%s deleted!\n", device.Name)
		device = nil
		return true
	} else {
		panic(err)
	}
}

func (device *Device) Update(name string, token *oauth2.Token) {
	headers := map[string]string { 
		"Authorization": "Bearer " + token.Token(),
		"Content-Type": "application/json",
		"Accept": "application/json"}

	body := []byte("{ \"name\": \""+ name +"\" }")
	raw := comunication.Perform("PUT", "http://api.lelylan.com/devices/" + device.Id, headers, body)
	err := json.Unmarshal(raw, &device)
	if err == nil {
		fmt.Printf("%s updated!\n", device.Name)
	} else {
		panic(err)
	}
}

func One(id string, token *oauth2.Token) Device {
	headers := map[string]string { "Authorization": "Bearer " + token.Token() }
	var device Device
	raw := comunication.Perform("GET", "http://api.lelylan.com/devices/" + id, headers, nil)
	err := json.Unmarshal(raw, &device)
	if err == nil {
		fmt.Printf("%s retrived!\n", device.Name)
		return device
	} else {
		panic(err)
	}
}

func All(token *oauth2.Token) []Device {
	headers := map[string]string { "Authorization": "Bearer " + token.Token() }
	var devices []Device
	raw := comunication.Perform("GET", "http://api.lelylan.com/devices", headers, nil)
	err := json.Unmarshal(raw, &devices)
	if err == nil {
		fmt.Printf("You own %d devices!\n", len(devices))
		return devices
	} else {
		panic(err)
	}
}

func Create(name string, id string, token *oauth2.Token) Device{
	headers := map[string]string { 
		"Authorization": "Bearer " + token.Token(),
		"Content-Type": "application/json",
		"Accept": "application/json"}

	body := []byte("{ \"name\": \""+ name +"\", \"type\": { \"id\": \""+id+"\" }}")
	var device Device
	raw := comunication.Perform("POST", "http://api.lelylan.com/devices", headers, body)
	err := json.Unmarshal(raw, &device)
	if err == nil {
		fmt.Printf("%s created!\n", device.Name)
		return device
	} else {
		panic(err)
	}
}