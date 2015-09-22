package main

import "time"

type Contact struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Age        int       `json:"age"`
	ReqAgeFrom int       `json:"req_age_from"`
	ReqAgeTo   int       `json:"req_age_to"`
	Tags       []string  `json:"tags"`
	Sex        string    `json:"sex"`
	CreatedAt  time.Time `json:"created_at"`
}
