# GetHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalculationHistory** | Pointer to [**[]GetHistoryHistoryEntry**](GetHistoryHistoryEntry.md) |  | [optional] 

## Methods

### NewGetHistoryResponse

`func NewGetHistoryResponse() *GetHistoryResponse`

NewGetHistoryResponse instantiates a new GetHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHistoryResponseWithDefaults

`func NewGetHistoryResponseWithDefaults() *GetHistoryResponse`

NewGetHistoryResponseWithDefaults instantiates a new GetHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalculationHistory

`func (o *GetHistoryResponse) GetCalculationHistory() []GetHistoryHistoryEntry`

GetCalculationHistory returns the CalculationHistory field if non-nil, zero value otherwise.

### GetCalculationHistoryOk

`func (o *GetHistoryResponse) GetCalculationHistoryOk() (*[]GetHistoryHistoryEntry, bool)`

GetCalculationHistoryOk returns a tuple with the CalculationHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationHistory

`func (o *GetHistoryResponse) SetCalculationHistory(v []GetHistoryHistoryEntry)`

SetCalculationHistory sets CalculationHistory field to given value.

### HasCalculationHistory

`func (o *GetHistoryResponse) HasCalculationHistory() bool`

HasCalculationHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


