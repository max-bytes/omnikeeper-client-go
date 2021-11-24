/*
Landscape omnikeeper REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package okclient

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"reflect"
	"time"
)

// Linger please
var (
	_ _context.Context
)

// CIApiService CIApi service
type CIApiService service

type ApiGetAllCIIDsRequest struct {
	ctx _context.Context
	ApiService *CIApiService
	version string
}


func (r ApiGetAllCIIDsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.GetAllCIIDsExecute(r)
}

/*
GetAllCIIDs list of all CI-IDs

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param version
 @return ApiGetAllCIIDsRequest
*/
func (a *CIApiService) GetAllCIIDs(ctx _context.Context, version string) ApiGetAllCIIDsRequest {
	return ApiGetAllCIIDsRequest{
		ApiService: a,
		ctx: ctx,
		version: version,
	}
}

// Execute executes the request
//  @return []string
func (a *CIApiService) GetAllCIIDsExecute(r ApiGetAllCIIDsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CIApiService.GetAllCIIDs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v{version}/CI/getAllCIIDs"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", _neturl.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;odata.metadata=minimal;odata.streaming=true", "application/json;odata.metadata=minimal;odata.streaming=false", "application/json;odata.metadata=minimal", "application/json;odata.metadata=full;odata.streaming=true", "application/json;odata.metadata=full;odata.streaming=false", "application/json;odata.metadata=full", "application/json;odata.metadata=none;odata.streaming=true", "application/json;odata.metadata=none;odata.streaming=false", "application/json;odata.metadata=none", "application/json;odata.streaming=true", "application/json;odata.streaming=false", "application/json", "application/xml", "application/odata", "text/plain", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCIByIDRequest struct {
	ctx _context.Context
	ApiService *CIApiService
	layerIDs *[]string
	cIID *string
	version string
	atTime *time.Time
}

// Specifies which layers contribute to the result, and in which order
func (r ApiGetCIByIDRequest) LayerIDs(layerIDs []string) ApiGetCIByIDRequest {
	r.layerIDs = &layerIDs
	return r
}
func (r ApiGetCIByIDRequest) CIID(cIID string) ApiGetCIByIDRequest {
	r.cIID = &cIID
	return r
}
// Specify datetime, for which point in time to get the data; leave empty to use current time (https://www.newtonsoft.com/json/help/html/DatesInJSON.htm)
func (r ApiGetCIByIDRequest) AtTime(atTime time.Time) ApiGetCIByIDRequest {
	r.atTime = &atTime
	return r
}

func (r ApiGetCIByIDRequest) Execute() (CIDTO, *_nethttp.Response, error) {
	return r.ApiService.GetCIByIDExecute(r)
}

/*
GetCIByID single CI by CI-ID

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param version
 @return ApiGetCIByIDRequest
*/
func (a *CIApiService) GetCIByID(ctx _context.Context, version string) ApiGetCIByIDRequest {
	return ApiGetCIByIDRequest{
		ApiService: a,
		ctx: ctx,
		version: version,
	}
}

// Execute executes the request
//  @return CIDTO
func (a *CIApiService) GetCIByIDExecute(r ApiGetCIByIDRequest) (CIDTO, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  CIDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CIApiService.GetCIByID")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v{version}/CI/getCIByID"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", _neturl.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.layerIDs == nil {
		return localVarReturnValue, nil, reportError("layerIDs is required and must be specified")
	}
	if r.cIID == nil {
		return localVarReturnValue, nil, reportError("cIID is required and must be specified")
	}

	{
		t := *r.layerIDs
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("layerIDs", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("layerIDs", parameterToString(t, "multi"))
		}
	}
	localVarQueryParams.Add("CIID", parameterToString(*r.cIID, ""))
	if r.atTime != nil {
		localVarQueryParams.Add("atTime", parameterToString(*r.atTime, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;odata.metadata=minimal;odata.streaming=true", "application/json;odata.metadata=minimal;odata.streaming=false", "application/json;odata.metadata=minimal", "application/json;odata.metadata=full;odata.streaming=true", "application/json;odata.metadata=full;odata.streaming=false", "application/json;odata.metadata=full", "application/json;odata.metadata=none;odata.streaming=true", "application/json;odata.metadata=none;odata.streaming=false", "application/json;odata.metadata=none", "application/json;odata.streaming=true", "application/json;odata.streaming=false", "application/json", "application/xml", "application/odata", "text/plain", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCIsByIDRequest struct {
	ctx _context.Context
	ApiService *CIApiService
	layerIDs *[]string
	cIIDs *[]string
	version string
	atTime *time.Time
}

// Specifies which layers contribute to the result, and in which order
func (r ApiGetCIsByIDRequest) LayerIDs(layerIDs []string) ApiGetCIsByIDRequest {
	r.layerIDs = &layerIDs
	return r
}
func (r ApiGetCIsByIDRequest) CIIDs(cIIDs []string) ApiGetCIsByIDRequest {
	r.cIIDs = &cIIDs
	return r
}
// Specify datetime, for which point in time to get the data; leave empty to use current time (https://www.newtonsoft.com/json/help/html/DatesInJSON.htm)
func (r ApiGetCIsByIDRequest) AtTime(atTime time.Time) ApiGetCIsByIDRequest {
	r.atTime = &atTime
	return r
}

func (r ApiGetCIsByIDRequest) Execute() ([]CIDTO, *_nethttp.Response, error) {
	return r.ApiService.GetCIsByIDExecute(r)
}

/*
GetCIsByID multiple CIs by CI-ID  !Watch out for the query URL getting too long because of a lot of CIIDs!  TODO: consider using POST

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param version
 @return ApiGetCIsByIDRequest
*/
func (a *CIApiService) GetCIsByID(ctx _context.Context, version string) ApiGetCIsByIDRequest {
	return ApiGetCIsByIDRequest{
		ApiService: a,
		ctx: ctx,
		version: version,
	}
}

// Execute executes the request
//  @return []CIDTO
func (a *CIApiService) GetCIsByIDExecute(r ApiGetCIsByIDRequest) ([]CIDTO, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CIDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CIApiService.GetCIsByID")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v{version}/CI/getCIsByID"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", _neturl.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.layerIDs == nil {
		return localVarReturnValue, nil, reportError("layerIDs is required and must be specified")
	}
	if r.cIIDs == nil {
		return localVarReturnValue, nil, reportError("cIIDs is required and must be specified")
	}

	{
		t := *r.layerIDs
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("layerIDs", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("layerIDs", parameterToString(t, "multi"))
		}
	}
	{
		t := *r.cIIDs
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("CIIDs", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("CIIDs", parameterToString(t, "multi"))
		}
	}
	if r.atTime != nil {
		localVarQueryParams.Add("atTime", parameterToString(*r.atTime, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;odata.metadata=minimal;odata.streaming=true", "application/json;odata.metadata=minimal;odata.streaming=false", "application/json;odata.metadata=minimal", "application/json;odata.metadata=full;odata.streaming=true", "application/json;odata.metadata=full;odata.streaming=false", "application/json;odata.metadata=full", "application/json;odata.metadata=none;odata.streaming=true", "application/json;odata.metadata=none;odata.streaming=false", "application/json;odata.metadata=none", "application/json;odata.streaming=true", "application/json;odata.streaming=false", "application/json", "application/xml", "application/odata", "text/plain", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
