package models

type Client struct {
	RootModel
	ClientId string `json:"client_id"`
}
