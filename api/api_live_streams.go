package api

import (
	"encoding/json"
	"fmt"
	"github.com/visionular/aurora-sdk-go/model"
	"io/ioutil"
	"net/url"
	"strings"
)

type LiveStreamsApiService service

func (a *LiveStreamsApiService) CreateLiveStream(request model.CreateStreamRequest, opts ...APIOption) (model.LiveStreamResponse, error) {

	resValue := model.LiveStreamResponse{}
	localVarAPIOptions := new(APIOptions)
	localVarPath := a.client.cfg.basePath + "/live/v1/live-streams"

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

func (a *LiveStreamsApiService) DisableLiveStream(LiveStreamID string, opts ...APIOption) (model.DisableLiveStreamResponse, error) {

	resValue := model.DisableLiveStreamResponse{}
	localVarAPIOptions := new(APIOptions)
	localVarPath := a.client.cfg.basePath + "/live/v1/live-streams/{LIVE_STREAM_ID}/disable"
	localVarPath = strings.Replace(localVarPath, "{"+"LIVE_STREAM_ID"+"}", fmt.Sprintf("%v", LiveStreamID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	r, err := a.client.prepareRequest(localVarAPIOptions, localVarPath, "PUT", nil, localVarHeaderParams, localVarQueryParams)
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

func (a *LiveStreamsApiService) EnableLiveStream(LiveStreamID string, opts ...APIOption) (model.EnableLiveStreamResponse, error) {

	resValue := model.EnableLiveStreamResponse{}
	localVarAPIOptions := new(APIOptions)
	localVarPath := a.client.cfg.basePath + "/live/v1/live-streams/{LIVE_STREAM_ID}/enable"
	localVarPath = strings.Replace(localVarPath, "{"+"LIVE_STREAM_ID"+"}", fmt.Sprintf("%v", LiveStreamID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

	r, err := a.client.prepareRequest(localVarAPIOptions, localVarPath, "PUT", nil, localVarHeaderParams, localVarQueryParams)
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

func (a *LiveStreamsApiService) ListLiveStreams(opts ...APIOption) (model.ListLiveStreamsResponse, error) {

	resValue := model.ListLiveStreamsResponse{}
	localVarAPIOptions := new(APIOptions)
	localVarPath := a.client.cfg.basePath + "/live/v1/live-streams"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

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

func (a *LiveStreamsApiService) QueryLiveStreamInfo(LiveStreamID string, opts ...APIOption) (model.LiveStreamResponse, error) {

	resValue := model.LiveStreamResponse{}
	localVarAPIOptions := new(APIOptions)
	localVarPath := a.client.cfg.basePath + "/live/v1/live-streams/{LIVE_STREAM_ID}"
	localVarPath = strings.Replace(localVarPath, "{"+"LIVE_STREAM_ID"+"}", fmt.Sprintf("%v", LiveStreamID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}

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
