# \AnsibleInventoryScanIngestAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnsibleInventoryScanIngestIngestAnsibleInventoryScan**](AnsibleInventoryScanIngestAPI.md#AnsibleInventoryScanIngestIngestAnsibleInventoryScan) | **Post** /api/v{version}/Ingest/AnsibleInventoryScan | 



## AnsibleInventoryScanIngestIngestAnsibleInventoryScan

> AnsibleInventoryScanIngestIngestAnsibleInventoryScan(ctx, version).WriteLayerID(writeLayerID).SearchLayerIDs(searchLayerIDs).AnsibleInventoryScanDTO(ansibleInventoryScanDTO).Execute()



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
    writeLayerID := "writeLayerID_example" // string | 
    searchLayerIDs := []string{"Inner_example"} // []string | 
    version := "version_example" // string | 
    ansibleInventoryScanDTO := *openapiclient.NewAnsibleInventoryScanDTO(map[string]string{"key": "Inner_example"}, map[string]string{"key": "Inner_example"}, map[string]string{"key": "Inner_example"}, map[string]string{"key": "Inner_example"}) // AnsibleInventoryScanDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnsibleInventoryScanIngestAPI.AnsibleInventoryScanIngestIngestAnsibleInventoryScan(context.Background(), version).WriteLayerID(writeLayerID).SearchLayerIDs(searchLayerIDs).AnsibleInventoryScanDTO(ansibleInventoryScanDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnsibleInventoryScanIngestAPI.AnsibleInventoryScanIngestIngestAnsibleInventoryScan``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAnsibleInventoryScanIngestIngestAnsibleInventoryScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writeLayerID** | **string** |  | 
 **searchLayerIDs** | **[]string** |  | 

 **ansibleInventoryScanDTO** | [**AnsibleInventoryScanDTO**](AnsibleInventoryScanDTO.md) |  | 

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

