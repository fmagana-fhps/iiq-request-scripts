package models

type Location struct {
	Name                 string       `json:"Name"`
	LocationId           string       `json:"LocationId"`
	SiteId               string       `json:"SiteId"`
	AddressId            string       `json:"AddressId"`
	LocationTypeId       string       `json:"LocationTypeId"`
	LocationType         LocationType `json:"LocationType"`
	LocationStatusTypeId string       `json:"LocationStatusTypeId"`
	IsReservable         bool         `json:"IsReservable"`
	IsDeleted            bool         `json:"IsDeleted"`
	CreatedDate          string       `json:"CreatedDate"`
	ModifiedDate         string       `json:"ModifiedDate"`
}

type LocationType struct {
	LocationTypeId string `json:"LocationTypeId"`
	Name           string `json:"Name"`
	CreatedDate    string `json:"CreatedDate"`
	ModifiedDate   string `json:"ModifiedDate"`
	IsDeleted      bool   `json:"IsDeleted"`
	Scope          string `json:"Scope"`
}
