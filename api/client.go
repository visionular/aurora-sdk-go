package api

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")
)

type APIClient struct {
	cfg        *Configuration
	httpClient *http.Client
	common     service

	LiveStreamsApi *LiveStreamsApiService
	VideoViewsApi  *VideoApiService
}

type service struct {
	client *APIClient
}

func NewAPIClient(cfg *Configuration) *APIClient {
	c := &APIClient{}
	c.cfg = cfg
	c.httpClient = &http.Client{
		Timeout: cfg.timeout,
	}
	c.common.client = c

	c.LiveStreamsApi = (*LiveStreamsApiService)(&c.common)
	c.VideoViewsApi = (*VideoApiService)(&c.common)

	return c
}

func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	return c.httpClient.Do(request)
}

type APIOption func(*APIOptions)

type APIOptions struct {
	ctx    context.Context
	params interface{}
}

func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

func (c *APIClient) prepareRequest(
	opts *APIOptions,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values) (requestInfo *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	url.RawQuery = query.Encode()

	if body != nil {
		requestInfo, err = http.NewRequest(method, url.String(), body)
	} else {
		requestInfo, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		requestInfo.Header = headers
	}

	if c.cfg.host != "" {
		requestInfo.Host = c.cfg.host
	}

	requestInfo.Header.Add("auth-type", "use-basic")
	requestInfo.Header.Add("User-Agent", c.cfg.userAgent)

	if opts.ctx != nil {
		requestInfo = requestInfo.WithContext(opts.ctx)
	}
	// BasicAuth
	if c.cfg.username == "" || c.cfg.password == "" {
		return nil, errors.New("unauthorized APIClient")
	}
	requestInfo.SetBasicAuth(c.cfg.username, c.cfg.password)

	return requestInfo, nil
}

func CheckForHttpError(code int, body []byte) error {

	if code >= 200 && code <= 299 {
		return nil
	}

	switch code {
	case 400:
		return fmt.Errorf("Bad Request")
	case 401:
		return UnauthorizedError{body: body, error: "Unauthorized"}
	case 404:
		return fmt.Errorf("Not Found")
	}

	return fmt.Errorf("Service Error")
}

type BadRequestError struct {
	body  []byte
	error string
}

func (e BadRequestError) Error() string {
	return e.error
}

// 401 Error
type UnauthorizedError struct {
	body  []byte
	error string
}

func (e UnauthorizedError) Error() string {
	return e.error
}

func (e UnauthorizedError) Body() []byte {
	return e.body
}

// 403 Error
type ForbiddenError struct {
	body  []byte
	error string
}

func (e ForbiddenError) Error() string {
	return e.error
}

func (e ForbiddenError) Body() []byte {
	return e.body
}

// 404 Error
type NotFoundError struct {
	body  []byte
	error string
}

func (e NotFoundError) Error() string {
	return e.error
}

func (e NotFoundError) Body() []byte {
	return e.body
}

// 429 Error
type TooManyRequestsError struct {
	body  []byte
	error string
}

func (e TooManyRequestsError) Error() string {
	return e.error
}

func (e TooManyRequestsError) Body() []byte {
	return e.body
}

// 5XX Error
type ServiceError struct {
	body   []byte
	status int
	error  string
}
