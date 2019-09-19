package main

import "github.com/ww24/lirc-web-api/service"

type status struct {
	Status       string `json:"status"`
	Message      string `json:"message"`
	Version      string `json:"version,omitempty"`
	LIRCDVersion string `json:"lircd_version,omitempty"`
}

type response struct {
	code    int
	Status  string           `json:"status"`
	Message string           `json:"message,omitempty"`
	Signals []service.Signal `json:"signals,omitempty"`
}

func (res *response) Error() string {
	return res.Message
}
