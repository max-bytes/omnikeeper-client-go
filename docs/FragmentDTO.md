# FragmentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | [**AttributeValueDTO**](AttributeValueDTO.md) |  | 
**Ciid** | **string** |  | 

## Methods

### NewFragmentDTO

`func NewFragmentDTO(name string, value AttributeValueDTO, ciid string, ) *FragmentDTO`

NewFragmentDTO instantiates a new FragmentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFragmentDTOWithDefaults

`func NewFragmentDTOWithDefaults() *FragmentDTO`

NewFragmentDTOWithDefaults instantiates a new FragmentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FragmentDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FragmentDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FragmentDTO) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *FragmentDTO) GetValue() AttributeValueDTO`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FragmentDTO) GetValueOk() (*AttributeValueDTO, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FragmentDTO) SetValue(v AttributeValueDTO)`

SetValue sets Value field to given value.


### GetCiid

`func (o *FragmentDTO) GetCiid() string`

GetCiid returns the Ciid field if non-nil, zero value otherwise.

### GetCiidOk

`func (o *FragmentDTO) GetCiidOk() (*string, bool)`

GetCiidOk returns a tuple with the Ciid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiid

`func (o *FragmentDTO) SetCiid(v string)`

SetCiid sets Ciid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


