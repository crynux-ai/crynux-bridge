package models

type LoraModel struct {
	RootModel
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Type         ModelType `json:"type"`
	DisplayLink  string    `json:"display_link"`
	DownloadLink string    `json:"download_link"`
}
