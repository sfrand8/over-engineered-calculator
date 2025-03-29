# CalculateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **string** | @Description The result of the mathematical expression evaluation @Example \&quot;42.0\&quot; // Optional: Provide an example result | [optional] 

## Methods

### NewCalculateResponse

`func NewCalculateResponse() *CalculateResponse`

NewCalculateResponse instantiates a new CalculateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalculateResponseWithDefaults

`func NewCalculateResponseWithDefaults() *CalculateResponse`

NewCalculateResponseWithDefaults instantiates a new CalculateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *CalculateResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CalculateResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CalculateResponse) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *CalculateResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


