package views

type ResponseStruct struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
