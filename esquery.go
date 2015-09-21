package main

import (
	"strconv"
)

type EsClient struct {
	host     string
	port     int
	endpoint string
}

func NewEsClient(host string, port int) *EsClient {
	endpoint := "http://" + host + ":" + strconv.Itoa(port)
	client := &EsClient{host, port, endpoint}
	return client
}

func (c EsClient) IntroCreate(intro Introduction) int {
	return 200
}

func (c EsClient) IntroResult(introId string) int {
	return 200

}

func (c EsClient) IntroDelete(introId string) int {
	return 200

}

func (c EsClient) ContactCreate(contract Contact) int {
	return 200

}

func (c EsClient) ContactDelete(contactId string) int {
	return 200

}

func (c EsClient) ContactShow(contactId string) int {
	return 200

}
