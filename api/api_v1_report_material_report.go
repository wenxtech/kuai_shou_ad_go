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


// V1ReportMaterialReportAPIService V1ReportMaterialReportAPI service
type V1ReportMaterialReportAPIService service

type ApiV1ReportMaterialReportPostRequest struct {
	ctx context.Context
	ApiService *V1ReportMaterialReportAPIService
	v1ReportMaterialReportPostRequest *V1ReportMaterialReportPostRequest
}

func (r *ApiV1ReportMaterialReportPostRequest) V1ReportMaterialReportPostRequest(v1ReportMaterialReportPostRequest V1ReportMaterialReportPostRequest) *ApiV1ReportMaterialReportPostRequest {
	r.v1ReportMaterialReportPostRequest = &v1ReportMaterialReportPostRequest
	return r
}

func (r *ApiV1ReportMaterialReportPostRequest) Execute() (*V1ReportMaterialReportResponse, *http.Response, error) {
	return r.ApiService.buildExecute(r)
}

func (r *ApiV1ReportMaterialReportPostRequest) WithLog(enable bool) *ApiV1ReportMaterialReportPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, ContextEnableLog, true)
	}
	return r
}

/*
V1ReportMaterialReportPost 素材报表

### 返回参数

说明：返回字段和请求参数中的 view_type 对应。 view_type 为视频报表时返回 photoUrl，coverUrl，photoCaption(作品广告语) 为封面报表时返回 coverUrl 为贴纸报表时返回 stickyUrl

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1ReportMaterialReportPostRequest
*/
func (a *V1ReportMaterialReportAPIService) Post(ctx context.Context) *ApiV1ReportMaterialReportPostRequest {
	return &ApiV1ReportMaterialReportPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V1ReportMaterialReportResponse
func (a *V1ReportMaterialReportAPIService) buildExecute(r *ApiV1ReportMaterialReportPostRequest) (*V1ReportMaterialReportResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []FormFile
		localVarReturnValue  *V1ReportMaterialReportResponse
	)
	
	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath, err := a.client.Cfg.ServerURLWithContext(r.ctx, "V1ReportMaterialReportAPIService.V1ReportMaterialReportPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/report/material_report"

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
	localVarPostBody = r.v1ReportMaterialReportPostRequest
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
