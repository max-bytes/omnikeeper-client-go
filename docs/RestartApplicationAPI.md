# \RestartApplicationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestartApplicationRestart**](RestartApplicationAPI.md#RestartApplicationRestart) | **Get** /api/v{version}/RestartApplication/restart | 



## RestartApplicationRestart

> RestartApplicationRestart(ctx, version).Execute()



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
    version := "version_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RestartApplicationAPI.RestartApplicationRestart(context.Background(), version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestartApplicationAPI.RestartApplicationRestart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartApplicationRestartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

