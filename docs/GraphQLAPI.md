# \GraphQLAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphQLDebug**](GraphQLAPI.md#GraphQLDebug) | **Post** /graphql-debug | 
[**GraphQLGet**](GraphQLAPI.md#GraphQLGet) | **Get** /graphql | 
[**GraphQLIndex**](GraphQLAPI.md#GraphQLIndex) | **Post** /graphql | 



## GraphQLDebug

> GraphQLDebug(ctx).GraphQLQuery(graphQLQuery).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/max-bytes/omnikeeper-client-go"
)

func main() {
	graphQLQuery := *openapiclient.NewGraphQLQuery() // GraphQLQuery |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GraphQLAPI.GraphQLDebug(context.Background()).GraphQLQuery(graphQLQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphQLAPI.GraphQLDebug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphQLDebugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphQLQuery** | [**GraphQLQuery**](GraphQLQuery.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/xml, text/plain, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphQLGet

> GraphQLGet(ctx).OperationName(operationName).Query(query).Variables(variables).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/max-bytes/omnikeeper-client-go"
)

func main() {
	operationName := "operationName_example" // string |  (optional)
	query := "query_example" // string |  (optional)
	variables := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GraphQLAPI.GraphQLGet(context.Background()).OperationName(operationName).Query(query).Variables(variables).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphQLAPI.GraphQLGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphQLGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operationName** | **string** |  | 
 **query** | **string** |  | 
 **variables** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphQLIndex

> GraphQLIndex(ctx).GraphQLQuery(graphQLQuery).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/max-bytes/omnikeeper-client-go"
)

func main() {
	graphQLQuery := *openapiclient.NewGraphQLQuery() // GraphQLQuery |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GraphQLAPI.GraphQLIndex(context.Background()).GraphQLQuery(graphQLQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphQLAPI.GraphQLIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphQLIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphQLQuery** | [**GraphQLQuery**](GraphQLQuery.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/xml, text/plain, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

