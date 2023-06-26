package models

type SingleAsset struct {
	Asset Asset `json:"Item"`
}

type MultipleAssets struct {
	Assets []Asset `json:"Items"`
}
