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


// V1ReportCampaignReportAPIService V1ReportCampaignReportAPI service
type V1ReportCampaignReportAPIService service

type ApiV1ReportCampaignReportPostRequest struct {
	ctx context.Context
	ApiService *V1ReportCampaignReportAPIService
	v1ReportCampaignReportPostRequest *V1ReportCampaignReportPostRequest
}

func (r *ApiV1ReportCampaignReportPostRequest) V1ReportCampaignReportPostRequest(v1ReportCampaignReportPostRequest V1ReportCampaignReportPostRequest) *ApiV1ReportCampaignReportPostRequest {
	r.v1ReportCampaignReportPostRequest = &v1ReportCampaignReportPostRequest
	return r
}

func (r *ApiV1ReportCampaignReportPostRequest) Execute() (*V1ReportCampaignReportResponse, *http.Response, error) {
	return r.ApiService.buildExecute(r)
}

func (r *ApiV1ReportCampaignReportPostRequest) WithLog(enable bool) *ApiV1ReportCampaignReportPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, ContextEnableLog, true)
	}
	return r
}

/*
V1ReportCampaignReportPost 广告计划数据



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1ReportCampaignReportPostRequest
*/
func (a *V1ReportCampaignReportAPIService) Post(ctx context.Context) *ApiV1ReportCampaignReportPostRequest {
	return &ApiV1ReportCampaignReportPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V1ReportCampaignReportResponse
func (a *V1ReportCampaignReportAPIService) buildExecute(r *ApiV1ReportCampaignReportPostRequest) (*V1ReportCampaignReportResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []FormFile
		localVarReturnValue  *V1ReportCampaignReportResponse
	)
	
	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath, err := a.client.Cfg.ServerURLWithContext(r.ctx, "V1ReportCampaignReportAPIService.V1ReportCampaignReportPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/report/campaign_report"

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
	localVarPostBody = r.v1ReportCampaignReportPostRequest
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
