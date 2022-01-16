package main

type SalDat struct {
	Data SalReq `json:"data"`
}
type SalReq struct {
	CoverUrl string   `json:"cover_url"`
	Title    string   `json:"title"`
	Artists  []string `json:"artists"`
	Status   string   `json:"status"`
	Progress int      `json:"progress"`
	Duration int      `json:"duration"`
	AlbumUrl string   `json:"album_url"`
}
