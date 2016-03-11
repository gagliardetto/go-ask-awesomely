package typeform

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

//
func (client *Client) BaseInfo() (BaseInfo, error) {

	path := fmt.Sprintf("/%v/", client.APIVersion)
	method := "GET"

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return BaseInfo{}, err
	}

	fmt.Println(string(response))

	var baseInfo BaseInfo
	err = json.Unmarshal(response, &baseInfo)
	if err != nil {
		return BaseInfo{}, err
	}
	return baseInfo, nil
}

func (client *Client) Do(ff Form) (Response, error) {
	path := fmt.Sprintf("/%v/forms", client.APIVersion)
	method := "POST"

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}
	bodyPayload = ff

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return Response{}, err
	}

	fmt.Println(string(response))

	var formInfo Response
	err = json.Unmarshal(response, &formInfo)
	if err != nil {
		return Response{}, err
	}
	return formInfo, nil
}

func (client *Client) GetForm(formID string) (Response, error) {

	path := fmt.Sprintf("/%v/form/%v", client.APIVersion, formID)
	method := "GET"

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return Response{}, err
	}

	fmt.Println(string(response))

	var formInfo Response
	err = json.Unmarshal(response, &formInfo)
	if err != nil {
		return Response{}, err
	}
	return formInfo, nil
}
