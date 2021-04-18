# Go API client for okclient

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 0.11.3
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./okclient"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AttributeApi* | [**BulkReplaceAttributesInLayer**](docs/AttributeApi.md#bulkreplaceattributesinlayer) | **Post** /api/v{version}/Attribute/bulkReplaceAttributesInLayer | bulk replace all attributes in specified layer
*AttributeApi* | [**FindMergedAttributesByName**](docs/AttributeApi.md#findmergedattributesbyname) | **Get** /api/v{version}/Attribute/findMergedAttributesByName | 
*AttributeApi* | [**GetMergedAttribute**](docs/AttributeApi.md#getmergedattribute) | **Get** /api/v{version}/Attribute/getMergedAttribute | 
*AttributeApi* | [**GetMergedAttributes**](docs/AttributeApi.md#getmergedattributes) | **Get** /api/v{version}/Attribute/getMergedAttributes | 
*AttributeApi* | [**GetMergedAttributesWithName**](docs/AttributeApi.md#getmergedattributeswithname) | **Get** /api/v{version}/Attribute/getMergedAttributesWithName | 
*CIApi* | [**GetAllCIIDs**](docs/CIApi.md#getallciids) | **Get** /api/v{version}/CI/getAllCIIDs | list of all CI-IDs
*CIApi* | [**GetCIByID**](docs/CIApi.md#getcibyid) | **Get** /api/v{version}/CI/getCIByID | single CI by CI-ID
*CIApi* | [**GetCIsByID**](docs/CIApi.md#getcisbyid) | **Get** /api/v{version}/CI/getCIsByID | multiple CIs by CI-ID  !Watch out for the query URL getting too long because of a lot of CIIDs!  TODO: consider using POST
*CISearchApi* | [**SearchCIsByTraits**](docs/CISearchApi.md#searchcisbytraits) | **Get** /api/v{version}/CISearch/searchCIsByTraits | 
*GraphQLApi* | [**Debug**](docs/GraphQLApi.md#debug) | **Post** /graphql-debug | 
*GraphQLApi* | [**Index**](docs/GraphQLApi.md#index) | **Post** /graphql | 
*GridViewApi* | [**AddContext**](docs/GridViewApi.md#addcontext) | **Post** /api/v{version}/GridView/context | Adds new context
*GridViewApi* | [**ChangeData**](docs/GridViewApi.md#changedata) | **Post** /api/v{version}/GridView/contexts/{context}/change | Saves grid view row changes and returns change results
*GridViewApi* | [**DeleteContext**](docs/GridViewApi.md#deletecontext) | **Delete** /api/v{version}/GridView/context/{name} | Deletes specific context
*GridViewApi* | [**EditContext**](docs/GridViewApi.md#editcontext) | **Put** /api/v{version}/GridView/context/{name} | Edits specific context
*GridViewApi* | [**GetContext**](docs/GridViewApi.md#getcontext) | **Get** /api/v{version}/GridView/context/{name} | Returns a single context in full
*GridViewApi* | [**GetContexts**](docs/GridViewApi.md#getcontexts) | **Get** /api/v{version}/GridView/contexts | Returns a list of contexts for grid view.
*GridViewApi* | [**GetData**](docs/GridViewApi.md#getdata) | **Get** /api/v{version}/GridView/contexts/{context}/data | Returns grid view data for specific context
*GridViewApi* | [**GetSchema**](docs/GridViewApi.md#getschema) | **Get** /api/v{version}/GridView/contexts/{context}/schema | Returns grid view schema for specific context
*LayerApi* | [**GetAllLayers**](docs/LayerApi.md#getalllayers) | **Get** /api/v{version}/Layer/getAllLayers | list of all layers
*LayerApi* | [**GetLayerByName**](docs/LayerApi.md#getlayerbyname) | **Get** /api/v{version}/Layer/getLayerByName | get a layer by name
*LayerApi* | [**GetLayersByName**](docs/LayerApi.md#getlayersbyname) | **Get** /api/v{version}/Layer/getLayersByName | get layers by name
*RelationApi* | [**GetAllMergedRelations**](docs/RelationApi.md#getallmergedrelations) | **Get** /api/v{version}/Relation/getAllMergedRelations | 
*RelationApi* | [**GetMergedRelation**](docs/RelationApi.md#getmergedrelation) | **Get** /api/v{version}/Relation/getMergedRelation | 
*RelationApi* | [**GetMergedRelationsFromOrToCI**](docs/RelationApi.md#getmergedrelationsfromortoci) | **Get** /api/v{version}/Relation/getMergedRelationsFromOrToCI | 
*RelationApi* | [**GetMergedRelationsOutgoingFromCI**](docs/RelationApi.md#getmergedrelationsoutgoingfromci) | **Get** /api/v{version}/Relation/getMergedRelationsOutgoingFromCI | 
*RelationApi* | [**GetMergedRelationsWithPredicate**](docs/RelationApi.md#getmergedrelationswithpredicate) | **Get** /api/v{version}/Relation/getMergedRelationsWithPredicate | 
*TraitApi* | [**GetEffectiveTraitsForTraitName**](docs/TraitApi.md#geteffectivetraitsfortraitname) | **Get** /api/v{version}/Trait/getEffectiveTraitsForTraitName | 


## Documentation For Models

 - [AddContextRequest](docs/AddContextRequest.md)
 - [AttributeState](docs/AttributeState.md)
 - [AttributeValueDTO](docs/AttributeValueDTO.md)
 - [AttributeValueType](docs/AttributeValueType.md)
 - [BulkCIAttributeLayerScopeDTO](docs/BulkCIAttributeLayerScopeDTO.md)
 - [CIAttributeDTO](docs/CIAttributeDTO.md)
 - [CIDTO](docs/CIDTO.md)
 - [ChangeDataCell](docs/ChangeDataCell.md)
 - [ChangeDataRequest](docs/ChangeDataRequest.md)
 - [EditContextRequest](docs/EditContextRequest.md)
 - [EffectiveTraitDTO](docs/EffectiveTraitDTO.md)
 - [FragmentDTO](docs/FragmentDTO.md)
 - [GraphQLQuery](docs/GraphQLQuery.md)
 - [GridViewColumn](docs/GridViewColumn.md)
 - [GridViewConfiguration](docs/GridViewConfiguration.md)
 - [LayerDTO](docs/LayerDTO.md)
 - [PredicateDTO](docs/PredicateDTO.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [RelatedCIDTO](docs/RelatedCIDTO.md)
 - [RelationDTO](docs/RelationDTO.md)
 - [RelationState](docs/RelationState.md)
 - [SparseRow](docs/SparseRow.md)


## Documentation For Authorization



### oauth2


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: file:///protocol/openid-connect/auth
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
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



