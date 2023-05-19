package api

import (
	"encoding/json"
	"fmt"
	"github.com/visionular/aurora-sdk-go/model"
	"io/ioutil"
	"net/url"
)

type VideoApiService service

func (a *VideoApiService) QueryTemplate(templateName string, opts ...APIOption) (model.TemplateApiResponse, error) {

	resValue := model.TemplateApiResponse{}
	localVarAPIOptions := new(APIOptions)
	localVarPath := a.client.cfg.basePath + "/vodencoding/v1/query_template?template_name=" + templateName

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	localVarHttpContentTypes := []string{"application/json"}
	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	r, err := a.client.prepareRequest(localVarAPIOptions, localVarPath, "GET", nil, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return resValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return resValue, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return resValue, err
	}

	err = CheckForHttpError(localVarHttpResponse.StatusCode, localVarBody)
	if err != nil {
		return resValue, err
	}
	err = json.Unmarshal(localVarBody, &resValue)
	if err != nil {
		return resValue, err
	}

	if resValue.Code != 0 {
		return resValue, fmt.Errorf(resValue.Msg)
	}

	return resValue, nil
}

func (a *VideoApiService) ListTemplate(opts ...APIOption) (model.TemplateListApiResponse, error) {

	resValue := model.TemplateListApiResponse{}
	localVarAPIOptions := new(APIOptions)
	localVarPath := a.client.cfg.basePath + "/vodencoding/v1/list_template"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	localVarHttpContentTypes := []string{"application/json"}
	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	r, err := a.client.prepareRequest(localVarAPIOptions, localVarPath, "GET", nil, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return resValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return resValue, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return resValue, err
	}

	err = CheckForHttpError(localVarHttpResponse.StatusCode, localVarBody)
	if err != nil {
		return resValue, err
	}
	err = json.Unmarshal(localVarBody, &resValue)
	if err != nil {
		return resValue, err
	}

	if resValue.Code != 0 {
		return resValue, fmt.Errorf(resValue.Msg)
	}

	return resValue, nil
}

func (a *VideoApiService) AddStorage(request model.StorageApiRequest, opts ...APIOption) (model.TemplateResponse, error) {

	resValue := model.TemplateResponse{}
	localVarAPIOptions := new(APIOptions)
	localVarPath := a.client.cfg.basePath + "/vodencoding/v1/add_storage"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	localVarHttpContentTypes := []string{"application/json"}
	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	r, err := a.client.prepareRequest(localVarAPIOptions, localVarPath, "POST", request, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return resValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return resValue, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return resValue, err
	}

	err = CheckForHttpError(localVarHttpResponse.StatusCode, localVarBody)
	if err != nil {
		return resValue, err
	}
	err = json.Unmarshal(localVarBody, &resValue)
	if err != nil {
		return resValue, err
	}

	if resValue.Code != 0 {
		return resValue, fmt.Errorf(resValue.Msg)
	}

	return resValue, nil
}

func (a *VideoApiService) ListStorage(opts ...APIOption) (model.StorageListApiResponse, error) {

	resValue := model.StorageListApiResponse{}
	localVarAPIOptions := new(APIOptions)
	localVarPath := a.client.cfg.basePath + "/vodencoding/v1/list_storage"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	localVarHttpContentTypes := []string{"application/json"}
	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	r, err := a.client.prepareRequest(localVarAPIOptions, localVarPath, "GET", nil, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return resValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return resValue, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return resValue, err
	}

	err = CheckForHttpError(localVarHttpResponse.StatusCode, localVarBody)
	if err != nil {
		return resValue, err
	}
	err = json.Unmarshal(localVarBody, &resValue)
	if err != nil {
		return resValue, err
	}

	if resValue.Code != 0 {
		return resValue, fmt.Errorf(resValue.Msg)
	}

	return resValue, nil
}

func (a *VideoApiService) DeleteStorage(storageID string, opts ...APIOption) (model.StorageDelResponse, error) {

	resValue := model.StorageDelResponse{}
	localVarAPIOptions := new(APIOptions)
	localVarPath := a.client.cfg.basePath + "/vodencoding/v1/del_storage?storage_id=" + storageID

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	localVarHttpContentTypes := []string{"application/json"}
	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	r, err := a.client.prepareRequest(localVarAPIOptions, localVarPath, "DELETE", nil, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return resValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return resValue, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return resValue, err
	}

	err = CheckForHttpError(localVarHttpResponse.StatusCode, localVarBody)
	if err != nil {
		return resValue, err
	}
	err = json.Unmarshal(localVarBody, &resValue)
	if err != nil {
		return resValue, err
	}

	if resValue.Code != 0 {
		return resValue, fmt.Errorf(resValue.Msg)
	}

	return resValue, nil
}
