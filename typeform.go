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

	var baseInfo BaseInfo
	err = json.Unmarshal(response, &baseInfo)
	if err != nil {
		return BaseInfo{}, err
	}
	return baseInfo, nil
}

func (client *Client) CreateForm(newForm Form) (FormInfo, error) {
	path := fmt.Sprintf("/%v/forms", client.APIVersion)
	method := "POST"

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}
	bodyPayload = newForm

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return FormInfo{}, err
	}

	var formInfo FormInfo
	err = json.Unmarshal(response, &formInfo)
	if err != nil {
		return FormInfo{}, err
	}
	return formInfo, nil
}

func (client *Client) GetForm(formID string) (FormInfo, error) {

	path := fmt.Sprintf("/%v/forms/%v", client.APIVersion, formID)
	method := "GET"

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return FormInfo{}, err
	}

	var formInfo FormInfo
	err = json.Unmarshal(response, &formInfo)
	if err != nil {
		return FormInfo{}, err
	}
	return formInfo, nil
}

func (client *Client) CreateImage(imageURL string) (NewImage, error) {

	path := fmt.Sprintf("/%v/images", client.APIVersion)
	method := "POST"

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}

	var newImage struct {
		URL string `json:"url"`
	}
	newImage.URL = imageURL
	bodyPayload = newImage

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return NewImage{}, err
	}

	var newImageResponse NewImage
	err = json.Unmarshal(response, &newImageResponse)
	if err != nil {
		return NewImage{}, err
	}
	return newImageResponse, nil
}

func (client *Client) GetImage(imageID string) (ImageInfo, error) {

	path := fmt.Sprintf("/%v/images/%v", client.APIVersion, imageID)
	method := "GET"

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return ImageInfo{}, err
	}

	var imageInfo ImageInfo
	err = json.Unmarshal(response, &imageInfo)
	if err != nil {
		return ImageInfo{}, err
	}
	return imageInfo, nil
}

func (client *Client) CreateDesign(newDesign Design) (DesignInfo, error) {
	path := fmt.Sprintf("/%v/designs", client.APIVersion)
	method := "POST"

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}
	bodyPayload = newDesign

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return DesignInfo{}, err
	}

	var designInfo DesignInfo
	err = json.Unmarshal(response, &designInfo)
	if err != nil {
		return DesignInfo{}, err
	}
	return designInfo, nil
}

func (client *Client) GetDesign(designID string) (DesignInfo, error) {

	path := fmt.Sprintf("/%v/designs/%v", client.APIVersion, designID)
	method := "GET"

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return DesignInfo{}, err
	}

	var designInfo DesignInfo
	err = json.Unmarshal(response, &designInfo)
	if err != nil {
		return DesignInfo{}, err
	}
	return designInfo, nil
}

func (client *Client) CreateURL(formID string) (URLInfo, error) {

	path := fmt.Sprintf("/%v/urls", client.APIVersion)
	method := "POST"

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}

	var newURL struct {
		FormID string `json:"form_id"`
	}
	newURL.FormID = formID
	bodyPayload = newURL

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return URLInfo{}, err
	}

	var newURLResponse URLInfo
	err = json.Unmarshal(response, &newURLResponse)
	if err != nil {
		return URLInfo{}, err
	}
	return newURLResponse, nil
}

func (client *Client) GetURL(URLID string) (URLInfo, error) {

	path := fmt.Sprintf("/%v/urls/%v", client.APIVersion, URLID)
	method := "GET"

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return URLInfo{}, err
	}

	var URLInfoResponse URLInfo
	err = json.Unmarshal(response, &URLInfoResponse)
	if err != nil {
		return URLInfo{}, err
	}
	return URLInfoResponse, nil
}

func (client *Client) ModifyURL(URLID string) (URLInfo, error) {

	path := fmt.Sprintf("/%v/urls/%v", client.APIVersion, URLID)
	method := "PUT"

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}

	var newURL struct {
		FormID string `json:"form_id"` // is it form_id or url_id ????
	}
	newURL.FormID = URLID
	bodyPayload = newURL

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return URLInfo{}, err
	}

	var newURLResponse URLInfo
	err = json.Unmarshal(response, &newURLResponse)
	if err != nil {
		return URLInfo{}, err
	}
	return newURLResponse, nil
}
