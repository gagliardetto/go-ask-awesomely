package typeform

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

// APIDomain is the domain of the typeform API
var APIDomain = "https://api.typeform.io/"

// NewClient creates a new API client
func NewClient(APIVersion APIVersion) (*Client, error) {
	return &Client{
		httpClient: http.DefaultClient,
		apiVersion: APIVersion,
		mu:         &sync.RWMutex{},
	}, nil
}

// SetAPIToken sets the API token used for making the requests to the API
func (client *Client) SetAPIToken(token string) error {
	if token == "" {
		return errors.New("token is empty")
	}
	client.mu.Lock()
	defer client.mu.Unlock()

	client.config.APIKey = token

	return nil
}

func (client *Client) fetchAndReturnPage(path string, method string, headers http.Header, queryParameters url.Values, bodyPayload interface{}) ([]byte, http.Header, error) {

	if client.config.APIKey == "" {
		return []byte(""), http.Header{}, fmt.Errorf("%s", "APIKey not provided")
	}

	requestURL, err := url.Parse(APIDomain)
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
	request.Header.Add("Accept-Encoding", "gzip")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "github.com/gagliardetto/go-ask-awesomely")
	request.Header.Add("X-API-TOKEN", client.config.APIKey)

	response, err := client.httpClient.Do(request)
	if err != nil {
		return []byte(""), http.Header{}, fmt.Errorf("Failed to get the URL %s: %s", requestURL, err)
	}
	defer response.Body.Close()

	var responseReader io.ReadCloser
	if strings.Contains(response.Header.Get("Content-Encoding"), "gzip") {
		decompressedBodyReader, err := gzip.NewReader(response.Body)
		if err != nil {
			return []byte(""), http.Header{}, err
		}
		responseReader = decompressedBodyReader
		defer responseReader.Close()
	} else {
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
