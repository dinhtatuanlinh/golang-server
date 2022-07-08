package template

type ResponseStruct struct {
	Data interface{}
	Time string
}

type DataMessage2xx struct {
	Status string
}

type DataErrorMessage4xx struct {
	Code    int `json:"Code" example:"6"`
	Message interface{}
	Time    string `json:"Time" example:"2021-12-10 17:53:16.279818"`
}

type DataErrorMessage4xxWithCustomerData struct {
	Code int
	Data interface{}
	Time string
}

type DataErrorMessage5xx struct {
	Message interface{}
}
