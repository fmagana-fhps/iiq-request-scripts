package models

type AssetVerificationResponse struct {
	AssetVerification AssetVerification `json:"Item"`
	Id                int               `json:"Id"`
	Uid               string            `json:"Uid"`
}

type AssetVerification struct {
	AssetVerificationId     string `json:"AssetVerificationId"`
	AssetId                 string `json:"AssetId"`
	CreatedDate             string `json:"CreatedDate"`
	VerifiedByUserId        string `json:"VerifiedByUserId"`
	AssetVerificationTypeId string `json:"AssetVerificationTypeId"`
	IsSuccessful            bool   `json:"IsSuccessful"`
	Comments                string `json:"Comments"`
	LocationId              string `json:"LocationId"`
}
