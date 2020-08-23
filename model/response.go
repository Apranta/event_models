package models

type Status struct {
	Code int `json:"code"`
	MessageClient string `json:"message_client"`
	MessageServer string `json:"message_server"`
}

type MetaList struct {
	CurrentPage int `json:"current_page"`
	PerPage int `json:"per_page"`
	Total int `json:"total"`
	TotalPage int `json:"total_page"`
}

type Response struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
	Status Status `json:"status"`
}

type UnprocessableEntityResponse struct {
	Details interface{} `json:"details"`
	Message string `json:"message"`
}