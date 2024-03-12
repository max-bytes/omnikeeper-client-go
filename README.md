# Go API client for okclient

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 30.0.0
- Generator version: 7.5.0-SNAPSHOT
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import okclient "github.com/max-bytes/omnikeeper-client-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `okclient.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), okclient.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `okclient.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), okclient.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `okclient.ContextOperationServerIndices` and `okclient.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), okclient.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), okclient.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AnsibleInventoryScanIngestAPI* | [**AnsibleInventoryScanIngestIngestAnsibleInventoryScan**](docs/AnsibleInventoryScanIngestAPI.md#ansibleinventoryscaningestingestansibleinventoryscan) | **Post** /api/v{version}/Ingest/AnsibleInventoryScan | 
*AuthRedirectAPI* | [**AuthRedirectIndex**](docs/AuthRedirectAPI.md#authredirectindex) | **Get** /.well-known/openid-configuration | 
*CytoscapeAPI* | [**CytoscapeTraitCentric**](docs/CytoscapeAPI.md#cytoscapetraitcentric) | **Get** /api/v{version}/Cytoscape/traitCentric | 
*GraphQLAPI* | [**GraphQLDebug**](docs/GraphQLAPI.md#graphqldebug) | **Post** /graphql-debug | 
*GraphQLAPI* | [**GraphQLGet**](docs/GraphQLAPI.md#graphqlget) | **Get** /graphql | 
*GraphQLAPI* | [**GraphQLIndex**](docs/GraphQLAPI.md#graphqlindex) | **Post** /graphql | 
*GraphvizDotAPI* | [**GraphvizDotLayerCentric**](docs/GraphvizDotAPI.md#graphvizdotlayercentric) | **Get** /api/v{version}/GraphvizDot/layerCentric | 
*GraphvizDotAPI* | [**GraphvizDotTraitCentric**](docs/GraphvizDotAPI.md#graphvizdottraitcentric) | **Get** /api/v{version}/GraphvizDot/traitCentric | 
*GridViewAPI* | [**GridViewAddContext**](docs/GridViewAPI.md#gridviewaddcontext) | **Post** /api/v{version}/GridView/context | Adds new context
*GridViewAPI* | [**GridViewChangeData**](docs/GridViewAPI.md#gridviewchangedata) | **Post** /api/v{version}/GridView/contexts/{context}/change | Saves grid view row changes and returns change results
*GridViewAPI* | [**GridViewDeleteContext**](docs/GridViewAPI.md#gridviewdeletecontext) | **Delete** /api/v{version}/GridView/context/{name} | Deletes specific context
*GridViewAPI* | [**GridViewEditContext**](docs/GridViewAPI.md#gridvieweditcontext) | **Put** /api/v{version}/GridView/context/{name} | Edits specific context
*GridViewAPI* | [**GridViewGetData**](docs/GridViewAPI.md#gridviewgetdata) | **Get** /api/v{version}/GridView/contexts/{context}/data | Returns grid view data for specific context
*GridViewAPI* | [**GridViewGetGridViewContext**](docs/GridViewAPI.md#gridviewgetgridviewcontext) | **Get** /api/v{version}/GridView/context/{name} | Returns a single context in full
*GridViewAPI* | [**GridViewGetGridViewContexts**](docs/GridViewAPI.md#gridviewgetgridviewcontexts) | **Get** /api/v{version}/GridView/contexts | Returns a list of contexts for grid view.
*GridViewAPI* | [**GridViewGetSchema**](docs/GridViewAPI.md#gridviewgetschema) | **Get** /api/v{version}/GridView/contexts/{context}/schema | Returns grid view schema for specific context
*ImportExportLayerAPI* | [**ImportExportLayerExportLayer**](docs/ImportExportLayerAPI.md#importexportlayerexportlayer) | **Get** /api/v{version}/ImportExportLayer/exportLayer | 
*MetadataAPI* | [**MetadataGetMetadata**](docs/MetadataAPI.md#metadatagetmetadata) | **Get** /api/odata/{context}/$metadata | 
*MetadataAPI* | [**MetadataGetServiceDocument**](docs/MetadataAPI.md#metadatagetservicedocument) | **Get** /api/odata/{context} | 
*OKPluginGenericJSONIngestAPI* | [**ManageContextGetAllContexts**](docs/OKPluginGenericJSONIngestAPI.md#managecontextgetallcontexts) | **Get** /api/v{version}/ingest/genericJSON/manage/context | 
*OKPluginGenericJSONIngestAPI* | [**ManageContextGetContext**](docs/OKPluginGenericJSONIngestAPI.md#managecontextgetcontext) | **Get** /api/v{version}/ingest/genericJSON/manage/context/{id} | 
*OKPluginGenericJSONIngestAPI* | [**ManageContextRemoveContext**](docs/OKPluginGenericJSONIngestAPI.md#managecontextremovecontext) | **Delete** /api/v{version}/ingest/genericJSON/manage/context/{id} | 
*OKPluginGenericJSONIngestAPI* | [**ManageContextUpsertContext**](docs/OKPluginGenericJSONIngestAPI.md#managecontextupsertcontext) | **Post** /api/v{version}/ingest/genericJSON/manage/context | 
*OKPluginGenericJSONIngestAPI* | [**PassiveDataIngest**](docs/OKPluginGenericJSONIngestAPI.md#passivedataingest) | **Post** /api/v{version}/ingest/genericJSON/data | 
*RestartApplicationAPI* | [**RestartApplicationRestart**](docs/RestartApplicationAPI.md#restartapplicationrestart) | **Get** /api/v{version}/RestartApplication/restart | 
*UsageStatsAPI* | [**UsageStatsFetch**](docs/UsageStatsAPI.md#usagestatsfetch) | **Get** /api/v{version}/UsageStats/fetch | 


## Documentation For Models

 - [AbstractInboundIDMethod](docs/AbstractInboundIDMethod.md)
 - [AddContextRequest](docs/AddContextRequest.md)
 - [AnsibleInventoryScanDTO](docs/AnsibleInventoryScanDTO.md)
 - [AttributeValueDTO](docs/AttributeValueDTO.md)
 - [AttributeValueType](docs/AttributeValueType.md)
 - [ChangeDataCell](docs/ChangeDataCell.md)
 - [ChangeDataRequest](docs/ChangeDataRequest.md)
 - [EditContextRequest](docs/EditContextRequest.md)
 - [EdmContainerElementKind](docs/EdmContainerElementKind.md)
 - [EdmExpressionKind](docs/EdmExpressionKind.md)
 - [EdmSchemaElementKind](docs/EdmSchemaElementKind.md)
 - [EdmTypeKind](docs/EdmTypeKind.md)
 - [GenericInboundAttribute](docs/GenericInboundAttribute.md)
 - [GenericInboundCI](docs/GenericInboundCI.md)
 - [GenericInboundCIIdMethod](docs/GenericInboundCIIdMethod.md)
 - [GenericInboundData](docs/GenericInboundData.md)
 - [GenericInboundRelation](docs/GenericInboundRelation.md)
 - [GraphQLQuery](docs/GraphQLQuery.md)
 - [GridViewColumn](docs/GridViewColumn.md)
 - [GridViewConfiguration](docs/GridViewConfiguration.md)
 - [IEdmEntityContainer](docs/IEdmEntityContainer.md)
 - [IEdmEntityContainerElement](docs/IEdmEntityContainerElement.md)
 - [IEdmExpression](docs/IEdmExpression.md)
 - [IEdmModel](docs/IEdmModel.md)
 - [IEdmSchemaElement](docs/IEdmSchemaElement.md)
 - [IEdmTerm](docs/IEdmTerm.md)
 - [IEdmType](docs/IEdmType.md)
 - [IEdmTypeReference](docs/IEdmTypeReference.md)
 - [IEdmVocabularyAnnotation](docs/IEdmVocabularyAnnotation.md)
 - [InboundIDMethodByAttribute](docs/InboundIDMethodByAttribute.md)
 - [InboundIDMethodByAttributeModifiers](docs/InboundIDMethodByAttributeModifiers.md)
 - [InboundIDMethodByByUnion](docs/InboundIDMethodByByUnion.md)
 - [InboundIDMethodByData](docs/InboundIDMethodByData.md)
 - [InboundIDMethodByIntersect](docs/InboundIDMethodByIntersect.md)
 - [InboundIDMethodByRelatedTempID](docs/InboundIDMethodByRelatedTempID.md)
 - [InboundIDMethodByTemporaryCIID](docs/InboundIDMethodByTemporaryCIID.md)
 - [NoFoundTargetCIHandling](docs/NoFoundTargetCIHandling.md)
 - [ODataEntitySetInfo](docs/ODataEntitySetInfo.md)
 - [ODataFunctionImportInfo](docs/ODataFunctionImportInfo.md)
 - [ODataServiceDocument](docs/ODataServiceDocument.md)
 - [ODataSingletonInfo](docs/ODataSingletonInfo.md)
 - [ODataTypeAnnotation](docs/ODataTypeAnnotation.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [SameTargetCIHandling](docs/SameTargetCIHandling.md)
 - [SameTempIDHandling](docs/SameTempIDHandling.md)
 - [SparseRow](docs/SparseRow.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### oauth2


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: N/A

Example

```go
auth := context.WithValue(context.Background(), okclient.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, okclient.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://[keycloak-url]/auth/realms/acme/protocol/openid-connect/auth
- **Scopes**: N/A

Example

```go
auth := context.WithValue(context.Background(), okclient.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, okclient.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



