package models

type Asset struct {
	AssetId                         string   `json:"AssetId"`
	SiteId                          string   `json:"SiteId"`
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

type Status struct {
	AssetStatusTypeId string `json:"AssetStatusTypeId"`
	Scope             string `json:"Scope"`
	Name              string `json:"Name"`
	IsRetired         bool   `json:"IsRetired"`
}
