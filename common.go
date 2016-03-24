package typeform

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// NewClient creates a new API client
func NewClient(APIVersion APIVersion) (*Client, error) {
	return &Client{
		httpClient: http.DefaultClient,
		APIVersion: APIVersion,
	}, nil
}

func (client *Client) fetchAndReturnPage(path string, method string, headers http.Header, queryParameters url.Values, bodyPayload interface{}) ([]byte, http.Header, error) {

	domain := "https://api.typeform.io/"
	requestURL, err := url.Parse(domain)
	if err != nil {
		return []byte(""), http.Header{}, err
	}
	requestURL.Path = path
	requestURL.RawQuery = queryParameters.Encode()

	if method != "GET" && method != "POST" && method != "PUT" && method != "PATCH" && method != "DELETE" {
		return []byte(""), http.Header{}, fmt.Errorf("Method not supported: %v", method)
	}

	encodedBody, err := json.Marshal(bodyPayload)
	if err != nil {
		return []byte(""), http.Header{}, err
	}

	//fmt.Println(requestURL.String())
	request, err := http.NewRequest(method, requestURL.String(), bytes.NewBuffer(encodedBody))
	if err != nil {
		return []byte(""), http.Header{}, fmt.Errorf("Failed to get the URL %s: %s", requestURL, err)
	}
	request.Header = headers
	request.Header.Add("Content-Length", strconv.Itoa(len(encodedBody)))

	request.Header.Add("Connection", "Keep-Alive")
	request.Header.Add("Accept-Encoding", "gzip, deflate")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Set("User-Agent", "github.com/gagliardetto/go-ask-awesomely")
	request.Header.Add("X-API-TOKEN", client.Config.APIKey)

	response, err := client.httpClient.Do(request)
	if err != nil {
		return []byte(""), http.Header{}, fmt.Errorf("Failed to get the URL %s: %s", requestURL, err)
	}
	defer response.Body.Close()

	var responseReader io.ReadCloser
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		decompressedBodyReader, err := gzip.NewReader(response.Body)
		if err != nil {
			return []byte(""), http.Header{}, err
		}
		responseReader = decompressedBodyReader
		defer responseReader.Close()
	default:
		responseReader = response.Body
	}

	responseBody, err := ioutil.ReadAll(responseReader)
	if err != nil {
		return []byte(""), http.Header{}, err
	}

	if response.StatusCode > 299 || response.StatusCode < 199 {
		var apiError APIError
		err = json.Unmarshal(responseBody, &apiError)
		if err != nil {
			return []byte(""), http.Header{}, nil
		}
		//fmt.Println(string(responseBody))
		return []byte(""), http.Header{}, fmt.Errorf("HTTPStatus %s: %s", strconv.Itoa(response.StatusCode), apiError.String())
	}

	return responseBody, response.Header, nil
}

func (apiError *APIError) String() string {
	return fmt.Sprintf("Error: %q; Field: %q; Description: %q", apiError.Error, apiError.Field, apiError.Description)
}
