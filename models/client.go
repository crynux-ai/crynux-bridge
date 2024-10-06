package models

type Client struct {
	RootModel
	ClientId string `json:"client_id"`
}

type ClientTask struct {
	RootModel
	ClientID       uint            `json:"client_id"`
	Client         Client          `json:"-"`
	InferenceTasks []InferenceTask `json:"-"`
}
