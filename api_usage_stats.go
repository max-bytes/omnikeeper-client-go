/*
Landscape omnikeeper REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package okclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)


// UsageStatsAPIService UsageStatsAPI service
type UsageStatsAPIService service

type ApiUsageStatsFetchRequest struct {
	ctx context.Context
	ApiService *UsageStatsAPIService
	from *time.Time
	to *time.Time
	version string
}

func (r ApiUsageStatsFetchRequest) From(from time.Time) ApiUsageStatsFetchRequest {
	r.from = &from
	return r
}

func (r ApiUsageStatsFetchRequest) To(to time.Time) ApiUsageStatsFetchRequest {
	r.to = &to
	return r
}

func (r ApiUsageStatsFetchRequest) Execute() (*http.Response, error) {
	return r.ApiService.UsageStatsFetchExecute(r)
}

/*
UsageStatsFetch Method for UsageStatsFetch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param version
 @return ApiUsageStatsFetchRequest
*/
func (a *UsageStatsAPIService) UsageStatsFetch(ctx context.Context, version string) ApiUsageStatsFetchRequest {
	return ApiUsageStatsFetchRequest{
		ApiService: a,
		ctx: ctx,
		version: version,
	}
}

// Execute executes the request
func (a *UsageStatsAPIService) UsageStatsFetchExecute(r ApiUsageStatsFetchRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageStatsAPIService.UsageStatsFetch")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v{version}/UsageStats/fetch"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", url.PathEscape(parameterValueToString(r.version, "version")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.from == nil {
		return nil, reportError("from is required and must be specified")
	}
	if r.to == nil {
		return nil, reportError("to is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "to", r.to, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
