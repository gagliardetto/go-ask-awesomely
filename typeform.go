package typeform

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// BaseInfo handles the endpoint used to get info about the API
func (client *Client) BaseInfo() (*BaseInfo, error) {

	path := fmt.Sprintf("/%v/", client.apiVersion)
	method := http.MethodGet

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return nil, err
	}

	var baseInfo BaseInfo
	err = json.Unmarshal(response, &baseInfo)
	if err != nil {
		return nil, err
	}
	return &baseInfo, nil
}

// CreateForm handles the endpoint used to get create a new form from the provided model
func (client *Client) CreateForm(newForm Form) (*FormInfo, error) {
	path := fmt.Sprintf("/%v/forms", client.apiVersion)
	method := http.MethodPost

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}
	bodyPayload = newForm

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return nil, err
	}

	var formInfo FormInfo
	err = json.Unmarshal(response, &formInfo)
	if err != nil {
		return nil, err
	}
	return &formInfo, nil
}

// GetForm handles the endpoint used to fetch a form by ID
func (client *Client) GetForm(formID string) (*FormInfo, error) {

	path := fmt.Sprintf("/%v/forms/%v", client.apiVersion, formID)
	method := http.MethodGet

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return nil, err
	}

	var formInfo FormInfo
	err = json.Unmarshal(response, &formInfo)
	if err != nil {
		return nil, err
	}
	return &formInfo, nil
}

// CreateImage handles the endpoint used to upload an image that
// will be then available for use in the forms as "Picture Choices"
func (client *Client) CreateImage(imageURL string) (*NewImage, error) {

	path := fmt.Sprintf("/%v/images", client.apiVersion)
	method := http.MethodPost

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
		return nil, err
	}

	var newImageResponse NewImage
	err = json.Unmarshal(response, &newImageResponse)
	if err != nil {
		return nil, err
	}
	return &newImageResponse, nil
}

// GetImage handles the endpoint used to get an image by ID
func (client *Client) GetImage(imageID string) (*ImageInfo, error) {

	path := fmt.Sprintf("/%v/images/%v", client.apiVersion, imageID)
	method := http.MethodGet

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return nil, err
	}

	var imageInfo ImageInfo
	err = json.Unmarshal(response, &imageInfo)
	if err != nil {
		return nil, err
	}
	return &imageInfo, nil
}

// CreateDesign handles the endpoint used to create a design that will
// be available to be used to style forms
func (client *Client) CreateDesign(newDesign Design) (*DesignInfo, error) {
	path := fmt.Sprintf("/%v/designs", client.apiVersion)
	method := http.MethodPost

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}
	bodyPayload = newDesign

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return nil, err
	}

	var designInfo DesignInfo
	err = json.Unmarshal(response, &designInfo)
	if err != nil {
		return nil, err
	}
	return &designInfo, nil
}

// GetDesign handles the endpoint used to get a design
func (client *Client) GetDesign(designID string) (*DesignInfo, error) {

	path := fmt.Sprintf("/%v/designs/%v", client.apiVersion, designID)
	method := http.MethodGet

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return nil, err
	}

	var designInfo DesignInfo
	err = json.Unmarshal(response, &designInfo)
	if err != nil {
		return nil, err
	}
	return &designInfo, nil
}

// CreateURL handles the endpoint used to create a new URL linking to a typeform
func (client *Client) CreateURL(formID string) (*URLInfo, error) {

	path := fmt.Sprintf("/%v/urls", client.apiVersion)
	method := http.MethodPost

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
		return nil, err
	}

	var newURLResponse URLInfo
	err = json.Unmarshal(response, &newURLResponse)
	if err != nil {
		return nil, err
	}
	return &newURLResponse, nil
}

// GetURL handles the endpoint used to get the typeform a URL links to
func (client *Client) GetURL(URLID string) (*URLInfo, error) {

	path := fmt.Sprintf("/%v/urls/%v", client.apiVersion, URLID)
	method := http.MethodGet

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return nil, err
	}

	var URLInfoResponse URLInfo
	err = json.Unmarshal(response, &URLInfoResponse)
	if err != nil {
		return nil, err
	}
	return &URLInfoResponse, nil
}

// ModifyURL handles the endpoint used to change an existing URL to link to a different typeform
func (client *Client) ModifyURL(URLID string, formID string) (*URLInfo, error) {

	path := fmt.Sprintf("/%v/urls/%v", client.apiVersion, URLID)
	method := http.MethodPut

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
		return nil, err
	}

	var newURLResponse URLInfo
	err = json.Unmarshal(response, &newURLResponse)
	if err != nil {
		return nil, err
	}
	return &newURLResponse, nil
}

// DeleteURL handles the endpoint used to delete a URL that links to a typeform
func (client *Client) DeleteURL(URLID string) error {

	path := fmt.Sprintf("/%v/urls/%v", client.apiVersion, URLID)
	method := http.MethodDelete

	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}

	_, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	if err != nil {
		return err
	}

	return nil
}
