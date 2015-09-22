package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
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

func (c *EsClient) InsertIndex(indexStr string, typeStr string, idStr string, jsonBody []byte) ([]byte, error) {

	url := (*c).endpoint + "/" + indexStr + "/" + typeStr + "/" + idStr
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return body, err
}

func (c *EsClient) IntroCreate(intro Introduction) int {
	return 200
}

func (c *EsClient) IntroResult(introId string) int {
	return 200

}

func (c *EsClient) IntroDelete(introId string) int {
	return 200

}

func (c *EsClient) ContactCreate(contract Contact) int {
	return 200

}

func (c *EsClient) ContactDelete(contactId string) int {
	return 200

}

func (c *EsClient) ContactShow(contactId string) int {
	return 200
}
