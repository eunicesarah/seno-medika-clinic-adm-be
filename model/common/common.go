package common

type Response struct {
	Message    string      `json:"message"`
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

type PutInput struct {
	Key   interface{} `json:"key"`
	Value interface{} `json:"value"`
}
