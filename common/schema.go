package common

type Response struct {
	Item Item `json:"Item"`
}

type Item struct {
	AssetId                         string   `json:"AssetId"`
	SiteId                          string   `json:"SiteId"`
	Site                            Site     `json:"Site"`
	ProductId                       string   `json:"ProductId"`
	CreatedDate                     string   `json:"CreatedDate"`
	ModifiedDate                    string   `json:"ModifiedDate"`
	AssetTypeId                     string   `json:"AssetTypeId"`
	AssetTypeName                   string   `json:"AssetTypeName"`
	IsDeleted                       bool     `json:"IsDeleted"`
	StatusTypeId                    string   `json:"StatusTypeId"`
	Status                          Status   `json:"Status"`
	AssetTag                        string   `json:"AssetTag"`
	SerialNumber                    string   `json:"SerialNumber"`
	ExternalId                      string   `json:"ExternalId"`
	Name                            string   `json:"Name"`
	CanOwnerManage                  bool     `json:"CanOwnerManage"`
	CanSubmitTicket                 bool     `json:"CanSubmitTicket"`
	IsFavorite                      bool     `json:"IsFavorite"`
	ModelId                         string   `json:"ModelId"`
	Model                           Model    `json:"Model"`
	LocationId                      string   `json:"LocationId"`
	Location                        Location `json:"Location"`
	HasOpenTickets                  bool     `json:"HasOpenTickets"`
	OpenTickets                     uint     `json:"OpenTickets"`
	PurchasePrice                   float32  `json:"PurchasePrice"`
	PurchasePoNumber                string   `json:"PurchasePoNumber"`
	WarrantyInfo                    string   `json:"WarrantyInfo"`
	LastInventoryDate               string   `json:"LastInventoryDate"`
	IsReadOnly                      bool     `json:"IsReadOnly"`
	IsExternallyManaged             bool     `json:"IsExternallyManaged"`
	AssetAuditPolicyStatusSortOrder uint     `json:"AssetAuditPolicyStatusSortOrder"`
	LastVerificationSuccessful      bool     `json:"LastVerificationSuccessful"`
}

type Site struct {
	SiteId                 string `json:"SiteId"`
	ProductId              string `json:"ProductId"`
	CreatedDate            string `json:"CreatedDate"`
	Name                   string `json:"Name"`
	OffsetInMinutesFromUtc uint   `json:"OffsetInMinutesFromUtc"`
	WorkMinutesPerDay      uint   `json:"WorkMinutesPerDay"`
	DbClusterIndex         uint   `json:"DbClusterIndex"`
	StatusTypeId           string `json:"StatusTypeId"`
	WorkHoursStartUtc      string `json:"WorkHoursStartUtc"`
	WorkHoursEndUtc        string `json:"WorkHoursEndUtc"`
	SystemUserId           string `json:"SystemUserId"`
	IsSandbox              bool   `json:"IsSandbox"`
}

type Status struct {
	AssetStatusTypeId string `json:"AssetStatusTypeId"`
	Scope             string `json:"Scope"`
	Name              string `json:"Name"`
	IsRetired         bool   `json:"IsRetired"`
}

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
