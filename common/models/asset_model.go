package models

type Model struct {
	ModelId        string       `json:"ModelId"`
	Scope          string       `json:"Scope"`
	Name           string       `json:"Name"`
	ModelName      string       `json:"ModelName"`
	AssetTypeId    string       `json:"AssetTypeId"`
	ManufacturerId string       `json:"ManufacturerId"`
	Manufacturer   Manufacturer `json:"Manufacturer"`
	CategoryId     string       `json:"CategoryId"`
	Category       Category     `json:"Category"`
	ImageId        string       `json:"ImageId"`
	CreatedDate    string       `json:"CreatedDate"`
	ModifiedDate   string       `json:"ModifiedDate"`
	EsoIsVisible   bool         `json:"EsoIsVisible"`
	IsGenericModel bool         `json:"IsGenericModel"`
}

type Manufacturer struct {
	ManufacturerId string `json:"ManufacturerId"`
	Scope          string `json:"Scope"`
	Name           string `json:"Name"`
	EsoIsVisible   bool   `json:"EsoIsVisible"`
}

type Category struct {
	Scope          string `json:"Scope"`
	CategoryId     string `json:"CategoryId"`
	Level          uint   `json:"Level"`
	SortOrder      uint   `json:"SortOrder"`
	CategoryTypeId string `json:"CategoryTypeId"`
	Name           string `json:"Name"`
	Icon           string `json:"Icon"`
}
