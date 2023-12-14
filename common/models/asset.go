// Imported over
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
	OwnerID                         string   `json:"OwnerId"`
	PreviousOwnerID                 string   `json:"PreviousOwnerId"`
	Owner                           Owner    `json:"Owner"`
	PreviousOwner                   Owner    `json:"PreviousOwner"`
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
	WarrantyExpirationDate          string   `json:"WarrantyExpirationDate"`
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

type Owner struct {
	UserID     string `json:"UserId"`
	SiteID     string `json:"SiteId"`
	Name       string `json:"Name"`
	FullName   string `json:"FullName"`
	Initials   string `json:"Initials"`
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	LocationID string `json:"LocationId"`
	RoleID     string `json:"RoleId"`
	RoleName   string `json:"RoleName"`
	Email      string `json:"Email"`
	Username   string `json:"Username"`
	IsOnline   bool   `json:"IsOnline"`
}
