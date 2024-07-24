/*
磁力引擎 copy

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
    . "github.com/wenxtech/kuai_shou_ad_go/models"
)


// V1AsyncTaskListAPIService V1AsyncTaskListAPI service
type V1AsyncTaskListAPIService service

type ApiV1AsyncTaskListPostRequest struct {
	ctx context.Context
	ApiService *V1AsyncTaskListAPIService
	v1AsyncTaskListPostRequest *V1AsyncTaskListPostRequest
}

func (r *ApiV1AsyncTaskListPostRequest) V1AsyncTaskListPostRequest(v1AsyncTaskListPostRequest V1AsyncTaskListPostRequest) *ApiV1AsyncTaskListPostRequest {
	r.v1AsyncTaskListPostRequest = &v1AsyncTaskListPostRequest
	return r
}

func (r *ApiV1AsyncTaskListPostRequest) Execute() (*V1AsyncTaskListResponse, *http.Response, error) {
	return r.ApiService.buildExecute(r)
}

func (r *ApiV1AsyncTaskListPostRequest) WithLog(enable bool) *ApiV1AsyncTaskListPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, ContextEnableLog, true)
	}
	return r
}

/*
V1AsyncTaskListPost 获取任务状态

### 请求参数

| 字段          | 类型   | 是否必填 | 描述                                                      |
| ------------- | ------ | -------- | --------------------------------------------------------- |
| advertiser_id | long   | 必填     | 广告主 ID（注：非账户快手 ID），在获取 accessToken 时返回 |
| task_ids      | long[] | 可选     | 任务 ID 集，不超过 10 个                                  |
| page          | int    | 可选     | 页码，默认 1                                              |
| page_size     | int    | 可选     | 每页行数，默认 20                                         |

### 返回参数

| 字段          | 类型     | 说明         | 备注                                         |
| ------------- | -------- | ------------ | -------------------------------------------- |
| code          | int      | 返回码       |                                              |
| message       | string   | 返回信息     |                                              |
| data          | struct   | JSON 返回值  |                                              |
| total_count   | int      | 任务总数     |                                              |
| details       | struct[] |              |                                              |
| advertiser_id | long     | 广告主 ID    |                                              |
| task_id       | long     | 任务 ID      |                                              |
| task_name     | string   | 任务名称     |                                              |
| create_time   | string   | 任务创建时间 | 格式如：yyyy-MM-dd HH:mm:ss                  |
| task_status   | int      | 任务状态     | 0：新建，1：处理中，2：处理成功，3：处理失败 |
| file_size     | int      | 文件大小     | 字节数                                       |



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1AsyncTaskListPostRequest
*/
func (a *V1AsyncTaskListAPIService) Post(ctx context.Context) *ApiV1AsyncTaskListPostRequest {
	return &ApiV1AsyncTaskListPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V1AsyncTaskListResponse
func (a *V1AsyncTaskListAPIService) buildExecute(r *ApiV1AsyncTaskListPostRequest) (*V1AsyncTaskListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []FormFile
		localVarReturnValue  *V1AsyncTaskListResponse
	)
	
	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath, err := a.client.Cfg.ServerURLWithContext(r.ctx, "V1AsyncTaskListAPIService.V1AsyncTaskListPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/async_task/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.v1AsyncTaskListPostRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
