package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"

	"github.com/visionular/aurora-sdk-go/model"
)

type VideoApiService service

func (a *VideoApiService) doRequest(
	opts []APIOption,
	path, method string,
	queryParams url.Values,
	body interface{},
	response interface{},
) error {
	// apply options
	localOpts := new(APIOptions)
	for _, opt := range opts {
		opt(localOpts)
	}

	// headers
	headers := make(map[string]string)
	if ct := selectHeaderContentType([]string{"application/json"}); ct != "" {
		headers["Content-Type"] = ct
	}

	// prepare and execute
	req, err := a.client.prepareRequest(localOpts, a.client.cfg.basePath+path, method, body, headers, queryParams)
	if err != nil {
		return err
	}

	httpResp, err := a.client.callAPI(req)
	if err != nil || httpResp == nil {
		return err
	}
	defer httpResp.Body.Close()

	// read and check
	data, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}
	if err := CheckForHttpError(httpResp.StatusCode, data); err != nil {
		return err
	}

	// unmarshal
	if err := json.Unmarshal(data, response); err != nil {
		return err
	}

	return nil
}

func (a *VideoApiService) QueryTemplate(templateName string, opts ...APIOption) (model.TemplateApiResponse, error) {
	var resp model.TemplateApiResponse
	path := "/vodencoding/v1/query_template"
	qp := url.Values{}
	qp.Set("template_name", templateName)

	err := a.doRequest(opts, path, "GET", qp, nil, &resp)
	if err != nil {
		return resp, err
	}
	if resp.Code != 0 {
		return resp, fmt.Errorf(resp.Msg)
	}

	return resp, nil
}

func (a *VideoApiService) ListTemplate(opts ...APIOption) (model.TemplateListApiResponse, error) {
	var resp model.TemplateListApiResponse
	err := a.doRequest(opts, "/vodencoding/v1/list_template", "GET", url.Values{}, nil, &resp)
	if err != nil {
		return resp, err
	}
	if resp.Code != 0 {
		return resp, fmt.Errorf(resp.Msg)
	}

	return resp, nil
}

func (a *VideoApiService) AddStorage(request model.StorageApiRequest, opts ...APIOption) (model.StorageResponse, error) {
	var resp model.StorageResponse
	err := a.doRequest(opts, "/vodencoding/v1/add_storage", "POST", url.Values{}, request, &resp)
	if err != nil {
		return resp, err
	}
	if resp.Code != 0 {
		return resp, fmt.Errorf(resp.Msg)
	}

	return resp, nil
}

func (a *VideoApiService) ListStorage(opts ...APIOption) (model.StorageListApiResponse, error) {
	var resp model.StorageListApiResponse
	err := a.doRequest(opts, "/vodencoding/v1/list_storage", "GET", url.Values{}, nil, &resp)
	if err != nil {
		return resp, err
	}
	if resp.Code != 0 {
		return resp, fmt.Errorf(resp.Msg)
	}

	return resp, nil
}

func (a *VideoApiService) DeleteStorage(storageID string, opts ...APIOption) (model.StorageDelResponse, error) {
	var resp model.StorageDelResponse
	qp := url.Values{}
	qp.Set("storage_id", storageID)
	err := a.doRequest(opts, "/vodencoding/v1/del_storage", "DELETE", qp, nil, &resp)
	if err != nil {
		return resp, err
	}
	if resp.Code != 0 {
		return resp, fmt.Errorf(resp.Msg)
	}

	return resp, nil
}

func (a *VideoApiService) CreateTask(request model.TaskApiRequest, opts ...APIOption) (model.TaskResponse, error) {
	var resp model.TaskResponse
	err := a.doRequest(opts, "/vodencoding/v1/create_task", "POST", url.Values{}, request, &resp)
	if err != nil {
		return resp, err
	}
	if resp.Code != 0 {
		return resp, fmt.Errorf(resp.Msg)
	}

	return resp, nil
}

func (a *VideoApiService) QueryTask(taskID string, opts ...APIOption) (model.TaskApiResponse, error) {
	var resp model.TaskApiResponse
	path := "/vodencoding/v1/query_task"
	qp := url.Values{}
	qp.Set("task_id", taskID)

	err := a.doRequest(opts, path, "GET", qp, nil, &resp)
	if err != nil {
		return resp, err
	}
	if resp.Code != 0 {
		return resp, fmt.Errorf(resp.Msg)
	}

	return resp, nil
}

func (a *VideoApiService) ListTask(request model.TaskListApiRequest, opts ...APIOption) (model.TaskListApiResponse, error) {
	var resp model.TaskListApiResponse
	path := "/vodencoding/v1/list_task"
	qp := url.Values{}
	qp.Set("count", strconv.Itoa(request.Count))
	qp.Set("start_num", strconv.Itoa(request.StartNum))
	qp.Set("start_time", strconv.FormatInt(request.StartTime, 10))
	qp.Set("end_time", strconv.FormatInt(request.EndTime, 10))
	qp.Set("status", request.Status)

	err := a.doRequest(opts, path, "GET", qp, nil, &resp)
	if err != nil {
		return resp, err
	}
	if resp.Code != 0 {
		return resp, fmt.Errorf(resp.Msg)
	}

	return resp, nil
}
