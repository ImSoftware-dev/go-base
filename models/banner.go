package models

type BannerModel struct {
	BannerID   int    `json:"banner_id"`
	BannerUrl  string `json:"banner_url"`
	CreateDate string `json:"create_date"`
	LinkUrl    string `json:"link_url"`
}
