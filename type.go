package main

import "time"

type signal struct {
	Remote string `json:"remote"`
	Name   string `json:"name"`
}

type send struct {
	*signal
	Duration int64 `json:"duration,omitempty"`
}

func (s *send) GetDuration() time.Duration {
	return time.Millisecond * time.Duration(s.Duration)
}

type status struct {
	Status       string `json:"status"`
	Message      string `json:"message"`
	Version      string `json:"version,omitempty"`
	LIRCDVersion string `json:"lircd_version,omitempty"`
}

type response struct {
	code    int
	Status  string   `json:"status"`
	Message string   `json:"message,omitempty"`
	Signals []signal `json:"signals,omitempty"`
}

func (res *response) Error() string {
	return res.Message
}
