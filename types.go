package main

type Params struct {
	AccessKey  string               `json:"access_key"`
	SecretKey  string               `json:"secret_key"`
	Region     *string              `json:"region"`
	StackID    *string              `json:"stack_id"`
	AppID      *string              `json:"app_id"`
	Command    *string              `json:"command"`
	Comment    *string              `json:"comment"`
	CustomJSON *string              `json:"custom_json"`
	Arguments  map[string][]*string `json:"arguments"`
	Instances  []*string            `json:"instances"`
}
