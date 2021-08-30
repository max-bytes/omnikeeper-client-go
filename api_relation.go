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

// RelationApiService RelationApi service
type RelationApiService service

type ApiGetAllMergedRelationsRequest struct {
	ctx _context.Context
	ApiService *RelationApiService
	layerIDs *[]string
	version string
	atTime *time.Time
}

func (r ApiGetAllMergedRelationsRequest) LayerIDs(layerIDs []string) ApiGetAllMergedRelationsRequest {
	r.layerIDs = &layerIDs
	return r
}
func (r ApiGetAllMergedRelationsRequest) AtTime(atTime time.Time) ApiGetAllMergedRelationsRequest {
	r.atTime = &atTime
	return r
}

func (r ApiGetAllMergedRelationsRequest) Execute() ([]RelationDTO, *_nethttp.Response, error) {
	return r.ApiService.GetAllMergedRelationsExecute(r)
}

/*
GetAllMergedRelations Method for GetAllMergedRelations

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param version
 @return ApiGetAllMergedRelationsRequest
*/
func (a *RelationApiService) GetAllMergedRelations(ctx _context.Context, version string) ApiGetAllMergedRelationsRequest {
	return ApiGetAllMergedRelationsRequest{
		ApiService: a,
		ctx: ctx,
		version: version,
	}
}

