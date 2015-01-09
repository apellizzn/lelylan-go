package devices

type Device struct {
	Name string `json:"name"`
	Id string `json:"id"`
	Uri string `json:"uri"`
	Cathegory string `json:"cathegory"`
	Pending bool `json:"pending"`
	Activated bool `json:"activated"`
	Type Type `json:"type"`
}

type Type struct{
	Id string `json:"id"`
}