package helpers

import "log"

type Response struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Totalrecord int         `json:"totalrecord"`
	Record      interface{} `json:"record"`
	Time        string      `json:"time"`
}
type Responsesetting struct {
	Status          int    `json:"status"`
	Website_status  string `json:"website_status"`
	Website_message string `json:"website_message"`
	Time            string `json:"time"`
}
type ResponseCustom struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Permainan   interface{} `json:"permainan"`
	Totalrecord int         `json:"totalrecord"`
	Totalbayar  int         `json:"totalbayar"`
	Record      interface{} `json:"record"`
	Time        string      `json:"time"`
}
type ErrorResponse struct {
	Field string
	Tag   string
}

func ErrorCheck(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}