// Execute executes the request
//  @return []RelationDTO
func (a *RelationApiService) GetAllMergedRelationsExecute(r ApiGetAllMergedRelationsRequest) ([]RelationDTO, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []RelationDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationApiService.GetAllMergedRelations")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v{version}/Relation/getAllMergedRelations"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", _neturl.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.layerIDs == nil {
		return localVarReturnValue, nil, reportError("layerIDs is required and must be specified")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

type ApiGetMergedRelationRequest struct {
	ctx _context.Context
	ApiService *RelationApiService
	fromCIID *string
	toCIID *string
	predicateID *string
	layerIDs *[]string
	version string
	atTime *time.Time
}

func (r ApiGetMergedRelationRequest) FromCIID(fromCIID string) ApiGetMergedRelationRequest {
	r.fromCIID = &fromCIID
	return r
}
func (r ApiGetMergedRelationRequest) ToCIID(toCIID string) ApiGetMergedRelationRequest {
	r.toCIID = &toCIID
	return r
}
func (r ApiGetMergedRelationRequest) PredicateID(predicateID string) ApiGetMergedRelationRequest {
	r.predicateID = &predicateID
	return r
}
func (r ApiGetMergedRelationRequest) LayerIDs(layerIDs []string) ApiGetMergedRelationRequest {
	r.layerIDs = &layerIDs
	return r
}
func (r ApiGetMergedRelationRequest) AtTime(atTime time.Time) ApiGetMergedRelationRequest {
	r.atTime = &atTime
	return r
}

func (r ApiGetMergedRelationRequest) Execute() (RelationDTO, *_nethttp.Response, error) {
	return r.ApiService.GetMergedRelationExecute(r)
}

/*
GetMergedRelation Method for GetMergedRelation

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param version
 @return ApiGetMergedRelationRequest
*/
func (a *RelationApiService) GetMergedRelation(ctx _context.Context, version string) ApiGetMergedRelationRequest {
	return ApiGetMergedRelationRequest{
		ApiService: a,
		ctx: ctx,
		version: version,
	}
}

// Execute executes the request
//  @return RelationDTO
func (a *RelationApiService) GetMergedRelationExecute(r ApiGetMergedRelationRequest) (RelationDTO, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RelationDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationApiService.GetMergedRelation")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v{version}/Relation/getMergedRelation"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", _neturl.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.fromCIID == nil {
		return localVarReturnValue, nil, reportError("fromCIID is required and must be specified")
	}
	if r.toCIID == nil {
		return localVarReturnValue, nil, reportError("toCIID is required and must be specified")
	}
	if r.predicateID == nil {
		return localVarReturnValue, nil, reportError("predicateID is required and must be specified")
	}
	if r.layerIDs == nil {
		return localVarReturnValue, nil, reportError("layerIDs is required and must be specified")
	}

	localVarQueryParams.Add("fromCIID", parameterToString(*r.fromCIID, ""))
	localVarQueryParams.Add("toCIID", parameterToString(*r.toCIID, ""))
	localVarQueryParams.Add("predicateID", parameterToString(*r.predicateID, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

type ApiGetMergedRelationsFromOrToCIRequest struct {
	ctx _context.Context
	ApiService *RelationApiService
	ciid *string
	layerIDs *[]string
	version string
	atTime *time.Time
}

func (r ApiGetMergedRelationsFromOrToCIRequest) Ciid(ciid string) ApiGetMergedRelationsFromOrToCIRequest {
	r.ciid = &ciid
	return r
}
func (r ApiGetMergedRelationsFromOrToCIRequest) LayerIDs(layerIDs []string) ApiGetMergedRelationsFromOrToCIRequest {
	r.layerIDs = &layerIDs
	return r
}
func (r ApiGetMergedRelationsFromOrToCIRequest) AtTime(atTime time.Time) ApiGetMergedRelationsFromOrToCIRequest {
	r.atTime = &atTime
	return r
}

func (r ApiGetMergedRelationsFromOrToCIRequest) Execute() ([]RelationDTO, *_nethttp.Response, error) {
	return r.ApiService.GetMergedRelationsFromOrToCIExecute(r)
}

/*
GetMergedRelationsFromOrToCI Method for GetMergedRelationsFromOrToCI

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param version
 @return ApiGetMergedRelationsFromOrToCIRequest
*/
func (a *RelationApiService) GetMergedRelationsFromOrToCI(ctx _context.Context, version string) ApiGetMergedRelationsFromOrToCIRequest {
	return ApiGetMergedRelationsFromOrToCIRequest{
		ApiService: a,
		ctx: ctx,
		version: version,
	}
}

// Execute executes the request
//  @return []RelationDTO
func (a *RelationApiService) GetMergedRelationsFromOrToCIExecute(r ApiGetMergedRelationsFromOrToCIRequest) ([]RelationDTO, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []RelationDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationApiService.GetMergedRelationsFromOrToCI")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v{version}/Relation/getMergedRelationsFromOrToCI"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", _neturl.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ciid == nil {
		return localVarReturnValue, nil, reportError("ciid is required and must be specified")
	}
	if r.layerIDs == nil {
		return localVarReturnValue, nil, reportError("layerIDs is required and must be specified")
	}

	localVarQueryParams.Add("ciid", parameterToString(*r.ciid, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

type ApiGetMergedRelationsOutgoingFromCIRequest struct {
	ctx _context.Context
	ApiService *RelationApiService
	fromCIID *string
	layerIDs *[]string
	version string
	atTime *time.Time
}

func (r ApiGetMergedRelationsOutgoingFromCIRequest) FromCIID(fromCIID string) ApiGetMergedRelationsOutgoingFromCIRequest {
	r.fromCIID = &fromCIID
	return r
}
func (r ApiGetMergedRelationsOutgoingFromCIRequest) LayerIDs(layerIDs []string) ApiGetMergedRelationsOutgoingFromCIRequest {
	r.layerIDs = &layerIDs
	return r
}
func (r ApiGetMergedRelationsOutgoingFromCIRequest) AtTime(atTime time.Time) ApiGetMergedRelationsOutgoingFromCIRequest {
	r.atTime = &atTime
	return r
}

func (r ApiGetMergedRelationsOutgoingFromCIRequest) Execute() ([]RelationDTO, *_nethttp.Response, error) {
	return r.ApiService.GetMergedRelationsOutgoingFromCIExecute(r)
}

/*
GetMergedRelationsOutgoingFromCI Method for GetMergedRelationsOutgoingFromCI

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param version
 @return ApiGetMergedRelationsOutgoingFromCIRequest
*/
func (a *RelationApiService) GetMergedRelationsOutgoingFromCI(ctx _context.Context, version string) ApiGetMergedRelationsOutgoingFromCIRequest {
	return ApiGetMergedRelationsOutgoingFromCIRequest{
		ApiService: a,
		ctx: ctx,
		version: version,
	}
}

// Execute executes the request
//  @return []RelationDTO
func (a *RelationApiService) GetMergedRelationsOutgoingFromCIExecute(r ApiGetMergedRelationsOutgoingFromCIRequest) ([]RelationDTO, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []RelationDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationApiService.GetMergedRelationsOutgoingFromCI")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v{version}/Relation/getMergedRelationsOutgoingFromCI"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", _neturl.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.fromCIID == nil {
		return localVarReturnValue, nil, reportError("fromCIID is required and must be specified")
	}
	if r.layerIDs == nil {
		return localVarReturnValue, nil, reportError("layerIDs is required and must be specified")
	}

	localVarQueryParams.Add("fromCIID", parameterToString(*r.fromCIID, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

type ApiGetMergedRelationsWithPredicateRequest struct {
	ctx _context.Context
	ApiService *RelationApiService
	predicateID *string
	layerIDs *[]string
	version string
	atTime *time.Time
}

func (r ApiGetMergedRelationsWithPredicateRequest) PredicateID(predicateID string) ApiGetMergedRelationsWithPredicateRequest {
	r.predicateID = &predicateID
	return r
}
func (r ApiGetMergedRelationsWithPredicateRequest) LayerIDs(layerIDs []string) ApiGetMergedRelationsWithPredicateRequest {
	r.layerIDs = &layerIDs
	return r
}
func (r ApiGetMergedRelationsWithPredicateRequest) AtTime(atTime time.Time) ApiGetMergedRelationsWithPredicateRequest {
	r.atTime = &atTime
	return r
}

func (r ApiGetMergedRelationsWithPredicateRequest) Execute() ([]RelationDTO, *_nethttp.Response, error) {
	return r.ApiService.GetMergedRelationsWithPredicateExecute(r)
}

/*
GetMergedRelationsWithPredicate Method for GetMergedRelationsWithPredicate

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param version
 @return ApiGetMergedRelationsWithPredicateRequest
*/
func (a *RelationApiService) GetMergedRelationsWithPredicate(ctx _context.Context, version string) ApiGetMergedRelationsWithPredicateRequest {
	return ApiGetMergedRelationsWithPredicateRequest{
		ApiService: a,
		ctx: ctx,
		version: version,
	}
}

// Execute executes the request
//  @return []RelationDTO
func (a *RelationApiService) GetMergedRelationsWithPredicateExecute(r ApiGetMergedRelationsWithPredicateRequest) ([]RelationDTO, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []RelationDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationApiService.GetMergedRelationsWithPredicate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v{version}/Relation/getMergedRelationsWithPredicate"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", _neturl.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.predicateID == nil {
		return localVarReturnValue, nil, reportError("predicateID is required and must be specified")
	}
	if r.layerIDs == nil {
		return localVarReturnValue, nil, reportError("layerIDs is required and must be specified")
	}

	localVarQueryParams.Add("predicateID", parameterToString(*r.predicateID, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
