package views

type ResponseStruct struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}
